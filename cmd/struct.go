package cmd

import (
	"GoWhen/Unify/Only"
	"fmt"
	"github.com/araddon/dateparse"
	"time"
)


type Data struct {
	format string
	Date *time.Time
	Duration *time.Duration
}

func (d *Data) SetDate(t time.Time) {
	d.Date = &t
	d.Duration = nil
}

func (d *Data) SetDuration(t time.Duration) {
	d.Date = nil
	d.Duration = &t
}

func (d *Data) Clear() {
	d.Date = nil
	d.Duration = nil
}

func (d *Data) IsWeekend() bool {
	if d.Date == nil {
		return false
	}
	switch d.Date.Weekday() {
	case time.Sunday:
		return true
	case time.Saturday:
		return true
	}
	return false
}

func (d *Data) IsWeekday() bool {
	return !d.IsWeekend()
}

func (d *Data) IsLeap() bool {
	if d.Date == nil {
		return false
	}
	year := d.Date.Year()
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func (d *Data) IsDST() bool {
	if d.Date == nil {
		return false
	}
	return d.Date.IsDST()
}

func (d *Data) Parse(format string, timeStr string) (time.Time, error) {
	var t time.Time
	var err error

	for range Only.Once {
		// If we have defined a specific format.
		if format != "" {
			format = StrToFormat(format)
			t, err = time.Parse(format, timeStr)
			if err == nil {
				// d.SetDate(t)
				break
			}
			break
		}

		// See if we can auto-discover the format.
		var l string
		l, err = dateparse.ParseFormat(timeStr)
		if err == nil {
			t, err = time.Parse(l, timeStr)
			break
		}

		// Else scan through common set of formats.
		for _, f := range TimeFormats {
			t, err = time.Parse(f, timeStr)
			if err == nil {
				// d.SetDate(t)
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
		if d.Date != nil {
			if d.format == "epoch" {
				fmt.Printf("%d\n", d.Date.Unix())
				break
			}

			if d.format == "week" {
				_, w := d.Date.ISOWeek()
				fmt.Printf("%d\n", w)
				break
			}

			if d.format == "" {
				d.format = time.RFC3339Nano
			}
			fmt.Printf("%s\n", d.Date.Format(d.format))
			break
		}

		if d.Duration != nil {
			fmt.Printf("%s\n", d.Duration.String())
			break
		}
	}
}
