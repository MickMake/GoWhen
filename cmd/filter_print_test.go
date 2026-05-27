package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"testing"
	"time"

	"GoWhen/cmd/cal"
)

func TestFilteredPrintRangeKeepsMonthEnd(t *testing.T) {
	cs := Cmds{}
	cs.Data.ClearSelectors()
	t.Cleanup(cs.Data.ClearSelectors)

	cs.Data.Format = "2006-01-02"
	cs.Data.SetFromDate(time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC))
	cs.Data.SetToDate(time.Date(2026, time.April, 1, 0, 0, 0, 0, time.UTC))
	cs.Data.SetRange(cal.Duration{Time: 24 * time.Hour})
	if err := cs.Data.AddKeepSelector("month end"); err != nil {
		t.Fatal(err)
	}

	got := captureStdout(t, cs.FilteredPrintRange)
	want := "2026-01-31\n2026-02-28\n2026-03-31\n"
	if got != want {
		t.Fatalf("FilteredPrintRange output = %q, want %q", got, want)
	}
}

func TestFilteredPrintRangeDropsWeekends(t *testing.T) {
	cs := Cmds{}
	cs.Data.ClearSelectors()
	t.Cleanup(cs.Data.ClearSelectors)

	cs.Data.Format = "2006-01-02"
	cs.Data.SetFromDate(time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC))
	cs.Data.SetToDate(time.Date(2026, time.January, 8, 0, 0, 0, 0, time.UTC))
	cs.Data.SetRange(cal.Duration{Time: 24 * time.Hour})
	if err := cs.Data.AddDropSelector("weekends"); err != nil {
		t.Fatal(err)
	}

	got := captureStdout(t, cs.FilteredPrintRange)
	want := "2026-01-01\n2026-01-02\n2026-01-05\n2026-01-06\n2026-01-07\n"
	if got != want {
		t.Fatalf("FilteredPrintRange output = %q, want %q", got, want)
	}
}

func TestFilteredPrintRangeExcludesEndDate(t *testing.T) {
	cs := Cmds{}
	cs.Data.ClearSelectors()
	t.Cleanup(cs.Data.ClearSelectors)

	cs.Data.Format = "2006-01-02"
	cs.Data.SetFromDate(time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC))
	cs.Data.SetToDate(time.Date(2026, time.January, 3, 0, 0, 0, 0, time.UTC))
	cs.Data.SetRange(cal.Duration{Time: 24 * time.Hour})

	got := captureStdout(t, cs.FilteredPrintRange)
	want := "2026-01-01\n2026-01-02\n"
	if got != want {
		t.Fatalf("FilteredPrintRange output = %q, want %q", got, want)
	}
}

func TestFilteredPrintReverseRangeKeepsMonthEnd(t *testing.T) {
	cs := Cmds{}
	cs.Data.ClearSelectors()
	t.Cleanup(cs.Data.ClearSelectors)

	cs.Data.Format = "2006-01-02"
	cs.Data.SetFromDate(time.Date(2026, time.April, 1, 0, 0, 0, 0, time.UTC))
	cs.Data.SetToDate(time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC))
	cs.Data.SetRange(cal.Duration{Time: 24 * time.Hour})
	if err := cs.Data.AddKeepSelector("month end"); err != nil {
		t.Fatal(err)
	}

	got := captureStdout(t, cs.FilteredPrintRange)
	want := "2026-03-31\n2026-02-28\n2026-01-31\n"
	if got != want {
		t.Fatalf("FilteredPrintRange output = %q, want %q", got, want)
	}
}

func TestFilteredPrintTimeSuppressesNonMatchingDate(t *testing.T) {
	cs := Cmds{}
	cs.Data.ClearSelectors()
	t.Cleanup(cs.Data.ClearSelectors)

	cs.Data.Format = "2006-01-02"
	if err := cs.Data.AddKeepSelector("weekday"); err != nil {
		t.Fatal(err)
	}

	got := captureStdout(t, func() {
		cs.FilteredPrintTime(time.Date(2026, time.May, 30, 0, 0, 0, 0, time.UTC))
	})
	if got != "" {
		t.Fatalf("FilteredPrintTime output = %q, want empty output", got)
	}
}

