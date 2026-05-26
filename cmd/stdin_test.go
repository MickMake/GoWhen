package cmd

import (
	"testing"
	"time"

	"GoWhen/cmd/cal"
)

func TestResetPipelineStatePreservesFormatSettings(t *testing.T) {
	cs := Cmds{}
	cs.reparse = true
	cs.last = true
	cs.Error = errForTest{}
	cs.Data.Convert = &cal.Convert{}
	cs.Data.FormatType = "java"
	cs.Data.GoFormat = false
	cs.Data.CppFormat = false
	cs.Data.JavaFormat = true
	cs.Data.SetFromDate(time.Date(2026, time.May, 26, 0, 0, 0, 0, time.UTC))
	cs.Data.SetToDate(time.Date(2026, time.May, 27, 0, 0, 0, 0, time.UTC))
	cs.Data.SetRange(cal.Duration{Time: 24 * time.Hour})
	if err := cs.Data.AddKeepSelector("weekday"); err != nil {
		t.Fatal(err)
	}

	cs.ResetPipelineState()

	if cs.reparse {
		t.Fatal("reparse should be reset")
	}
	if cs.last {
		t.Fatal("last should be reset")
	}
	if cs.Error != nil {
		t.Fatal("Error should be reset")
	}
	if cs.Data.Convert == nil {
		t.Fatal("Convert should be preserved")
	}
	if cs.Data.FormatType != "java" {
		t.Fatalf("FormatType = %q, want java", cs.Data.FormatType)
	}
	if cs.Data.GoFormat || cs.Data.CppFormat || !cs.Data.JavaFormat {
		t.Fatalf("format flags not preserved: go=%v cpp=%v java=%v", cs.Data.GoFormat, cs.Data.CppFormat, cs.Data.JavaFormat)
	}
	if cs.Data.FromDate.Time != nil {
		t.Fatal("FromDate should be cleared")
	}
	if cs.Data.ToDate.Time != nil {
		t.Fatal("ToDate should be cleared")
	}
	if cs.Data.Range != nil {
		t.Fatal("Range should be cleared")
	}
	if !cs.Data.ShouldPrint(time.Date(2026, time.May, 30, 0, 0, 0, 0, time.UTC)) {
		t.Fatal("selectors should be cleared, so Saturday should print")
	}
}

type errForTest struct{}

func (errForTest) Error() string { return "test error" }
