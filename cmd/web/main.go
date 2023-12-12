package main

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"jirku.sk/mcmamina/components/pages"
	"jirku.sk/mcmamina/handlers"
)

// .env keys
const (
	GOOGLE_API_KEY     = "GOOGLE_API_KEY"
	GOOGLE_CALENDAR_ID = "GOOGLE_CALENDAR_ID"
)

func main() {
	log := slog.Default()
	router := mux.NewRouter()

	var env map[string]string
	env, err := godotenv.Read()
	if err != nil {
		log.Error("Error loading .env file", slog.Any("error", err))
	}

	handleFiles(router, "/images/", "./assets/images")
	handleFiles(router, "/dist/", "./dist")
	router.HandleFunc("/", defaultHandler(log, env[GOOGLE_API_KEY], env[GOOGLE_CALENDAR_ID]).ServeHTTP)
	router.HandleFunc("/pages.css", CssHandler())
	http.ListenAndServe("localhost:3000", router)
}

func defaultHandler(log *slog.Logger, apiKey, calendarID string) http.Handler {
	defaultHandler := handlers.DefaultHandler{
		Log:        log,
		ApiKey:     apiKey,
		CalendarID: calendarID,
	}
	return &defaultHandler
}

func CssHandler() func(http.ResponseWriter, *http.Request) {
	handler := templ.NewCSSHandler(pages.IndexCss())
	return handler.ServeHTTP
}

func handleFiles(r *mux.Router, path, dir string) {
	handler := http.StripPrefix(path, http.FileServer(http.Dir(dir)))
	r.PathPrefix(path).Handler(handler)
}
