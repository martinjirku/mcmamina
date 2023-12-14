package handlers

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"jirku.sk/mcmamina/components"
	"jirku.sk/mcmamina/components/pages"
	"jirku.sk/mcmamina/models"
)

type EventsGetter interface {
	GetEvents(ctx context.Context, fromDate, toDate time.Time) ([]models.Event, error)
}

type CSSPathGetter interface {
	GetCssPath() (string, error)
}
type DefaultHandler struct {
	Log           *slog.Logger
	EventsGetter  EventsGetter
	cssPathGetter CSSPathGetter
}

func NewdefaultHandler(log *slog.Logger, getter EventsGetter, cssPathGetter CSSPathGetter) http.Handler {
	defaultHandler := DefaultHandler{
		Log:           log,
		EventsGetter:  getter,
		cssPathGetter: cssPathGetter,
	}
	return &defaultHandler
}

type ViewProps struct {
	days  []models.Day
	today time.Time
}

func (h *DefaultHandler) View(w http.ResponseWriter, r *http.Request, props ViewProps) {
	cssPath, _ := h.cssPathGetter.GetCssPath()
	components.Page(components.NewPage(
		"Index",
		"Domov",
		cssPath,
		pages.IndexPage(pages.IndexPageDto{
			Days:  props.days,
			Today: props.today,
		}),
	)).Render(r.Context(), w)
}

func (h *DefaultHandler) Get(w http.ResponseWriter, r *http.Request) {
	var err error
	props := ViewProps{
		days:  models.GetNumberOfDaysAfter(time.Now(), 15),
		today: time.Now(),
	}
	events, err := h.EventsGetter.
		GetEvents(r.Context(), props.days[0].Date, props.days[len(props.days)-1].Date)
	if err != nil {
		h.Log.Error("failed to get events", slog.Any("error", err))
	} else {
		for i := range props.days {
			for _, event := range events {
				start := event.Start
				end := event.End
				currentDate := props.days[i].Date
				if start.Format(time.DateOnly) == currentDate.Format(time.DateOnly) || end.Format(time.DateOnly) == currentDate.Format(time.DateOnly) {
					props.days[i].AddEvent(event)
				} else if currentDate.After(start) && currentDate.Before(end) {
					props.days[i].AddEvent(event)
				}
			}
		}
	}
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
