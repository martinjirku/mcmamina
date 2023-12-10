package main

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
	"jirku.sk/mcmamina/components/pages"
	"jirku.sk/mcmamina/handlers"
)

func main() {
	log := slog.Default()
	router := mux.NewRouter()

	handleFiles(router, "/images/", "./assets/images")
	handleFiles(router, "/dist/", "./dist")
	router.HandleFunc("/", defaultHandler(log).ServeHTTP)
	router.HandleFunc("/pages.css", CssHandler())
	http.ListenAndServe("localhost:3000", router)
}

func defaultHandler(log *slog.Logger) http.Handler {
	defaultHandler := handlers.DefaultHandler{
		Log: log,
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
