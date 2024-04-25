package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"jirku.sk/mcmamina/handlers"
	"jirku.sk/mcmamina/pkg/middleware"
	"jirku.sk/mcmamina/pkg/services"
)

// .env keys
const (
	GOOGLE_API_KEY      = "GOOGLE_API_KEY"
	GOOGLE_CALENDAR_ID  = "GOOGLE_CALENDAR_ID"
	GOOGLE_CAPTCHA_SITE = "GOOGLE_CAPTCHA_SITE"
)

//go:embed dist dist/.vite
var distFS embed.FS

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file", slog.Any("error", err))
	}
	setupWebserver(log)
}

type configuration struct {
	publicPath  string
	port        int
	host        string
	panoramaURL string
}

func setupWebserver(log *slog.Logger) {
	config := parseConfig(log)
	router := mux.NewRouter()

	var workingFolder fs.FS
	log.Info("reading public folder")
	if config.publicPath == "" {
		log.Info("No public folder found, using embed FS")
		log.Info("accessing embed FS")
		folder, err := fs.Sub(distFS, "dist")
		if err != nil {
			log.Error("accessing sub 'dist' to embed FS %w", err)
		}
		workingFolder = folder
	} else {
		log.Info(fmt.Sprintf("serving specified location %s", config.publicPath))
		workingFolder = os.DirFS(config.publicPath)
	}

	cssService, sponsorService, recaptchaService, calendarService := prepareServices(log, workingFolder)
	prepareMiddleware(router, log)

	router.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	setupPanorama(config.panoramaURL, router, log)

	router.HandleFunc("/", handlers.NewIndexHandler(log, calendarService, cssService).ServeHTTP).Methods("GET")
	router.HandleFunc("/o-nas", handlers.AboutUs(log, cssService)).Methods("GET")
	// MCMAMINA -->> GENERATED CODE
	router.HandleFunc("/prihlasenie", handlers.Login(log, cssService, recaptchaService))
	router.HandleFunc("/aktivity/burzy", handlers.Marketplace(log, cssService)).Methods("GET")
	router.HandleFunc("/aktivity/podporne-skupiny", handlers.SupportGroups(log, cssService)).Methods("GET")
	router.HandleFunc("/aktivity/predporodny-kurz", handlers.BabyDeliveryCourse(log, cssService)).Methods("GET")
	router.HandleFunc("/podpora/2-percenta-z-dane", handlers.TaxBonus(log, cssService)).Methods("GET")
	router.HandleFunc("/podpora/dobrovolnici", handlers.Volunteers(log, cssService)).Methods("GET")
	router.HandleFunc("/podpora", handlers.SupportedUs(log, cssService, sponsorService)).Methods("GET")
	router.HandleFunc("/aktivity", handlers.Activities(log, cssService)).Methods("GET")
	router.HandleFunc("/aktivity/kalendar", handlers.Calendar(log, cssService)).Methods("GET")
	// MCMAMINA <<-- GENERATED CODE
	handleFiles(router, http.FS(workingFolder))

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT)

	addr := fmt.Sprintf("%s:%d", config.host, config.port)
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	log.Info(fmt.Sprintf("starting server at http://%s", addr))
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error(fmt.Errorf("failed to start server %s", err).Error())
			os.Exit(1)
		}
	}()
	// Block until we receive our signal.
	signal := <-sigs
	log.Info(fmt.Sprintf("received \"%s\" signal, shutting down", signal))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	log.Info("shutting down...")
	srv.Shutdown(ctx)
	log.Info("server shutted down")
	os.Exit(0)
}

func parseConfig(log *slog.Logger) configuration {
	config := configuration{}
	flag.StringVar(&config.publicPath, "public", "", "Usage description of the flag")
	flag.StringVar(&config.host, "host", "0.0.0.0", "specify the app host")
	flag.StringVar(&config.panoramaURL, "panorama", "http://mcmamina.panfoto.sk/", "specify the panorama URL for reverse proxy")
	flag.IntVar(&config.port, "port", 8080, "specfiy the port application will listen")
	flag.Parse()

	if !validatePort(config.port) {
		log.Error("invalid port")
		panic("invalid port")
	}
	return config
}

func prepareServices(log *slog.Logger, fs fs.FS) (*services.CSS, *services.SponsorService, *services.RecaptchaService, *services.CalendarService) {
	cssService := services.NewCSS(fs, serviceLog(log, "css"))
	sponsorService := services.NewSponsorService()
	recaptchaService := services.NewRecaptchaService(os.Getenv(GOOGLE_API_KEY), os.Getenv(GOOGLE_CAPTCHA_SITE))
	calendarService := services.NewCalendarService(os.Getenv(GOOGLE_API_KEY), os.Getenv(GOOGLE_CALENDAR_ID))
	return cssService, sponsorService, recaptchaService, calendarService
}

func prepareMiddleware(router *mux.Router, log *slog.Logger) {
	router.Use(middleware.Recover(middlewareLog(log, "recover")))
	router.Use(middleware.RequestID(middlewareLog(log, "requestID")))
	router.Use(middleware.Logger(middlewareLog(log, "logger")))
	router.Use(middleware.Csrf)
	// router.Use(middleware.AuthMiddleware(store))
}

func setupPanorama(panoramaPath string, router *mux.Router, log *slog.Logger) {
	panoramaURL, err := url.Parse(panoramaPath)
	if err != nil {
		log.Error("invalid panorama URL")
	} else {
		router.PathPrefix("/panorama").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			proxy := httputil.NewSingleHostReverseProxy(panoramaURL)
			proxy.Director = func(req *http.Request) {
				req.URL.Scheme = panoramaURL.Scheme
				req.URL.Host = panoramaURL.Host
				originalPath := req.URL.Path
				if strings.HasPrefix(originalPath, "/panorama") {
					req.RequestURI = strings.TrimPrefix(originalPath, "/panorama")
					req.URL.Path = strings.TrimPrefix(originalPath, "/panorama")
				}
				req.Host = panoramaURL.Host
			}
			proxy.ServeHTTP(w, r)
		})
	}
}

func handleFiles(r *mux.Router, folder http.FileSystem) {
	fs := http.FileServer(folder)

	r.PathPrefix("/").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := getContentType(r.URL.Path)
		w.Header().Set("Content-Type", contentType)
		if contentType != "application/javascript" {
			w.Header().Set("Cache-Control", "public, max-age=3600") // 1 hour in seconds
		} else {
			w.Header().Set("Cache-Control", "public, max-age=60") // 1 hour in seconds
		}
		fs.ServeHTTP(w, r)
	}))
}

func getContentType(filePath string) string {
	// Default to "text/plain" if MIME type cannot be determined
	contentType := "text/plain"

	// Map file extensions to MIME types
	mimeTypes := map[string]string{
		".es":  "application/javascript",
		".js":  "application/javascript",
		".css": "text/css",
		// Add more MIME types as needed
	}

	// Get the file extension
	ext := filepath.Ext(filePath)

	// Look up the MIME type in the map
	if mt, ok := mimeTypes[ext]; ok {
		contentType = mt
	}

	return contentType
}

func validatePort(port int) bool {
	return port > 0 && port <= 65535
}

func middlewareLog(log *slog.Logger, name string) *slog.Logger {
	return log.With(slog.String("type", "middleware"), slog.String("name", name))
}

func serviceLog(log *slog.Logger, name string) *slog.Logger {
	return log.With(slog.String("type", "service"), slog.String("name", name))
}
