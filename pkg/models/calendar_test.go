package models_test

import (
	"testing"
	"time"

	"jirku.sk/mcmamina/pkg/models"
)

func getTime(s string) time.Time {
	t, _ := time.Parse(time.DateOnly, s)
	return t
}

func TestDay(t *testing.T) {
	day := models.NewDay(getTime("2023-12-01"))
	if day.GetTitle() != "Piatok" {
		t.Errorf("Expected %v, got %v", "Piatok", day.GetTitle())
	}
	if day.GetAbbr() != "Pi" {
		t.Errorf("Expected %v, got %v", "Pi", day.GetAbbr())
	}
}

func TestGetNumberOfDaysAfter(t *testing.T) {
	days := models.GetNumberOfDaysAfter(getTime("2023-12-01"), 2)
	if len(days) != 2 {
		t.Errorf("Expected %v, got %v", 2, len(days))
	}
	if days[0].GetTitle() != "Piatok" {
		t.Errorf("Expected %v, got %v", "Štvrtok", days[0].GetTitle())
	}
	if days[1].GetTitle() != "Sobota" {
		t.Errorf("Expected %v, got %v", "Piatok", days[1].GetTitle())
	}
	if days[0].GetAbbr() != "Pi" {
		t.Errorf("Expected %v, got %v", "Št", days[0].GetAbbr())
	}
	if days[1].GetAbbr() != "So" {
		t.Errorf("Expected %v, got %v", "Št", days[0].GetAbbr())
	}
}
