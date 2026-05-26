package cmd

import (
	"bytes"
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
