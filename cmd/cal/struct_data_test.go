package cal

import (
	"path/filepath"
	"testing"
	"time"
)

func TestConvertFormatPreservesGoLayout(t *testing.T) {
	d := Data{GoFormat: true}

	d.ConvertFormat("2006-01-02")

	if d.Format != "2006-01-02" {
		t.Fatalf("expected Go layout to be preserved, got %q", d.Format)
	}
}

func TestParseDateStringUsesConvertedGoLayout(t *testing.T) {
	d := Data{GoFormat: true}

	got, err := d.ParseDateString("2006-01-02", "2026-05-26")
	if err != nil {
		t.Fatal(err)
	}

	want := time.Date(2026, time.May, 26, 0, 0, 0, 0, time.UTC)
	if !got.Equal(want) {
		t.Fatalf("expected %s, got %s", want, got)
	}
}

func TestParseDateStringUsesConvertedJavaLayout(t *testing.T) {
	convert, err := ReadConvert(filepath.Join(t.TempDir(), "convert.json"))
	if err != nil {
		t.Fatal(err)
	}

	d := Data{JavaFormat: true, Convert: convert}

	got, err := d.ParseDateString("yyyy-MM-dd", "2026-05-26")
	if err != nil {
		t.Fatal(err)
	}

	want := time.Date(2026, time.May, 26, 0, 0, 0, 0, time.UTC)
	if !got.Equal(want) {
		t.Fatalf("expected %s, got %s", want, got)
	}
}

func TestIsDateBeforeReturnsTrueAndNilError(t *testing.T) {
	d := Data{GoFormat: true}
	d.SetFromDate(time.Date(2026, time.May, 26, 0, 0, 0, 0, time.UTC))

	got, err := d.IsDateBefore("2006-01-02", "2026-05-27")
	if err != nil {
		t.Fatal(err)
	}
	if !got {
		t.Fatal("expected date to be before comparison date")
	}
}

func TestIsDateAfterReturnsTrueAndNilError(t *testing.T) {
	d := Data{GoFormat: true}
	d.SetFromDate(time.Date(2026, time.May, 26, 0, 0, 0, 0, time.UTC))

	got, err := d.IsDateAfter("2006-01-02", "2026-05-25")
	if err != nil {
		t.Fatal(err)
	}
	if !got {
		t.Fatal("expected date to be after comparison date")
	}
}

func TestIsDateBeforeReturnsParseError(t *testing.T) {
	d := Data{GoFormat: true}
	d.SetFromDate(time.Date(2026, time.May, 26, 0, 0, 0, 0, time.UTC))

	got, err := d.IsDateBefore("2006-01-02", "potato")
	if err == nil {
		t.Fatal("expected parse error")
	}
	if got {
		t.Fatal("expected false result when parsing fails")
	}
}

func TestIsDateAfterReturnsParseError(t *testing.T) {
	d := Data{GoFormat: true}
	d.SetFromDate(time.Date(2026, time.May, 26, 0, 0, 0, 0, time.UTC))

	got, err := d.IsDateAfter("2006-01-02", "potato")
	if err == nil {
		t.Fatal("expected parse error")
	}
	if got {
		t.Fatal("expected false result when parsing fails")
	}
}
