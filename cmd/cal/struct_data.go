package cal

import (
	"GoWhen/Unify/Only"
	"errors"
	"fmt"
	"github.com/araddon/dateparse"
	"strings"
	"time"
)


type Diff struct {
	Year   int
	Month  int
	Day    int
	Hour   int
	Minute int
	Second int
}

func (d *Diff) String() string {
	var s string

	for range Only.Once {
		if d.Year != 0 {
			s += fmt.Sprintf("%dy ", d.Year)
		}

		if d.Month != 0 {
			s += fmt.Sprintf("%dM ", d.Month)
		}

		if d.Day != 0 {
			s += fmt.Sprintf("%dd ", d.Day)
		}

		if d.Hour != 0 {
			s += fmt.Sprintf("%dh ", d.Hour)
		}

		if d.Minute != 0 {
			s += fmt.Sprintf("%dm ", d.Minute)
		}

		if d.Second != 0 {
			s += fmt.Sprintf("%ds ", d.Second)
		}
		s = strings.TrimSpace(s)
	}

	return s
}


type Data struct {
	Command    string
	Format     string

	Convert    *Convert
	JavaFormat bool
	CppFormat  bool
	GoFormat   bool
	// FormatType []string
	FormatType string

	FromDate DateTime
	ToDate   DateTime
	Diff     *Diff
	Duration *Duration
	Range    *Duration
}
type DateTime struct {
	*time.Time
}

func (d *Data) ConvertFormat(format string) {
	for range Only.Once {
		format = StrToFormat(format)

		if d.JavaFormat {
			d.Format = d.Convert.FromJava(format)
			break
		}

		if d.CppFormat {
			d.Format = d.Convert.FromCpp(format)
			break
		}
	}
}

func (d *Data) SetCmd(c string) {
	d.Command = c
}

func (d *Data) SetFromDate(t time.Time) {
	d.FromDate.Time = &t
}

func (d *Data) SetToDate(t time.Time) {
	d.ToDate.Time = &t
}

func (d *Data) SetDiff(t Diff) {
	d.Diff = &t
}

func (d *Data) SetDuration(t Duration) {
	d.Duration = &t
}

func (d *Data) SetRange(t Duration) {
	d.Range = &t
}

func (d *Data) DateParse(format string, timeStr string) error {
	var err error
	for range Only.Once {
		d.ConvertFormat(format)

		var t time.Time
		t, err = d.ParseDateString(d.Format, timeStr)
		if err != nil {
			break
		}
		d.SetFromDate(t)
	}
	return err
}


func (d *Data) DateTruncate(duration string) error {
	var err error
	for range Only.Once {
		var dur Duration
		dur, err = ParseDuration(duration)
		if err != nil {
			break
		}
		dur.Time += time.Duration(dur.Months * 30 * 24) * time.Hour
		dur.Time += time.Duration(dur.Years * 365 * 24) * time.Hour

		t := *d.FromDate.Time
		// Truncate and round only works in UTC.
		// So need to convert.
		_, o := d.FromDate.Time.Zone()
		l := d.FromDate.Time.Location()
		if o > 0 {
			// Convert to UTC, without the actual zone conversion.
			t = d.FromDate.Time.UTC().Add(time.Second * time.Duration(o))
		}

		t = t.Truncate(dur.Time)

		if o > 0 {
			// Convert back to previous zone, (if there was one).
			t = t.In(l).Add(time.Second * time.Duration(-o))
		}

		d.SetToDate(t)
	}
	return err
}

func (d *Data) DateRound(duration string) error {
	var err error
	for range Only.Once {
		var dur Duration
		dur, err = ParseDuration(duration)
		if err != nil {
			break
		}
		dur.Time += time.Duration(dur.Months * 30 * 24) * time.Hour
		dur.Time += time.Duration(dur.Years * 365 * 24) * time.Hour

		t := *d.FromDate.Time
		// Truncate and round only works in UTC.
		// So need to convert.
		_, o := d.FromDate.Time.Zone()
		l := d.FromDate.Time.Location()
		if o > 0 {
			// Convert to UTC, without the actual zone conversion.
			t = d.FromDate.Time.UTC().Add(time.Second * time.Duration(o))
		}

		t = t.Round(dur.Time)

		if o > 0 {
			// Convert back to previous zone, (if there was one).
			t = t.In(l).Add(time.Second * time.Duration(-o))
		}

		d.SetToDate(t)
	}
	return err
}

func (d *Data) DateTimezone(loc string) error {
	var err error
	for range Only.Once {
		if (loc == "") || (loc == ".") {
			// Strip zone info if empty.
			_, o := d.FromDate.Time.Zone()
			t := d.FromDate.Time.UTC().Add(time.Second * time.Duration(o))
			d.FromDate.Time = &t
			break
		}

		var l *time.Location
		l, err = time.LoadLocation(loc)
		if err != nil {
			err = errors.New("unknown timezone '" + loc + "'")
			break
		}
		t := d.FromDate.Time.In(l)
		d.SetToDate(t)
	}
	return err
}

