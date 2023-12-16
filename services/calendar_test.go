package services_test

import (
	"context"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"jirku.sk/mcmamina/services"
)

func TestCalendar(t *testing.T) {
	var env map[string]string
	env, err := godotenv.Read("../.env")
	if err != nil {
		t.Errorf("Error loading .env file %v", err)
	}
	calendarService := services.NewCalendarService(env["GOOGLE_API_KEY"], env["GOOGLE_CALENDAR_ID"])
	events, err := calendarService.GetEvents(context.Background(), time.Now(), time.Now().AddDate(0, 0, 5))
	if err != nil {
		t.Errorf("Failed to get events: %v", err)
	}
	if len(events) == 0 {
		t.Errorf("Expected at least one event")
	}
}
