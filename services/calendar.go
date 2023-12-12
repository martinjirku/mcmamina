package services

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"jirku.sk/mcmamina/models"
)

type CalendarService struct {
	apiKey     string
	calendarID string
}

func NewCalendarService(apiKey, calendarID string) *CalendarService {
	return &CalendarService{
		apiKey:     apiKey,
		calendarID: calendarID,
	}
}

func (s *CalendarService) GetEvents(ctx context.Context, timeMin, timeMax time.Time) ([]models.Event, error) {
	calendarService, err := calendar.NewService(ctx, option.WithAPIKey(s.apiKey))
	if err != nil {
		return nil, fmt.Errorf("creating calendar service: %w", err)
	}
	eventsService := calendar.NewEventsService(calendarService)
	call := eventsService.
		List(s.calendarID).
		TimeMin(timeMin.Format(time.RFC3339)).
		TimeMax(timeMax.Format(time.RFC3339))
	events, err := call.Do()
	if err != nil {
		return nil, fmt.Errorf("getting events: %w", err)
	}
	result := make([]models.Event, len(events.Items))
	for i, event := range events.Items {
		if event == nil {
			continue
		}
		result[i] = *models.NewEvent(event.Summary)
		if event.Start != nil {
			result[i].SetStart(event.Start.DateTime)
		}
		if event.End != nil {
			result[i].SetEnd(event.End.DateTime)
		}

	}
	return result, nil
}