// DateAdd - SetDuration / SetToDate
func (d *Data) DateAdd(duration string) error {
	var err error
	for range Only.Once {
		var dur Duration
		dur, err = ParseDuration(duration)
		if err != nil {
			break
		}
		d.SetDuration(dur)
		t := d.FromDate.AddDate(int(dur.Years), int(dur.Months), 0)
		t = t.Add(dur.Time)
		// d.SetFromDate(t)
		d.SetToDate(t)
	}
	return err
}


// DateRange - SetToDate
func (d *Data) DateRange(format string, toStr string, duration string) error {
	var err error
	for range Only.Once {
		d.ConvertFormat(format)

		if (d.ToDate.Time == nil) || (toStr != "") {
			var t time.Time
			t, err = d.ParseDateString(d.Format, toStr)
			if err != nil {
				break
			}
			d.SetToDate(t)
		}

		var td Duration
		td, err = ParseDuration(duration)
		if err != nil {
			break
		}
		d.SetRange(td)
	}
	return err
}

// DateDiff - SetToDate / SetDiff
func (d *Data) DateDiff(format string, timeStr string) error {
	var err error
	for range Only.Once {
		d.ConvertFormat(format)

		var t time.Time
		t, err = d.ParseDateString(d.Format, timeStr)
		if err != nil {
			break
		}
		d.SetToDate(t)
		diff := DateDiff(*d.FromDate.Time, t)
		d.SetDiff(diff)
	}
	return err
}


func (d *Data) IsDateNil() bool {
	if d.FromDate.Time == nil {
		return true
	}
	return false
}

func (d *Data) SetDateIfNil() {
	if d.FromDate.Time == nil {
		d.SetFromDate(time.Now())
	}
}

func (d *Data) IsDateWeekend() bool {
	if d.FromDate.Time == nil {
		return false
	}
	switch d.FromDate.Time.Weekday() {
		case time.Sunday:
			return true
		case time.Saturday:
			return true
	}
	return false
}

func (d *Data) IsDateWeekday() bool {
	return !d.IsDateWeekend()
}

