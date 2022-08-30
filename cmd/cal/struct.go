package cal

import (
	"github.com/jinzhu/now"
	"time"
)


// New returns new Calendar pointer.
func New(t time.Time) *Calendar {
	return &Calendar{
		Now: now.New(t),
	}
}

// Now returns Calendar
func Now() *Calendar {
	return &Calendar{
		Now: now.New(time.Now()),
	}
}

// Next returns time.Time collection to represent next week.
func (w Week) Next() (nextWeek Week) {
	for _, t := range w {
		nextWeek = append(nextWeek, t.AddDate(0, 0, 7))
	}
	return
}

// Previous returns time.Time collection to represent previous week.
func (w Week) Previous() (previousWeek Week) {
	for _, t := range w {
		previousWeek = append(previousWeek, t.AddDate(0, 0, -7))
	}
	return
}

// Calendar have now.Now data.
type Calendar struct {
	Now *now.Now
}

// Next sets new *now.Now for next month.
func (c *Calendar) Next() {
	newDate := c.Now.BeginningOfMonth().AddDate(0, 1, 0)
	c.Now = now.New(newDate)
}

// Previous sets new *now.Now for previous month.
func (c *Calendar) Previous() {
	newDate := c.Now.BeginningOfMonth().AddDate(0, -1, 0)
	c.Now = now.New(newDate)
}

// NextCalendar returns next Calendar.
func (c *Calendar) NextCalendar() (nextCalendar *Calendar) {
	newDate := c.Now.BeginningOfMonth().AddDate(0, 1, 0)
	nextCalendar = &Calendar{
		Now: now.New(newDate),
	}
	return
}

// PreviousCalendar returns previous Calendar.
func (c *Calendar) PreviousCalendar() (previousCalendar *Calendar) {
	newDate := c.Now.BeginningOfMonth().AddDate(0, -1, 0)
	previousCalendar = &Calendar{
		Now: now.New(newDate),
	}
	return
}

// NextYearCalendar returns next Calendar.
func (c *Calendar) NextYearCalendar() (nextCalendar *Calendar) {
	newDate := c.Now.BeginningOfMonth().AddDate(1, 0, 0)
	nextCalendar = &Calendar{
		Now: now.New(newDate),
	}
	return
}

// PreviousYearCalendar returns previous Calendar.
func (c *Calendar) PreviousYearCalendar() (previousCalendar *Calendar) {
	newDate := c.Now.BeginningOfMonth().AddDate(-1, 0, 0)
	previousCalendar = &Calendar{
		Now: now.New(newDate),
	}
	return
}