func TestPrintFormattedTimeNamedFormats(t *testing.T) {
	tm := time.Date(2026, time.May, 26, 15, 4, 5, 123456789, time.UTC)
	tests := []struct {
		format string
		want   string
	}{
		{format: "epoch", want: "1787756645\n"},
		{format: "unix", want: "1787756645\n"},
		{format: "unix-ms", want: "1787756645123\n"},
		{format: "unix-us", want: "1787756645123456\n"},
		{format: "unix-ns", want: "1787756645123456789\n"},
		{format: "iso", want: "2026-05-26T15:04:05Z\n"},
		{format: "date", want: "2026-05-26\n"},
		{format: "time", want: "15:04:05\n"},
		{format: "datetime", want: "2026-05-26 15:04:05\n"},
	}

	for _, tt := range tests {
		t.Run(tt.format, func(t *testing.T) {
			cs := Cmds{}
			cs.Data.Format = tt.format
			got := captureStdout(t, func() {
				cs.PrintFormattedTime(tm)
			})
			if got != tt.want {
				t.Fatalf("PrintFormattedTime(%q) = %q, want %q", tt.format, got, tt.want)
			}
		})
	}
}

func TestPrintFormattedTimeStructuredFormats(t *testing.T) {
	tm := time.Date(2026, time.May, 26, 15, 4, 5, 123456789, time.UTC)
	tests := []struct {
		format string
		want   string
	}{
		{format: "csv", want: "date,epoch,weekday,year,month,day,week\n2026-05-26T15:04:05Z,1787756645,Tuesday,2026,5,26,22\n"},
		{format: "tsv", want: "date\tepoch\tweekday\tyear\tmonth\tday\tweek\n2026-05-26T15:04:05Z\t1787756645\tTuesday\t2026\t5\t26\t22\n"},
	}

	for _, tt := range tests {
		t.Run(tt.format, func(t *testing.T) {
			cs := Cmds{}
			cs.Data.Format = tt.format
			got := captureStdout(t, func() {
				cs.PrintFormattedTime(tm)
			})
			if got != tt.want {
				t.Fatalf("PrintFormattedTime(%q) = %q, want %q", tt.format, got, tt.want)
			}
		})
	}
}

func TestPrintFormattedTimeJSONL(t *testing.T) {
	cs := Cmds{}
	cs.Data.Format = "jsonl"
	tm := time.Date(2026, time.May, 26, 15, 4, 5, 123456789, time.UTC)

	got := captureStdout(t, func() {
		cs.PrintFormattedTime(tm)
	})

	var decoded map[string]interface{}
	if err := json.Unmarshal([]byte(got), &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded["date"] != "2026-05-26T15:04:05Z" || decoded["weekday"] != "Tuesday" {
		t.Fatalf("unexpected JSONL string fields: %#v", decoded)
	}
	if decoded["epoch"] != float64(1787756645) || decoded["year"] != float64(2026) || decoded["month"] != float64(5) || decoded["day"] != float64(26) || decoded["week"] != float64(22) {
		t.Fatalf("unexpected JSONL numeric fields: %#v", decoded)
	}
}

func TestFilteredPrintRangeUsesNamedFormats(t *testing.T) {
	cs := Cmds{}
	cs.Data.ClearSelectors()
	t.Cleanup(cs.Data.ClearSelectors)

	cs.Data.Format = "date"
	cs.Data.SetFromDate(time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC))
	cs.Data.SetToDate(time.Date(2026, time.January, 3, 0, 0, 0, 0, time.UTC))
	cs.Data.SetRange(cal.Duration{Time: 24 * time.Hour})

	got := captureStdout(t, cs.FilteredPrintRange)
	want := "2026-01-01\n2026-01-02\n"
	if got != want {
		t.Fatalf("FilteredPrintRange output = %q, want %q", got, want)
	}
}

func TestFilteredPrintRangeUsesStructuredFormats(t *testing.T) {
	cs := Cmds{}
	cs.Data.ClearSelectors()
	t.Cleanup(cs.Data.ClearSelectors)

	cs.Data.Format = "csv"
	cs.Data.SetFromDate(time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC))
	cs.Data.SetToDate(time.Date(2026, time.January, 3, 0, 0, 0, 0, time.UTC))
	cs.Data.SetRange(cal.Duration{Time: 24 * time.Hour})

	got := captureStdout(t, cs.FilteredPrintRange)
	want := "date,epoch,weekday,year,month,day,week\n2026-01-01T00:00:00Z,1767225600,Thursday,2026,1,1,1\n2026-01-02T00:00:00Z,1767312000,Friday,2026,1,2,1\n"
	if got != want {
		t.Fatalf("FilteredPrintRange output = %q, want %q", got, want)
	}
}

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w

	fn()

	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	os.Stdout = old

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatal(err)
	}
	if err := r.Close(); err != nil {
		t.Fatal(err)
	}

	return buf.String()
}
