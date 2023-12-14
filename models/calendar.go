package models

import "time"

func IsSameDay(d, t time.Time) bool {
	return d.Year() == t.Year() && d.Month() == t.Month() && d.Day() == t.Day()
}

type Event struct {
	Title string
	Start time.Time
	End   time.Time
}

func NewEvent(title string) Event {
	return Event{
		Title: title,
	}
}

func (e *Event) SetStart(t string) *Event {
	start, err := time.Parse(time.RFC3339, t)
	if err != nil {
		return e
	}
	e.Start = start
	return e
}

func (e *Event) SetEnd(t string) *Event {
	end, err := time.Parse(time.RFC3339, t)
	if err != nil {
		return e
	}
	e.End = end
	return e
}

func (e *Event) RenderEventRange(currentDay time.Time) string {
	var Label string
	if !IsSameDay(e.Start, currentDay) {
		Label += e.Start.Format("02. 01. 2006 ")
	}
	Label += e.Start.Format("15:04")
	Label += " - "
	if !IsSameDay(e.End, currentDay) {
		Label += e.End.Format("02. 01. 2006 ")
	}
	Label += e.End.Format("15:04")
	return Label
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
	return IsSameDay(d.Date, t)
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

func (d *Day) GetFormatedDate() string {
	return d.Date.Format("02. 01. 2006")
}

func (d *Day) GetDateIdentifier() string {
	return "day-" + d.Date.Format("2006-01-02")
}
