package cal

import (
	"GoWhen/Unify/Only"
	"errors"
	"strconv"
	"strings"
	"time"
)


type Duration struct {
	Time time.Duration
	Years int64
	Months int64
	// Weeks int	// Handled by classic time.Duration
	// Days int	// Handled by classic time.Duration
}

func ParseDuration(s string) (Duration, error) {
	var duration Duration
	var err error

	for range Only.Once {
		times := strings.Split(s, " ")

		for _, ds := range times {
			ds = strings.TrimSpace(ds)
			if ds == "" {
				continue
			}

			var d time.Duration
			d, err = time.ParseDuration(ds)
			if err == nil {
				duration.Time += d
				continue
			}

			//
			// neg := false
			// c := ds[0]
			// if c == '-' || c == '+' {
			// 	neg = c == '-'
			// 	ds = ds[1:]
			// }

			lb := ds[len(ds)-1]

			switch lb {
				case 'Y':
					fallthrough
				case 'y':
					// Using DateAdd type duration.
					var lbv int64
					lbv, err = strconv.ParseInt(ds[:len(ds)-1], 10, 64)
					if err != nil {
						break
					}
					duration.Years += lbv

				case 'M':
					// Using DateAdd type duration.
					var lbv int64
					lbv, err = strconv.ParseInt(ds[:len(ds)-1], 10, 64)
					if err != nil {
						break
					}
					duration.Months += lbv

				case 'W':
					fallthrough
				case 'w':
					// Straight-forward conversion.
					var lbv float64
					lbv, err = strconv.ParseFloat(ds[:len(ds)-1], 10)
					if err != nil {
						break
					}
					v := float64(int64(time.Hour) * 168) * lbv
					duration.Time += time.Duration(v)

				case 'D':
					fallthrough
				case 'd':
					// Straight-forward conversion.
					var lbv float64
					lbv, err = strconv.ParseFloat(ds[:len(ds)-1], 10)
					if err != nil {
						break
					}
					v := float64(int64(time.Hour) * 24) * lbv
					duration.Time += time.Duration(v)

				default:
					err = errors.New("time: invalid duration " + ds)
					break
			}
		}
	}

	// for range Only.Once {
	// 	var fv int64
	// 	for _, v := range duration.Time {
	// 		fv += int64(v)
	// 	}
	// 	duration.Time = []time.Duration{time.Duration(fv)}
	//
	// 	for _, v := range duration.Time {
	// 		fv += int64(v)
	// 	}
	// 	duration.Years = []int64{fv}
	//
	// 	for _, v := range duration.Time {
	// 		fv += int64(v)
	// 	}
	// 	duration.Months = []int64{fv}
	// }

	return duration, err
}

// DateDiff - Stolen from https://stackoverflow.com/questions/36530251/time-since-with-months-and-years
//goland:noinspection GoRedundantConversion
func DateDiff(a, b time.Time) Diff {
	var d Diff

	for range Only.Once {
		var year, month, day, hour, min, sec int

		if a.Location() != b.Location() {
			b = b.In(a.Location())
		}

		if a.After(b) {
			a, b = b, a
		}

		y1, M1, d1 := a.Date()
		y2, M2, d2 := b.Date()

		h1, m1, s1 := a.Clock()
		h2, m2, s2 := b.Clock()

		year = int(y2 - y1)
		month = int(M2 - M1)
		day = int(d2 - d1)
		hour = int(h2 - h1)
		min = int(m2 - m1)
		sec = int(s2 - s1)

		// Normalize negative values
		if sec < 0 {
			sec += 60
			min--
		}

		if min < 0 {
			min += 60
			hour--
		}

		if hour < 0 {
			hour += 24
			day--
		}

		if day < 0 {
			// days in month:
			t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
			day += 32 - t.Day()
			month--
		}

		if month < 0 {
			month += 12
			year--
		}

		d = Diff{
			Year:   year,
			Month:  month,
			Day:    day,
			Hour:   hour,
			Minute: min,
			Second: sec,
		}
	}

	return d
}

func StrToFormat(str string) string {
	switch strings.ToLower(str) {
		case ".":
			str = ""
		case "simple":
			str = "2006-01-02T15:04:05"
		case "layout":
			str = "01/02 03:04:05PM '06 -0700"
		case "ansic":
			str = "Mon Jan _2 15:04:05 2006"
		case "unixdate":
			str = "Mon Jan _2 15:04:05 MST 2006"
		case "rubydate":
			str = "Mon Jan 02 15:04:05 -0700 2006"
		case "rfc822":
			str = "02 Jan 06 15:04 MST"
		case "rfc822z":
			str = "02 Jan 06 15:04 -0700"
		case "rfc850":
			str = "Monday, 02-Jan-06 15:04:05 MST"
		case "rfc1123":
			str = "Mon, 02 Jan 2006 15:04:05 MST"
		case "rfc1123z":
			str = "Mon, 02 Jan 2006 15:04:05 -0700"
		case "rfc3339":
			str = "2006-01-02T15:04:05Z07:00"
		case "rfc3339nano":
			str = "2006-01-02T15:04:05.999999999Z07:00"
		case "kitchen":
			str = "3:04PM"
		case "stamp":
			str = "Jan _2 15:04:05"
		case "stampmilli":
			str = "Jan _2 15:04:05.000"
		case "stampmicro":
			str = "Jan _2 15:04:05.000000"
		case "stampnano":
			str = "Jan _2 15:04:05.000000000"

		// Special cases.
		case "epoch":
			str = "epoch"
		case "week":
			str = "week"
	}
	return str
}

func StrToDate(str string) *time.Time {
	var ret *time.Time

	for range Only.Once {
		s := strings.ToLower(str)
		switch {
			case s == "":
				fallthrough
			case s == "now":
				fallthrough
			case s == "today":
				r := time.Now()
				ret = &r

			case s == "tomorrow":
				r := time.Now().Add(time.Hour * 24)
				ret = &r

			case s == "next-week":
				r := time.Now().Add(time.Hour * 168)
				ret = &r

			case s == "yesterday":
				r := time.Now().Add(time.Hour * -24)
				ret = &r

			case s == "last-week":
				r := time.Now().Add(time.Hour * -168)
				ret = &r

			case s == "epoch":
				r, _ := time.Parse("2006-01-02 15:04:05", "1970-01-01 00:00:00")
				ret = &r

			// case strings.HasPrefix(s, "last"):
			// 	r := strings.TrimPrefix(s, "last")
			// 	r = strings.TrimSpace(r)
			// 	switch r {
			// 		case "decade":
			// 			str = time.Now().Add(time.Hour * -24).Format(time.RFC3339)
			//
			// 		case "year":
			// 			str = time.Now().Add(time.Hour * -24).Format(time.RFC3339)
			//
			// 		case "month":
			// 			str = time.Now().Add(time.Hour * -24).Format(time.RFC3339)
			//
			// 		case "week":
			// 			str = time.Now().Add(time.Hour * -168).Format(time.RFC3339)
			// 	}
		}
	}
	return ret
}
