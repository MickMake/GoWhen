package cal

import (
	"testing"
	"time"
)

func TestParseDurationComposite(t *testing.T) {
	got, err := ParseDuration("1y 2M 3w 4d 5h 6m 7s")
	if err != nil {
		t.Fatal(err)
	}

	if got.Years != 1 {
		t.Fatalf("Years = %d, want 1", got.Years)
	}
	if got.Months != 2 {
		t.Fatalf("Months = %d, want 2", got.Months)
	}

	want := (3 * 7 * 24 * time.Hour) + (4 * 24 * time.Hour) + (5 * time.Hour) + (6 * time.Minute) + (7 * time.Second)
	if got.Time != want {
		t.Fatalf("Time = %s, want %s", got.Time, want)
	}
}

func TestParseDurationDecimalWeeksAndDays(t *testing.T) {
	got, err := ParseDuration("1.5w .5d")
	if err != nil {
		t.Fatal(err)
	}

	want := (time.Duration(float64(int64(time.Hour)*168) * 1.5)) + (12 * time.Hour)
	if got.Time != want {
		t.Fatalf("Time = %s, want %s", got.Time, want)
	}
}

func TestParseDurationNegativeValues(t *testing.T) {
	got, err := ParseDuration("-1y -2M -3w -4d -5h")
	if err != nil {
		t.Fatal(err)
	}

	if got.Years != -1 {
		t.Fatalf("Years = %d, want -1", got.Years)
	}
	if got.Months != -2 {
		t.Fatalf("Months = %d, want -2", got.Months)
	}

	want := -(3 * 7 * 24 * time.Hour) - (4 * 24 * time.Hour) - (5 * time.Hour)
	if got.Time != want {
		t.Fatalf("Time = %s, want %s", got.Time, want)
	}
}

func TestParseDurationPlusValues(t *testing.T) {
	got, err := ParseDuration("+1y +2M +7d")
	if err != nil {
		t.Fatal(err)
	}

	if got.Years != 1 {
		t.Fatalf("Years = %d, want 1", got.Years)
	}
	if got.Months != 2 {
		t.Fatalf("Months = %d, want 2", got.Months)
	}
	if got.Time != 7*24*time.Hour {
		t.Fatalf("Time = %s, want %s", got.Time, 7*24*time.Hour)
	}
}

func TestParseDurationEmptyString(t *testing.T) {
	got, err := ParseDuration("")
	if err != nil {
		t.Fatal(err)
	}
	if got.Years != 0 || got.Months != 0 || got.Time != 0 {
		t.Fatalf("ParseDuration empty = %+v, want zero value", got)
	}
}

func TestParseDurationInvalidToken(t *testing.T) {
	if _, err := ParseDuration("bananas"); err == nil {
		t.Fatal("expected invalid duration error")
	}
}

func TestParseDurationInvalidMiddleTokenReturnsError(t *testing.T) {
	if _, err := ParseDuration("1d bananas 2h"); err == nil {
		t.Fatal("expected invalid duration error")
	}
}

func TestDateDiffString(t *testing.T) {
	diff := DateDiff(
		time.Date(2026, time.January, 1, 1, 2, 3, 0, time.UTC),
		time.Date(2027, time.March, 4, 5, 6, 7, 0, time.UTC),
	)

	if diff.String() != "1y 2M 3d 4h 4m 4s" {
		t.Fatalf("Diff.String() = %q", diff.String())
	}
}

func TestStrToFormatKnownAliases(t *testing.T) {
	tests := map[string]string{
		".":           "",
		"simple":      "2006-01-02T15:04:05",
		"RFC3339":     "2006-01-02T15:04:05Z07:00",
		"RFC3339Nano": "2006-01-02T15:04:05.999999999Z07:00",
		"epoch":       "epoch",
		"week":        "week",
	}

	for input, want := range tests {
		if got := StrToFormat(input); got != want {
			t.Fatalf("StrToFormat(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestStrToDateEpoch(t *testing.T) {
	got := StrToDate("epoch")
	if got == nil {
		t.Fatal("StrToDate(epoch) returned nil")
	}

	want := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	if !got.Equal(want) {
		t.Fatalf("StrToDate(epoch) = %s, want %s", got, want)
	}
}

func TestStrToDateRelativeValues(t *testing.T) {
	tests := []struct {
		name string
		want time.Duration
	}{
		{name: "tomorrow", want: 24 * time.Hour},
		{name: "yesterday", want: -24 * time.Hour},
		{name: "next-week", want: 168 * time.Hour},
		{name: "last-week", want: -168 * time.Hour},
	}

	for _, tt := range tests {
		before := time.Now()
		got := StrToDate(tt.name)
		after := time.Now()
		if got == nil {
			t.Fatalf("StrToDate(%q) returned nil", tt.name)
		}

		min := before.Add(tt.want).Add(-time.Second)
		max := after.Add(tt.want).Add(time.Second)
		if got.Before(min) || got.After(max) {
			t.Fatalf("StrToDate(%q) = %s, want between %s and %s", tt.name, got, min, max)
		}
	}
}

func TestStrToDateUnknownReturnsNil(t *testing.T) {
	if got := StrToDate("not-a-date-keyword"); got != nil {
		t.Fatalf("StrToDate unknown = %s, want nil", got)
	}
}
