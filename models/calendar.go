package models

import "time"

type Event struct {
	Title string
	Start time.Time
	End   time.Time
}

type Day struct {
	Date   time.Time
	Events []Event
}

func NewDay(date time.Time) Day {
	return Day{
		Date:   date,
		Events: []Event{},
	}
}

func GetNumberOfDaysAfter(date time.Time, count int) []Day {
	days := make([]Day, count)
	for i := 0; i < count; i++ {
		days[i] = NewDay(date.AddDate(0, 0, i))
	}
	return days
}

func (d *Day) GetTitle() string {
	dayInEnglish := d.Date.Format("Monday")

	dayMap := map[string]string{
		"Monday":    "Pondelok",
		"Tuesday":   "Utorok",
		"Wednesday": "Streda",
		"Thursday":  "Štvrtok",
		"Friday":    "Piatok",
		"Saturday":  "Sobota",
		"Sunday":    "Nedeľa",
	}

	dayInSlovak, ok := dayMap[dayInEnglish]
	if !ok {
		return "Unknown day"
	}

	return dayInSlovak
}

func (d *Day) GetAbbr() string {
	dayInEnglish := d.Date.Format("Monday")

	dayMap := map[string]string{
		"Monday":    "Po",
		"Tuesday":   "Ut",
		"Wednesday": "St",
		"Thursday":  "Št",
		"Friday":    "Pi",
		"Saturday":  "So",
		"Sunday":    "Ne",
	}

	dayInSlovak, ok := dayMap[dayInEnglish]
	if !ok {
		return "--"
	}

	return dayInSlovak
}

func (d *Day) AddEvent(event Event) {
	d.Events = append(d.Events, event)
}

func (d *Day) Is(t time.Time) bool {
	return d.Date.Year() == t.Year() && d.Date.Month() == t.Month() && d.Date.Day() == t.Day()
}

func (d *Day) IsWeekend() bool {
	return d.Date.Weekday() == time.Saturday || d.Date.Weekday() == time.Sunday
}

func (d *Day) IsBefore(t time.Time) bool {
	return d.Date.Before(t)
}

func (d *Day) HasEvents() bool {
	return len(d.Events) > 0
}

func (d *Day) GetDay() string {
	return d.Date.Format("2")
}
