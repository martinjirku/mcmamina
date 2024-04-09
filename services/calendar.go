package services

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"jirku.sk/mcmamina/models"
)

type CalendarService struct {
	apiKey     string
	calendarID string
	validUntil atomic.Int64
	data       sync.Map
}

func NewCalendarService(apiKey, calendarID string) *CalendarService {
	return &CalendarService{
		apiKey:     apiKey,
		calendarID: calendarID,
	}
}

func (s *CalendarService) getEvents(ctx context.Context, timeMin, timeMax time.Time) ([]models.Event, error) {
	calendarService, err := calendar.NewService(ctx, option.WithAPIKey(s.apiKey))
	if err != nil {
		return nil, fmt.Errorf("creating calendar service: %w", err)
	}
	eventsService := calendar.NewEventsService(calendarService)
	events, err := eventsService.
		List(s.calendarID).
		Fields(googleapi.Field("items(id,summary,start,end,recurrence)")).
		TimeMin(timeMin.Format(time.RFC3339)).
		TimeMax(timeMax.Format(time.RFC3339)).
		Do()
	if err != nil {
		return nil, fmt.Errorf("getting events: %w", err)
	}

	var result []models.Event

	eventTitleToIgnore := strings.ToLower("Prestávka - sme zatvorení")

	var eventsToGet []string
	for _, event := range events.Items {
		if event == nil {
			continue
		}
		if event.Recurrence != nil && len(event.Recurrence) > 0 {
			eventsToGet = append(eventsToGet, event.Id)
		} else if strings.EqualFold(strings.ToLower(event.Summary), eventTitleToIgnore) {
			continue
		} else {
			result = append(result, newEventFromGoogle(event))
		}
	}
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 20) // number of goroutines to run in parallel
	gatheredEvents := make(chan []models.Event, len(eventsToGet))
	for _, eventID := range eventsToGet {
		wg.Add(1)
		semaphore <- struct{}{}
		go func(eventID string) {
			defer wg.Done()
			events, err := s.getEventInstances(calendarService, eventID, timeMin, timeMax)
			if err != nil {
				fmt.Printf("getting occurenced events instances: %s", err.Error())
				gatheredEvents <- []models.Event{}
			} else {
				gatheredEvents <- events
			}
			<-semaphore
		}(eventID)
	}
	go func() {
		wg.Wait()
		close(gatheredEvents)
	}()
	for events := range gatheredEvents {
		if len(events) > 0 && strings.EqualFold(strings.ToLower(events[0].Title), eventTitleToIgnore) {
			continue
		}
		result = append(result, events...)
	}
	return result, nil
}

func (s *CalendarService) getEventInstances(calendarService *calendar.Service, eventId string, timeMin, timeMax time.Time) ([]models.Event, error) {
	instances, err := calendarService.Events.Instances(s.calendarID, eventId).
		Fields(googleapi.Field("items(id,summary,start,end,recurrence)")).
		TimeMin(timeMin.Format(time.RFC3339)).
		TimeMax(timeMax.Format(time.RFC3339)).
		Do()
	if err != nil {
		fmt.Printf("getting occurenced events instances: %s", err.Error())
		return nil, err
	}
	result := make([]models.Event, 0, len(instances.Items))
	for _, instance := range instances.Items {
		result = append(result, newEventFromGoogle(instance))
	}
	return result, nil
}

func (s *CalendarService) GetEvents(ctx context.Context, timeMin, timeMax time.Time) ([]models.Event, error) {
	key := timeMin.Format(time.DateOnly) + timeMax.Format(time.DateOnly)
	validUntil := s.validUntil.Load()
	// first, try to get the events from the cache
	cached, ok := s.data.Load(key)
	if ok && validUntil > time.Now().Unix() {
		events, ok := cached.(([]models.Event))
		if ok {
			return events, nil
		}

	}

	events, err := s.getEvents(ctx, timeMin, timeMax)
	if err != nil {
		return nil, err
	}

	// cache the fetched events
	s.data.Store(key, events)
	s.validUntil.Store(time.Now().Add(30 * time.Minute).Unix())

	return events, nil
}

func newEventFromGoogle(googleEvent *calendar.Event) models.Event {
	event := models.NewEvent(googleEvent.Summary)
	if googleEvent.Start != nil {
		event.SetStart(googleEvent.Start.DateTime)
	}
	if googleEvent.End != nil {
		event.SetEnd(googleEvent.End.DateTime)
	}
	return event
}
