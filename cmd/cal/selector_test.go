package cal

import (
	"testing"
	"time"
)

func TestParseSelectorNormalisesNames(t *testing.T) {
	tests := map[string]string{
		" Month End ":     "month-end",
		"month_end":      "month-end",
		"month   end":    "month-end",
		"last Friday":    "last-friday",
		"third_thursday": "third-thursday",
	}

	for input, want := range tests {
		selector, err := ParseSelector(input)
		if err != nil {
			t.Fatalf("ParseSelector(%q) returned error: %v", input, err)
		}
		if selector.Name != want {
			t.Fatalf("ParseSelector(%q) name = %q, want %q", input, selector.Name, want)
		}
	}
}

func TestParseSelectorUnknown(t *testing.T) {
	if _, err := ParseSelector("bogus selector"); err == nil {
		t.Fatal("expected unknown selector error")
	}
}

func TestSelectorAliases(t *testing.T) {
	tests := []struct {
		name  string
		date  time.Time
		match bool
	}{
		{name: "weekday", date: date(2026, time.January, 5), match: true},
		{name: "weekdays", date: date(2026, time.January, 5), match: true},
		{name: "weekdays", date: date(2026, time.January, 4), match: false},
		{name: "workday", date: date(2026, time.January, 5), match: true},
		{name: "workdays", date: date(2026, time.January, 5), match: true},
		{name: "business day", date: date(2026, time.January, 5), match: true},
		{name: "business-days", date: date(2026, time.January, 5), match: true},
		{name: "weekend", date: date(2026, time.January, 4), match: true},
		{name: "weekends", date: date(2026, time.January, 4), match: true},
		{name: "weekends", date: date(2026, time.January, 5), match: false},
	}

	for _, tt := range tests {
		selector, err := ParseSelector(tt.name)
		if err != nil {
			t.Fatalf("ParseSelector(%q) returned error: %v", tt.name, err)
		}
		if got := selector.Match(tt.date); got != tt.match {
			t.Fatalf("selector %q match = %v, want %v", tt.name, got, tt.match)
		}
	}
}

func TestWeekdaySelectors(t *testing.T) {
	tests := []struct {
		name string
		date time.Time
	}{
		{name: "mon", date: date(2026, time.January, 5)},
		{name: "monday", date: date(2026, time.January, 5)},
		{name: "tue", date: date(2026, time.January, 6)},
		{name: "tuesday", date: date(2026, time.January, 6)},
		{name: "wed", date: date(2026, time.January, 7)},
		{name: "wednesday", date: date(2026, time.January, 7)},
		{name: "thu", date: date(2026, time.January, 8)},
		{name: "thursday", date: date(2026, time.January, 8)},
		{name: "fri", date: date(2026, time.January, 9)},
		{name: "friday", date: date(2026, time.January, 9)},
		{name: "sat", date: date(2026, time.January, 10)},
		{name: "saturday", date: date(2026, time.January, 10)},
		{name: "sun", date: date(2026, time.January, 11)},
		{name: "sunday", date: date(2026, time.January, 11)},
	}

	for _, tt := range tests {
		selector, err := ParseSelector(tt.name)
		if err != nil {
			t.Fatalf("ParseSelector(%q) returned error: %v", tt.name, err)
		}
		if !selector.Match(tt.date) {
			t.Fatalf("selector %q should match %s", tt.name, tt.date)
		}
	}
}

