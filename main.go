package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"jirku.sk/mcmamina/handlers"
	"jirku.sk/mcmamina/pkg/middleware"
	"jirku.sk/mcmamina/services"
)

// .env keys
const (
	GOOGLE_API_KEY     = "GOOGLE_API_KEY"
	GOOGLE_CALENDAR_ID = "GOOGLE_CALENDAR_ID"
)

//go:embed dist dist/.vite
var distFS embed.FS

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file", slog.Any("error", err))
	}
	calendarService := services.NewCalendarService(os.Getenv(GOOGLE_API_KEY), os.Getenv(GOOGLE_CALENDAR_ID))
	setupWebserver(log, calendarService)
}

type configuration struct {
	publicPath string
	port       int
	host       string
}

func setupWebserver(log *slog.Logger, calendarService *services.CalendarService) {
	config := configuration{}
	router := mux.NewRouter()

	flag.StringVar(&config.publicPath, "public", "", "Usage description of the flag")
	flag.StringVar(&config.host, "host", "0.0.0.0", "specify the app host")
	flag.IntVar(&config.port, "port", 8080, "specfiy the port application will listen")
	flag.Parse()

	if !validatePort(config.port) {
		log.Error("invalid port")
		panic("invalid port")
	}

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

	cssService := services.NewCSS(workingFolder, serviceLog(log, "css"))
	sponsorService := services.NewSponsorService()

	// Logger middleware
	router.Use(middleware.Recover(middlwareLog(log, "recover")))
	router.Use(middleware.RequestID(middlwareLog(log, "requestID")))
	router.Use(middleware.Logger(middlwareLog(log, "logger")))

	router.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	router.HandleFunc("/", handlers.NewIndexHandler(log, calendarService, cssService).ServeHTTP)
	router.HandleFunc("/o-nas", handlers.AboutUs(log, cssService))
	// MCMAMINA -->> GENERATED CODE
	router.HandleFunc("/podpora/2-percenta-z-dane", handlers.TaxBonus(log, cssService))
	router.HandleFunc("/podpora", handlers.SupportedUs(log, cssService, sponsorService))
	router.HandleFunc("/aktivity", handlers.Activities(log, cssService))
	router.HandleFunc("/aktivity/kalendar", handlers.Calendar(log, cssService))
	// MCMAMINA <<-- GENERATED CODE

	handleFiles(router, http.FS(workingFolder))

	addr := fmt.Sprintf("%s:%d", config.host, config.port)
	log.Info(fmt.Sprintf("starting server at %s", addr))
	http.ListenAndServe(addr, router)
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

func middlwareLog(log *slog.Logger, name string) *slog.Logger {
	return log.With(slog.String("type", "middleware"), slog.String("name", name))
}

func serviceLog(log *slog.Logger, name string) *slog.Logger {
	return log.With(slog.String("type", "service"), slog.String("name", name))
}
