package cal

import (
	"GoWhen/Unify/Only"
	"fmt"
	"github.com/araddon/dateparse"
	"github.com/olekukonko/tablewriter"
	"os"
	"strings"
	"time"
)


type DateTime struct {
	Format string
	Date   *time.Time
	Duration *time.Duration
	// Months int		// Special - because months could be 28, 29, 30, 31 days.
	Diff *Diff
}

type Diff struct {
	Year int
	Month int
	Day int
	Hour int
	Minute int
	Second int
}


func (d *DateTime) SetDate(t time.Time) {
	d.Date = &t
	d.Duration = nil
}

func (d *DateTime) SetDiff(t Diff) {
	d.Date = nil
	d.Duration = nil
	d.Diff = &t
}

func (d *DateTime) SetDuration(t time.Duration) {
	d.Date = nil
	d.Duration = &t
	// d.Months = months
}

func (d *DateTime) Clear() {
	d.Date = nil
	d.Duration = nil
}

func (d *DateTime) IsWeekend() bool {
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

func (d *DateTime) IsWeekday() bool {
	return !d.IsWeekend()
}

func (d *DateTime) IsLeap() bool {
	if d.Date == nil {
		return false
	}
	year := d.Date.Year()
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func (d *DateTime) IsDST() bool {
	if d.Date == nil {
		return false
	}
	return d.Date.IsDST()
}

func (d *DateTime) Parse(format string, timeStr string) (time.Time, error) {
	var t time.Time
	var err error

	for range Only.Once {
		timeStr = StrToDate(timeStr)
		format = StrToFormat(format)

		// If we have defined a specific format.
		if format != "" {
			t, err = time.Parse(format, timeStr)
			if err != nil {
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

func (d *DateTime) Print() {
	for range Only.Once {
		if d.Date != nil {
			if d.Format == "epoch" {
				fmt.Printf("%d\n", d.Date.Unix())
				break
			}

			if d.Format == "week" {
				_, w := d.Date.ISOWeek()
				fmt.Printf("%d\n", w)
				break
			}

			if d.Format == "cal-week" {
				m := New(*d.Date).Week()
				m.Print()
				break
			}

			if d.Format == "cal-month" {
				m := New(*d.Date).Month()
				m.Print()
				break
			}

			if d.Format == "cal-year" {
				y := New(*d.Date).Year()
				y.Print()
				break
			}

			if d.Format == "" {
				d.Format = time.RFC3339Nano
			}
			fmt.Printf("%s\n", d.Date.Format(d.Format))
			break
		}

		if d.Duration != nil {
			// s := d.Duration.String()
			// replacer := strings.NewReplacer("", ":", "!", "?")
			// vd := d.Duration.Hours() / 24

			fmt.Printf("%s\n", d.Duration.String())
			break
		}

		if d.Diff != nil {
			var s string
			if d.Diff.Year != 0 {
				s += fmt.Sprintf("%dy ", d.Diff.Year)
			}

			if d.Diff.Month != 0 {
				s += fmt.Sprintf("%dM ", d.Diff.Month)
			}

			if d.Diff.Day != 0 {
				s += fmt.Sprintf("%dd ", d.Diff.Day)
			}

			if d.Diff.Hour != 0 {
				s += fmt.Sprintf("%dh ", d.Diff.Hour)
			}

			if d.Diff.Minute != 0 {
				s += fmt.Sprintf("%dm ", d.Diff.Minute)
			}

			if d.Diff.Second != 0 {
				s += fmt.Sprintf("%ds ", d.Diff.Second)
			}
			s = strings.TrimSpace(s)

			fmt.Println(s)
			break
		}
	}
}


func (y *Year) Print()  {
	x := (*y)[1][1][1]
	fmt.Printf("|-------------- %s --------------|\n", x.Format("2006"))
	for _, m := range *y {
		m.Print()
		fmt.Println()
	}
}

func (m *Month) Print()  {
	x := (*m)[1][1]
	fmt.Printf("| %s\n", x.Format("Jan 2006"))
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"S", "M", "T", "W", "T", "F", "S"})
	table.SetBorder(true)
	for _, month := range *m {
		var row []string
		for _, week := range month {
			row = append(row, fmt.Sprintf("%d", week.Day()))
		}
		table.Append(row)
	}
	table.Render()
}

func (w *Week) Print()  {
	x := (*w)[1]
	fmt.Printf("| %s\n", x.Format("Jan 2006"))
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"S", "M", "T", "W", "T", "F", "S"})
	table.SetBorder(true)
	var row []string
	for _, week := range *w {
		row = append(row, fmt.Sprintf("%d", week.Day()))
	}
	table.Append(row)
	table.Render()
}