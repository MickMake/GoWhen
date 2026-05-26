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

func TestDateParseSetsFromDate(t *testing.T) {
	d := Data{GoFormat: true}

	if err := d.DateParse("2006-01-02", "2026-05-26"); err != nil {
		t.Fatal(err)
	}

	want := time.Date(2026, time.May, 26, 0, 0, 0, 0, time.UTC)
	if d.FromDate.Time == nil {
		t.Fatal("expected FromDate to be set")
	}
	if !d.FromDate.Time.Equal(want) {
		t.Fatalf("FromDate = %s, want %s", d.FromDate.Time, want)
	}
}

func TestDateParseReturnsErrorForInvalidDate(t *testing.T) {
	d := Data{GoFormat: true}

	if err := d.DateParse("2006-01-02", "potato"); err == nil {
		t.Fatal("expected parse error")
	}
}

func TestDateAddSetsToDateAndDuration(t *testing.T) {
	d := Data{GoFormat: true}
	d.SetFromDate(time.Date(2026, time.May, 26, 1, 2, 3, 0, time.UTC))

	if err := d.DateAdd("1y 2M 3d 4h"); err != nil {
		t.Fatal(err)
	}

	if d.Duration == nil {
		t.Fatal("expected Duration to be set")
	}
	if d.ToDate.Time == nil {
		t.Fatal("expected ToDate to be set")
	}

	want := time.Date(2027, time.July, 29, 5, 2, 3, 0, time.UTC)
	if !d.ToDate.Time.Equal(want) {
		t.Fatalf("ToDate = %s, want %s", d.ToDate.Time, want)
	}
}

func TestDateAddReturnsInvalidDurationError(t *testing.T) {
	d := Data{GoFormat: true}
	d.SetFromDate(time.Date(2026, time.May, 26, 1, 2, 3, 0, time.UTC))

	if err := d.DateAdd("1d bananas 2h"); err == nil {
		t.Fatal("expected invalid duration error")
	}
}

func TestDateRangeSetsToDateAndRange(t *testing.T) {
	d := Data{GoFormat: true}
	d.SetFromDate(time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC))

	if err := d.DateRange("2006-01-02", "2026-02-01", "1d"); err != nil {
		t.Fatal(err)
	}

	if d.ToDate.Time == nil {
		t.Fatal("expected ToDate to be set")
	}
	if d.Range == nil {
		t.Fatal("expected Range to be set")
	}

	want := time.Date(2026, time.February, 1, 0, 0, 0, 0, time.UTC)
	if !d.ToDate.Time.Equal(want) {
		t.Fatalf("ToDate = %s, want %s", d.ToDate.Time, want)
	}
	if d.Range.Time != 24*time.Hour {
		t.Fatalf("Range = %s, want 24h", d.Range.Time)
	}
}

func TestDateRangeReturnsInvalidDurationError(t *testing.T) {
	d := Data{GoFormat: true}
	d.SetFromDate(time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC))

	if err := d.DateRange("2006-01-02", "2026-02-01", "bad-duration"); err == nil {
		t.Fatal("expected invalid duration error")
	}
}

func TestDateDiffSetsToDateAndDiff(t *testing.T) {
	d := Data{GoFormat: true}
	d.SetFromDate(time.Date(2026, time.January, 1, 1, 2, 3, 0, time.UTC))

	if err := d.DateDiff("2006-01-02 15:04:05", "2027-03-04 05:06:07"); err != nil {
		t.Fatal(err)
	}

	if d.ToDate.Time == nil {
		t.Fatal("expected ToDate to be set")
	}
	if d.Diff == nil {
		t.Fatal("expected Diff to be set")
	}
	if got := d.Diff.String(); got != "1y 2M 3d 4h 4m 4s" {
		t.Fatalf("Diff = %q", got)
	}
}

func TestDateDiffReturnsParseError(t *testing.T) {
	d := Data{GoFormat: true}
	d.SetFromDate(time.Date(2026, time.January, 1, 1, 2, 3, 0, time.UTC))

	if err := d.DateDiff("2006-01-02", "potato"); err == nil {
		t.Fatal("expected parse error")
	}
}

func TestDateTimezoneDotStripsZoneOffset(t *testing.T) {
	loc := time.FixedZone("TEST", 10*60*60)
	d := Data{GoFormat: true}
	d.SetFromDate(time.Date(2026, time.May, 26, 12, 0, 0, 0, loc))

	if err := d.DateTimezone("."); err != nil {
		t.Fatal(err)
	}

	if d.FromDate.Time == nil {
		t.Fatal("expected FromDate to stay set")
	}
	if _, offset := d.FromDate.Time.Zone(); offset != 0 {
		t.Fatalf("offset = %d, want UTC offset 0", offset)
	}
}

func TestDateTimezoneReturnsUnknownTimezoneError(t *testing.T) {
	d := Data{GoFormat: true}
	d.SetFromDate(time.Date(2026, time.May, 26, 12, 0, 0, 0, time.UTC))

	if err := d.DateTimezone("Nope/NotAZone"); err == nil {
		t.Fatal("expected unknown timezone error")
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

func TestDateStateHelpers(t *testing.T) {
	d := Data{}
	if !d.IsDateNil() {
		t.Fatal("new Data should have nil date")
	}

	d.SetDateIfNil()
	if d.IsDateNil() {
		t.Fatal("SetDateIfNil should set FromDate")
	}

	d.SetFromDate(time.Date(2024, time.February, 29, 0, 0, 0, 0, time.UTC))
	if !d.IsDateLeap() {
		t.Fatal("2024 should be leap year")
	}
	if d.IsDateWeekend() {
		t.Fatal("2024-02-29 should not be weekend")
	}
	if !d.IsDateWeekday() {
		t.Fatal("2024-02-29 should be weekday")
	}

	d.SetFromDate(time.Date(2026, time.May, 30, 0, 0, 0, 0, time.UTC))
	if !d.IsDateWeekend() {
		t.Fatal("2026-05-30 should be weekend")
	}
}