func TestBoundarySelectors(t *testing.T) {
	tests := []struct {
		name  string
		date  time.Time
		match bool
	}{
		{name: "bom", date: date(2026, time.February, 1), match: true},
		{name: "month start", date: date(2026, time.February, 1), match: true},
		{name: "month begin", date: date(2026, time.February, 1), match: true},
		{name: "month start", date: date(2026, time.February, 2), match: false},
		{name: "eom", date: date(2026, time.February, 28), match: true},
		{name: "month end", date: date(2026, time.February, 28), match: true},
		{name: "month end", date: date(2026, time.February, 27), match: false},
		{name: "month end", date: date(2024, time.February, 29), match: true},
		{name: "boy", date: date(2026, time.January, 1), match: true},
		{name: "year start", date: date(2026, time.January, 1), match: true},
		{name: "year begin", date: date(2026, time.January, 1), match: true},
		{name: "year start", date: date(2026, time.January, 2), match: false},
		{name: "eoy", date: date(2026, time.December, 31), match: true},
		{name: "year end", date: date(2026, time.December, 31), match: true},
		{name: "year end", date: date(2026, time.December, 30), match: false},
		{name: "quarter start", date: date(2026, time.January, 1), match: true},
		{name: "quarter begin", date: date(2026, time.April, 1), match: true},
		{name: "quarter start", date: date(2026, time.February, 1), match: false},
		{name: "quarter end", date: date(2026, time.March, 31), match: true},
		{name: "quarter end", date: date(2026, time.June, 30), match: true},
		{name: "quarter end", date: date(2026, time.September, 30), match: true},
		{name: "quarter end", date: date(2026, time.December, 31), match: true},
		{name: "quarter end", date: date(2026, time.June, 29), match: false},
	}

	for _, tt := range tests {
		selector, err := ParseSelector(tt.name)
		if err != nil {
			t.Fatalf("ParseSelector(%q) returned error: %v", tt.name, err)
		}
		if got := selector.Match(tt.date); got != tt.match {
			t.Fatalf("selector %q match = %v, want %v", tt.name, got, tt.match)
		}
	}
}

func TestDayOfMonthSelector(t *testing.T) {
	selector, err := ParseSelector("day-15")
	if err != nil {
		t.Fatalf("ParseSelector(day-15) returned error: %v", err)
	}
	if !selector.Match(date(2026, time.January, 15)) {
		t.Fatal("day-15 should match January 15")
	}
	if selector.Match(date(2026, time.January, 16)) {
		t.Fatal("day-15 should not match January 16")
	}
}

func TestDayOfMonthSelectorNoMatchForMissingMonthDay(t *testing.T) {
	selector, err := ParseSelector("day-31")
	if err != nil {
		t.Fatalf("ParseSelector(day-31) returned error: %v", err)
	}
	if selector.Match(date(2026, time.February, 28)) {
		t.Fatal("day-31 should not match February 28")
	}
}

func TestInvalidDayOfMonthSelector(t *testing.T) {
	for _, name := range []string{"day-0", "day-32", "day-nope"} {
		if _, err := ParseSelector(name); err == nil {
			t.Fatalf("ParseSelector(%q) expected error", name)
		}
	}
}

func TestOrdinalWeekdaySelectors(t *testing.T) {
	tests := []struct {
		name  string
		date  time.Time
		match bool
	}{
		{name: "first monday", date: date(2026, time.January, 5), match: true},
		{name: "first-mon", date: date(2026, time.January, 5), match: true},
		{name: "first monday", date: date(2026, time.January, 12), match: false},
		{name: "second-tue", date: date(2026, time.January, 13), match: true},
		{name: "second-tuesday", date: date(2026, time.January, 13), match: true},
		{name: "third-thu", date: date(2026, time.January, 15), match: true},
		{name: "third-thursday", date: date(2026, time.January, 15), match: true},
		{name: "fourth fri", date: date(2026, time.January, 23), match: true},
		{name: "fourth friday", date: date(2026, time.January, 23), match: true},
		{name: "last fri", date: date(2026, time.January, 30), match: true},
		{name: "last friday", date: date(2026, time.January, 30), match: true},
		{name: "last friday", date: date(2026, time.January, 23), match: false},
		{name: "last monday", date: date(2024, time.February, 26), match: true},
		{name: "last monday", date: date(2024, time.February, 19), match: false},
	}

	for _, tt := range tests {
		selector, err := ParseSelector(tt.name)
		if err != nil {
			t.Fatalf("ParseSelector(%q) returned error: %v", tt.name, err)
		}
		if got := selector.Match(tt.date); got != tt.match {
			t.Fatalf("selector %q match = %v, want %v", tt.name, got, tt.match)
		}
	}
}

func date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
