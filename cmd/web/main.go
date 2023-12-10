package main

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"jirku.sk/mcmamina/handlers"
)

func main() {
	log := slog.Default()
	r := mux.NewRouter()

	r.HandleFunc("/", defaultHandler(log).ServeHTTP)

	imagesHandler := http.StripPrefix("/images/", http.FileServer(http.Dir("./assets/images")))
	r.PathPrefix("/images/").Handler(imagesHandler)

	distHandler := http.FileServer(http.Dir("./dist"))
	r.PathPrefix("/").Handler(http.StripPrefix("/", distHandler))

	http.ListenAndServe("localhost:3000", r)
}

func defaultHandler(log *slog.Logger) http.Handler {
	defaultHandler := handlers.DefaultHandler{
		Log: log,
	}
	return &defaultHandler
}
