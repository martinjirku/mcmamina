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
	log := slog.Default()
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file", slog.Any("error", err))
	}
	calendarService := services.NewCalendarService(os.Getenv(GOOGLE_API_KEY), os.Getenv(GOOGLE_CALENDAR_ID))
	setupWebserver(log, calendarService)
}

func setupWebserver(log *slog.Logger, calendarService *services.CalendarService) {
	router := mux.NewRouter()
	var publicPath string
	flag.StringVar(&publicPath, "public", "", "Usage description of the flag")
	flag.Parse()

	var workingFolder fs.FS
	log.Info("reading public folder")
	if publicPath == "" {
		log.Info("No public folder found, using embed FS")
		log.Info("accessing embed FS")
		folder, err := fs.Sub(distFS, "dist")
		if err != nil {
			log.Error("accessing sub 'dist' to embed FS %w", err)
		}
		workingFolder = folder
	} else {
		log.Info(fmt.Sprintf("serving specified location %s", publicPath))
		workingFolder = os.DirFS(publicPath)
	}

	cssService := services.NewCSS(workingFolder, log)
	sponsorService := services.NewSponsorService()
	router.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	router.HandleFunc("/", handlers.NewIndexHandler(log, calendarService, cssService).ServeHTTP)
	router.HandleFunc("/o-nas", handlers.AboutUs(log, cssService))
	// MCMAMINA -->> GENERATED CODE
	router.HandleFunc("/podpora/2-percenta-z-dane", handlers.TaxBonus(log, cssService))
	router.HandleFunc("/podpora", handlers.SupportedUs(log, cssService, sponsorService))
	router.HandleFunc("/kalendar", handlers.Calendar(log, cssService))
	// MCMAMINA <<-- GENERATED CODE

	handleFiles(router, http.FS(workingFolder))
	http.ListenAndServe("0.0.0.0:8080", router)
}

func handleFiles(r *mux.Router, folder http.FileSystem) {
	fs := http.FileServer(folder)

	r.PathPrefix("/").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filePath := filepath.Join("your_directory_path", r.URL.Path)
		contentType := getContentType(filePath)
		w.Header().Set("Content-Type", contentType)
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
