package services_test

import (
	"context"
	"testing"
	"time"

	"jirku.sk/mcmamina/services"
)

func TestCalendar(t *testing.T) {
	calendarService := services.NewCalendarService("AIzaSyAkRcuTDqhS6FYZvOASB8gXNWhQ_grK3lg", "n4bgt6kl18u5ueku1g38f5kic8@group.calendar.google.com")
	// calendarService := services.NewCalendarService("AIzaSyAkRcuTDqhS6FYZvOASB8gXNWhQ_grK3lg", "450bfae1def53fcc04f16ae1a9787c3901090ea5b3e87aaab4643e50745bb91d@group.calendar.google.com")
	events, err := calendarService.GetEvents(context.Background(), time.Now(), time.Now().AddDate(0, 0, 5))
	if err != nil {
		t.Errorf("Failed to get events: %v", err)
	}
	if len(events) == 0 {
		t.Errorf("Expected at least one event")
	}
}
