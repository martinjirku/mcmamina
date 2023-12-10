package handlers

import (
	"log/slog"
	"net/http"

	"jirku.sk/mcmamina/components"
	"jirku.sk/mcmamina/components/pages"
)

type DefaultHandler struct {
	Log *slog.Logger
}

type ViewProps struct{}

func (h *DefaultHandler) View(w http.ResponseWriter, r *http.Request, props ViewProps) {
	components.Page(components.NewPage("Index", "Domov", pages.IndexPage())).Render(r.Context(), w)
}

func (h *DefaultHandler) Get(w http.ResponseWriter, r *http.Request) {
	var err error
	props := ViewProps{}
	if err != nil {
		h.Log.Error("failed to get counts", slog.Any("error", err))
		http.Error(w, "failed to get counts", http.StatusInternalServerError)
		return
	}
	h.View(w, r, props)
}

func (h *DefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Log.Info("request", slog.String("method", r.Method), slog.String("path", r.URL.Path))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.Method == http.MethodPost {
		h.Get(w, r)
		return
	}
	h.Get(w, r)
}
