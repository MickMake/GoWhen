package cmd


// type DateTime struct {
// 	format string
// 	Date *time.Time
// 	Duration *time.Duration
// 	// Months int		// Special - because months could be 28, 29, 30, 31 days.
// 	Diff *Diff
// }
//
// type Diff struct {
// 	Year int
// 	Month int
// 	Day int
// 	Hour int
// 	Minute int
// 	Second int
// }
//
// func (d *DateTime) SetDate(t time.Time) {
// 	d.Date = &t
// 	d.Duration = nil
// }
//
// func (d *DateTime) SetDiff(t Diff) {
// 	d.Date = nil
// 	d.Duration = nil
// 	d.Diff = &t
// }
//
// func (d *DateTime) SetDuration(t time.Duration) {
// 	d.Date = nil
// 	d.Duration = &t
// 	// d.Months = months
// }
//
// func (d *DateTime) Clear() {
// 	d.Date = nil
// 	d.Duration = nil
// }
//
// func (d *DateTime) IsWeekend() bool {
// 	if d.Date == nil {
// 		return false
// 	}
// 	switch d.Date.Weekday() {
// 		case time.Sunday:
// 			return true
// 		case time.Saturday:
// 			return true
// 	}
// 	return false
// }
//
// func (d *DateTime) IsWeekday() bool {
// 	return !d.IsWeekend()
// }
//
// func (d *DateTime) IsLeap() bool {
// 	if d.Date == nil {
// 		return false
// 	}
// 	year := d.Date.Year()
// 	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
// }
//
// func (d *DateTime) IsDST() bool {
// 	if d.Date == nil {
// 		return false
// 	}
// 	return d.Date.IsDST()
// }
//
// func (d *DateTime) Parse(format string, timeStr string) (time.Time, error) {
// 	var t time.Time
// 	var err error
//
// 	for range Only.Once {
// 		timeStr = cal.StrToDate(timeStr)
// 		format = cal.StrToFormat(format)
//
// 		// If we have defined a specific format.
// 		if format != "" {
// 			t, err = time.Parse(format, timeStr)
// 			if err != nil {
// 				break
// 			}
// 			break
// 		}
//
// 		// See if we can auto-discover the format.
// 		var l string
// 		l, err = dateparse.ParseFormat(timeStr)
// 		if err == nil {
// 			t, err = time.Parse(l, timeStr)
// 			break
// 		}
//
// 		// Else scan through common set of formats.
// 		for _, f := range cal.TimeFormats {
// 			t, err = time.Parse(f, timeStr)
// 			if err == nil {
// 				// d.SetDate(t)
// 				break
// 			}
// 		}
// 		if err != nil {
// 			break
// 		}
// 	}
//
// 	return t, err
// }
//
// func (d *DateTime) Print() {
// 	for range Only.Once {
// 		if d.Date != nil {
// 			if d.format == "epoch" {
// 				fmt.Printf("%d\n", d.Date.Unix())
// 				break
// 			}
//
// 			if d.format == "week" {
// 				_, w := d.Date.ISOWeek()
// 				fmt.Printf("%d\n", w)
// 				break
// 			}
//
// 			if d.format == "cal" {
// 				month := cal.New(*d.Date).Month()
// 				fmt.Println(month) // Print current month of the calendar.
// 				break
// 			}
//
// 			if d.format == "" {
// 				d.format = time.RFC3339Nano
// 			}
// 			fmt.Printf("%s\n", d.Date.Format(d.format))
// 			break
// 		}
//
// 		if d.Duration != nil {
// 			// s := d.Duration.String()
// 			// replacer := strings.NewReplacer("", ":", "!", "?")
// 			// vd := d.Duration.Hours() / 24
//
// 			fmt.Printf("%s\n", d.Duration.String())
// 			break
// 		}
//
// 		if d.Diff != nil {
// 			var s string
// 			if d.Diff.Year != 0 {
// 				s += fmt.Sprintf("%dy ", d.Diff.Year)
// 			}
//
// 			if d.Diff.Month != 0 {
// 				s += fmt.Sprintf("%dM ", d.Diff.Month)
// 			}
//
// 			if d.Diff.Day != 0 {
// 				s += fmt.Sprintf("%dd ", d.Diff.Day)
// 			}
//
// 			if d.Diff.Hour != 0 {
// 				s += fmt.Sprintf("%dh ", d.Diff.Hour)
// 			}
//
// 			if d.Diff.Minute != 0 {
// 				s += fmt.Sprintf("%dm ", d.Diff.Minute)
// 			}
//
// 			if d.Diff.Second != 0 {
// 				s += fmt.Sprintf("%ds ", d.Diff.Second)
// 			}
// 			s = strings.TrimSpace(s)
//
// 			fmt.Println(s)
// 			break
// 		}
// 	}
// }
