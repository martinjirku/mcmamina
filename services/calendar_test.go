package services_test

import (
	"context"
	"testing"
	"time"

	"jirku.sk/mcmamina/services"
)

func TestCalendar(t *testing.T) {
	calendarService := services.NewCalendarService("random", "randomID")
	events, err := calendarService.GetEvents(context.Background(), time.Now(), time.Now().AddDate(0, 0, 5))
	if err != nil {
		t.Errorf("Failed to get events: %v", err)
	}
	if len(events) == 0 {
		t.Errorf("Expected at least one event")
	}
}
