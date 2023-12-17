package main

import (
	"log/slog"
	"net/http"
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

func main() {
	log := slog.Default()
	var env map[string]string
	env, err := godotenv.Read()
	if err != nil {
		log.Error("Error loading .env file", slog.Any("error", err))
	}
	calendarService := services.NewCalendarService(env[GOOGLE_API_KEY], env[GOOGLE_CALENDAR_ID])
	setupWebserver(log, calendarService)

}

func setupWebserver(log *slog.Logger, calendarService *services.CalendarService) {
	router := mux.NewRouter()
	distPath := "/dist"
	cssService := services.NewCSS(distPath)
	sponsorService := services.NewSponsorService()

	router.HandleFunc("/", handlers.NewIndexHandler(log, calendarService, cssService).ServeHTTP)
	router.HandleFunc("/o-nas", handlers.AboutUs(log, cssService))
	// MCMAMINA -->> GENERATED CODE
	router.HandleFunc("/podpora/2-percenta-z-dane", handlers.TaxBonus(log, cssService))

	router.HandleFunc("/podpora", handlers.SupportedUs(log, cssService, sponsorService))
	router.HandleFunc("/kalendar", handlers.Calendar(log, cssService))

	// MCMAMINA <<-- GENERATED CODE

	handleFiles(router, "/images/", "./assets/images")
	handleFiles(router, "/", "."+distPath)

	http.ListenAndServe("localhost:3000", router)
}

func handleFiles(r *mux.Router, path, dir string) {
	fs := http.StripPrefix(path, http.FileServer(http.Dir(dir)))

	r.PathPrefix(path).Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
