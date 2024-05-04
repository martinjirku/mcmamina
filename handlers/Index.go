package handlers

import (
	"context"
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"
	"time"

	"jirku.sk/mcmamina/pkg/models"
)

type EventsGetter interface {
	GetEvents(ctx context.Context, fromDate, toDate time.Time) ([]models.Event, error)
}

type CSSPathGetter interface {
	GetCssPath() (string, error)
}

func getMenuItems(activePath string) []map[string]any {
	return []map[string]any{
		{"label": "Domov", "href": "/", "isActive": activePath == "/"},
		{"label": "O nás", "href": "/o-nas", "isActive": activePath == "/o-nas"},
		{"label": "Aktivity", "href": "/aktivity", "isActive": activePath == "/aktivity"},
		{"label": "Podporili nás", "href": "/podpora", "isActive": activePath == "/podpora"},
	}
}

func HandleError(log *slog.Logger, cssPathGetter CSSPathGetter, tmpl *template.Template, fs fs.FS) http.HandlerFunc {
	currentTmpl, err := tmpl.Clone()
	if err != nil {
		log.Error("cloning template: %w", err)
	}
	currentTmpl, err = currentTmpl.ParseFS(fs, "templates/pages/error.tmpl")
	if err != nil {
		log.Error("ParseFS 'templates/pages/index.tmpl': %w", err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Redirect(w, r, "/error", http.StatusMethodNotAllowed)
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		cssPath, _ := cssPathGetter.GetCssPath()
		if err := currentTmpl.ExecuteTemplate(w, "page", map[string]any{
			"Title":       "Domov",
			"Css":         cssPath,
			"layoutClass": "index-page w-full bg-cover bg-center text-indigo-800 font-light",
			"Menu":        getMenuItems("/error"),
		}); err != nil {
			log.Error("page executing context", err)
		}
	}
}

func NewPageHandler(log *slog.Logger, getter EventsGetter, cssPathGetter CSSPathGetter, tmpl *template.Template, fs fs.FS) http.Handler {
	currentTmpl, err := tmpl.Clone()
	if err != nil {
		log.Error("cloning template: %w", err)
	}
	currentTmpl, err = currentTmpl.ParseFS(fs, "templates/pages/index.tmpl")
	if err != nil {
		log.Error("ParseFS 'templates/pages/index.tmpl': %w", err)
	}
	handler := PageHandler{
		Log:           log,
		EventsGetter:  getter,
		cssPathGetter: cssPathGetter,
		tmpl:          currentTmpl,
	}
	return &handler
}

type PageHandler struct {
	Log           *slog.Logger
	EventsGetter  EventsGetter
	cssPathGetter CSSPathGetter
	tmpl          *template.Template
}

func (p *PageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	cssPath, _ := p.cssPathGetter.GetCssPath()

	days := models.GetNumberOfDaysAfter(time.Now(), 15)
	today := time.Now()
	events, err := p.EventsGetter.
		GetEvents(r.Context(), days[0].Date, days[len(days)-1].Date)
	if err != nil {
		p.Log.Error("failed to get events", slog.Any("error", err))
	} else {
		for i := range days {
			for _, event := range events {
				start := event.Start
				end := event.End
				currentDate := days[i].Date
				if start.Format(time.DateOnly) == currentDate.Format(time.DateOnly) || end.Format(time.DateOnly) == currentDate.Format(time.DateOnly) {
					days[i].AddEvent(event)
				} else if currentDate.After(start) && currentDate.Before(end) {
					days[i].AddEvent(event)
				}
			}
		}
	}
	if err != nil {
		p.Log.Error("failed to get counts", slog.Any("error", err))
		http.Redirect(w, r, "/error", http.StatusInternalServerError)
		return
	}

	model := map[string]any{
		"Title":       "Domov",
		"Css":         cssPath,
		"Module":      "index",
		"layoutClass": "index-page w-full bg-cover bg-center text-indigo-800 font-light",
		"Menu":        getMenuItems("/"),
		"calendar": map[string]any{
			"days":  days,
			"today": today,
		},
		"activities": models.GetActivities(),
	}
	if err := p.tmpl.ExecuteTemplate(w, "page", model); err != nil {
		p.Log.Error("page executing context", err)
		http.Redirect(w, r, "/error", http.StatusInternalServerError)
	}
}
