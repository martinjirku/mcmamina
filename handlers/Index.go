package handlers

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"jirku.sk/mcmamina/models"
	"jirku.sk/mcmamina/template/components"
	"jirku.sk/mcmamina/template/pages"
)

type EventsGetter interface {
	GetEvents(ctx context.Context, fromDate, toDate time.Time) ([]models.Event, error)
}

type CSSPathGetter interface {
	GetCssPath() (string, error)
}
type IndexHandler struct {
	Log           *slog.Logger
	EventsGetter  EventsGetter
	cssPathGetter CSSPathGetter
}

func NewIndexHandler(log *slog.Logger, getter EventsGetter, cssPathGetter CSSPathGetter) http.Handler {
	IndexHandler := IndexHandler{
		Log:           log,
		EventsGetter:  getter,
		cssPathGetter: cssPathGetter,
	}
	return &IndexHandler
}

type ViewProps struct {
	days  []models.Day
	today time.Time
}

func (h *IndexHandler) View(w http.ResponseWriter, r *http.Request, props ViewProps) {
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

func (h *IndexHandler) Get(w http.ResponseWriter, r *http.Request) {
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

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.Method == http.MethodPost {
		// todo: handle POST
		h.Get(w, r)
		return
	}
	h.Get(w, r)
}
