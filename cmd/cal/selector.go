package cal

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Selector struct {
	Name string

	Weekday *time.Weekday
	Weekend *bool

	MonthStart   bool
	MonthEnd     bool
	YearStart    bool
	YearEnd      bool
	QuarterStart bool
	QuarterEnd   bool

	DayOfMonth int

	Ordinal        int
	OrdinalWeekday *time.Weekday
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
	selector := Selector{Name: normaliseSelectorName(name)}

	if selector.Name == "" {
		return selector, fmt.Errorf("unknown selector %q", name)
	}

	switch selector.Name {
	case "weekday", "weekdays", "workday", "workdays", "business-day", "business-days":
		weekend := false
		selector.Weekend = &weekend
	case "weekend", "weekends":
		weekend := true
		selector.Weekend = &weekend
	case "bom", "month-start", "month-begin":
		selector.MonthStart = true
	case "eom", "month-end":
		selector.MonthEnd = true
	case "boy", "year-start", "year-begin":
		selector.YearStart = true
	case "eoy", "year-end":
		selector.YearEnd = true
	case "quarter-start", "quarter-begin":
		selector.QuarterStart = true
	case "quarter-end":
		selector.QuarterEnd = true
	default:
		if day, ok := parseWeekdayName(selector.Name); ok {
			selector.Weekday = &day
			return selector, nil
		}

		if day, ok, err := parseDayOfMonthSelector(selector.Name); ok || err != nil {
			if err != nil {
				return selector, err
			}
			selector.DayOfMonth = day
			return selector, nil
		}

		if ordinal, day, ok := parseOrdinalWeekdaySelector(selector.Name); ok {
			selector.Ordinal = ordinal
			selector.OrdinalWeekday = &day
			return selector, nil
		}

		return selector, fmt.Errorf("unknown selector %q", name)
	}

	return selector, nil
}

func normaliseSelectorName(name string) string {
	name = strings.ToLower(strings.TrimSpace(name))
	name = strings.ReplaceAll(name, "_", " ")
	name = strings.ReplaceAll(name, "-", " ")
	return strings.Join(strings.Fields(name), "-")
}

func parseWeekdayName(name string) (time.Weekday, bool) {
	switch name {
	case "mon", "monday":
		return time.Monday, true
	case "tue", "tuesday":
		return time.Tuesday, true
	case "wed", "wednesday":
		return time.Wednesday, true
	case "thu", "thursday":
		return time.Thursday, true
	case "fri", "friday":
		return time.Friday, true
	case "sat", "saturday":
		return time.Saturday, true
	case "sun", "sunday":
		return time.Sunday, true
	default:
		return time.Sunday, false
	}
}

func parseDayOfMonthSelector(name string) (int, bool, error) {
	if !strings.HasPrefix(name, "day-") {
		return 0, false, nil
	}

	value := strings.TrimPrefix(name, "day-")
	day, err := strconv.Atoi(value)
	if err != nil || day < 1 || day > 31 {
		return 0, true, fmt.Errorf("invalid day selector %q", name)
	}

	return day, true, nil
}

func parseOrdinalWeekdaySelector(name string) (int, time.Weekday, bool) {
	ordinalName, weekdayName, ok := strings.Cut(name, "-")
	if !ok {
		return 0, time.Sunday, false
	}

	ordinal, ok := parseOrdinalName(ordinalName)
	if !ok {
		return 0, time.Sunday, false
	}

	weekday, ok := parseWeekdayName(weekdayName)
	if !ok {
		return 0, time.Sunday, false
	}

	return ordinal, weekday, true
}

func parseOrdinalName(name string) (int, bool) {
	switch name {
	case "first":
		return 1, true
	case "second":
		return 2, true
	case "third":
		return 3, true
	case "fourth":
		return 4, true
	case "last":
		return -1, true
	default:
		return 0, false
	}
}

func (s Selector) Match(t time.Time) bool {
	if s.Weekday != nil {
		return t.Weekday() == *s.Weekday
	}

	if s.Weekend != nil {
		isWeekend := t.Weekday() == time.Saturday || t.Weekday() == time.Sunday
		return isWeekend == *s.Weekend
	}

	if s.MonthStart {
		return t.Day() == 1
	}

	if s.MonthEnd {
		return t.AddDate(0, 0, 1).Day() == 1
	}

	if s.YearStart {
		return t.Month() == time.January && t.Day() == 1
	}

	if s.YearEnd {
		return t.Month() == time.December && t.Day() == 31
	}

	if s.QuarterStart {
		return t.Day() == 1 && isQuarterStartMonth(t.Month())
	}

	if s.QuarterEnd {
		return isQuarterEndMonth(t.Month()) && t.AddDate(0, 0, 1).Day() == 1
	}

	if s.DayOfMonth > 0 {
		return t.Day() == s.DayOfMonth
	}

	if s.OrdinalWeekday != nil {
		return matchOrdinalWeekday(t, s.Ordinal, *s.OrdinalWeekday)
	}

	return false
}

func isQuarterStartMonth(month time.Month) bool {
	switch month {
	case time.January, time.April, time.July, time.October:
		return true
	default:
		return false
	}
}

func isQuarterEndMonth(month time.Month) bool {
	switch month {
	case time.March, time.June, time.September, time.December:
		return true
	default:
		return false
	}
}

func matchOrdinalWeekday(t time.Time, ordinal int, weekday time.Weekday) bool {
	if t.Weekday() != weekday {
		return false
	}

	if ordinal == -1 {
		return t.AddDate(0, 0, 7).Month() != t.Month()
	}

	firstDay := ((ordinal - 1) * 7) + 1
	lastDay := ordinal * 7
	return t.Day() >= firstDay && t.Day() <= lastDay
}
