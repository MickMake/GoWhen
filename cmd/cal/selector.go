package cal

import (
	"fmt"
	"strings"
	"time"
)

type Selector struct {
	Name    string
	Weekday *time.Weekday
	Weekend *bool
}

var keepSelectors []Selector
var dropSelectors []Selector

func (d *Data) AddKeepSelector(name string) error {
	selector, err := ParseSelector(name)
	if err != nil {
		return err
	}
	keepSelectors = append(keepSelectors, selector)
	return nil
}

func (d *Data) AddDropSelector(name string) error {
	selector, err := ParseSelector(name)
	if err != nil {
		return err
	}
	dropSelectors = append(dropSelectors, selector)
	return nil
}

func (d *Data) ClearSelectors() {
	keepSelectors = nil
	dropSelectors = nil
}

func (d *Data) ShouldPrint(t time.Time) bool {
	if len(keepSelectors) > 0 {
		matched := false
		for _, selector := range keepSelectors {
			if selector.Match(t) {
				matched = true
				break
			}
		}
		if !matched {
			return false
		}
	}

	for _, selector := range dropSelectors {
		if selector.Match(t) {
			return false
		}
	}

	return true
}

func ParseSelector(name string) (Selector, error) {
	selector := Selector{Name: strings.ToLower(strings.TrimSpace(name))}

	switch selector.Name {
	case "weekday":
		weekend := false
		selector.Weekend = &weekend
	case "weekend":
		weekend := true
		selector.Weekend = &weekend
	case "mon", "monday":
		day := time.Monday
		selector.Weekday = &day
	case "tue", "tuesday":
		day := time.Tuesday
		selector.Weekday = &day
	case "wed", "wednesday":
		day := time.Wednesday
		selector.Weekday = &day
	case "thu", "thursday":
		day := time.Thursday
		selector.Weekday = &day
	case "fri", "friday":
		day := time.Friday
		selector.Weekday = &day
	case "sat", "saturday":
		day := time.Saturday
		selector.Weekday = &day
	case "sun", "sunday":
		day := time.Sunday
		selector.Weekday = &day
	default:
		return selector, fmt.Errorf("unknown selector %q", name)
	}

	return selector, nil
}

func (s Selector) Match(t time.Time) bool {
	if s.Weekday != nil {
		return t.Weekday() == *s.Weekday
	}

	if s.Weekend != nil {
		isWeekend := t.Weekday() == time.Saturday || t.Weekday() == time.Sunday
		return isWeekend == *s.Weekend
	}

	return false
}