func (d *Data) IsDateLeap() bool {
	if d.FromDate.Time == nil {
		return false
	}
	year := d.FromDate.Time.Year()
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func (d *Data) IsDateDST() bool {
	if d.FromDate.Time == nil {
		return false
	}
	return d.FromDate.Time.IsDST()
}

func (d *Data) IsDateBefore(format string, timeStr string) bool {
	var yes bool
	for range Only.Once {
		d.ConvertFormat(format)

		t, err := d.ParseDateString(d.Format, timeStr)
		if err != nil {
			break
		}
		if d.FromDate.Time.Before(t) {
			yes = true
			break
		}
		yes = false
	}
	return yes
}

func (d *Data) IsDateAfter(format string, timeStr string) bool {
	var yes bool
	for range Only.Once {
		d.ConvertFormat(format)

		t, err := d.ParseDateString(d.Format, timeStr)
		if err != nil {
			break
		}
		if d.FromDate.Time.After(t) {
			yes = true
			break
		}
		yes = false
	}
	return yes
}


// func (d *Data) IsDate() bool {
// 	if d.FromDate.Time != nil {
// 		return true
// 	}
// 	return false
// }
//
// func (d *Data) IsDiff() bool {
// 	if d.Diff != nil {
// 		return true
// 	}
// 	return false
// }
//
// func (d *Data) IsDuration() bool {
// 	if d.Duration != nil {
// 		return true
// 	}
// 	return false
// }


func (d *Data) Clear() {
	d.FromDate.Time = nil
	d.ToDate.Time = nil
	d.Diff = nil
	d.Duration = nil
}

// func (d *Data) DateSet(t time.Time) {
// 	d.Date.Time = &t
// 	d.Duration = nil
// 	d.Diff = nil
// }
//
// func (d *Data) DiffSet(t Diff) {
// 	d.Date.Time = nil
// 	d.Duration = nil
// 	d.Diff = &t
// }
//
// func (d *Data) DurationSet(t time.Duration) {
// 	d.Date.Time = nil
// 	d.Duration = &t
// 	d.Diff = nil
// }

func (d *Data) ParseDateString(format string, timeStr string) (time.Time, error) {
	var t time.Time
	var err error

	for range Only.Once {
		if format != "" {
			d.ConvertFormat(format)
		}

		if timeStr == "." {
			timeStr = ""
		}
		if (d.ToDate.Time != nil) && (timeStr == "") {
			t = *d.ToDate.Time
			break
		}

		t2 := StrToDate(timeStr)
		if t2 != nil {
			t = *t2
			break
		}

		// If we have defined a specific format.
		if d.Format != "" {
			t, err = time.Parse(format, timeStr)
			if err == nil {
				break
			}
		}

		// See if we can auto-discover the format.
		format, err = dateparse.ParseFormat(timeStr)
		if err == nil {
			d.Format = format	// Will be in GoLang layout format.
			t, err = time.Parse(format, timeStr)
			break
		}

		// Else scan through common set of formats.
		for _, f := range TimeFormats {
			t, err = time.Parse(f, timeStr)
			if err == nil {
				d.Format = f
				break
			}
		}
		if err != nil {
			break
		}
	}

	return t, err
}


func (d *Data) Print() {
	for range Only.Once {
		if d.Range != nil {
			d.PrintRange()
			break
		}

		if d.Diff != nil {
			d.PrintDiff()
			break
		}

		// if d.Duration != nil {
		// 	d.PrintDuration()
		// 	break
		// }

		if d.ToDate.Time != nil {
			d.PrintToDate()
			break
		}

		if d.FromDate.Time != nil {
			d.PrintFromDate()
			break
		}
	}
}

func (d *Data) PrintFromDate() {
	for range Only.Once {
		if d.FromDate.Time == nil {
			break
		}

		if d.Format == "epoch" {
			fmt.Printf("%d\n", d.FromDate.Time.Unix())
			break
		}

		if d.Format == "week" {
			_, w := d.FromDate.Time.ISOWeek()
			fmt.Printf("%d\n", w)
			break
		}

		if d.Format == "list" {
			m := New(*d.FromDate.Time).Week()
			m.Print()
			break
		}

		if d.Format == "cal-week" {
			m := New(*d.FromDate.Time).Week()
			m.Print()
			break
		}

		if d.Format == "cal-month" {
			m := New(*d.FromDate.Time).Month()
			m.Print()
			break
		}

		if d.Format == "cal-year" {
			y := New(*d.FromDate.Time).Year()
			y.Print()
			break
		}

		if d.Format == "" {
			d.Format = time.RFC3339Nano
		}
		fmt.Printf("%s\n", d.FromDate.Time.Format(d.Format))
	}
}

func (d *Data) PrintToDate() {
	for range Only.Once {
		if d.ToDate.Time == nil {
			break
		}

		if d.Format == "epoch" {
			fmt.Printf("%d\n", d.ToDate.Time.Unix())
			break
		}

		if d.Format == "week" {
			_, w := d.ToDate.Time.ISOWeek()
			fmt.Printf("%d\n", w)
			break
		}

		if d.Format == "list" {
			m := New(*d.ToDate.Time).Week()
			m.Print()
			break
		}

		if d.Format == "cal-week" {
			m := New(*d.ToDate.Time).Week()
			m.Print()
			break
		}

		if d.Format == "cal-month" {
			m := New(*d.ToDate.Time).Month()
			m.Print()
			break
		}

		if d.Format == "cal-year" {
			y := New(*d.ToDate.Time).Year()
			y.Print()
			break
		}

		if d.Format == "" {
			d.Format = time.RFC3339Nano
		}
		fmt.Printf("%s\n", d.ToDate.Time.Format(d.Format))
	}
}

func (d *Data) PrintDuration() {
	for range Only.Once {
		if d.Duration == nil {
			break
		}
		fmt.Printf("%dy %dM %s\n", d.Duration.Years, d.Duration.Months, d.Duration.Time.String())
	}
}

func (d *Data) PrintDiff() {
	for range Only.Once {
		if d.Diff == nil {
			break
		}
		fmt.Println(d.Diff.String())
	}
}

func (d *Data) PrintRange() {
	for range Only.Once {
		if d.Range == nil {
			break
		}

		if d.Format == "" {
			d.Format = time.RFC3339
		}

		var lt time.Time
		if d.ToDate.Time.Before(*d.FromDate.Time) {
			for t := *d.FromDate.Time; t.After(*d.ToDate.Time); {
				fmt.Println(t.Format(d.Format))
				t = t.AddDate(-int(d.Range.Years), -int(d.Range.Months), 0).Add(-d.Range.Time)
				if lt == t {
					// Avoid endless loops
					break
				}
				lt = t
			}
			break
		}

		for t := *d.FromDate.Time; t.Before(*d.ToDate.Time); {
			fmt.Println(t.Format(d.Format))
			t = t.AddDate(int(d.Range.Years), int(d.Range.Months), 0).Add(d.Range.Time)
			if lt == t {
				// Avoid endless loops
				break
			}
			lt = t

		}
	}
}
