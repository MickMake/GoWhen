package cal

import (
	"GoWhen/Unify/Only"
	"errors"
	"strconv"
	"strings"
	"time"
)


/*
   Alternative formats - https://programming.guide/go/format-parse-string-time-date-example.html

   Go layout						Java notation				C notation				Notes
   2006-01-02						yyyy-MM-dd					%F						ISO 8601
   20060102						yyyyMMdd					%Y%m%d					ISO 8601
   January 02, 2006				MMMM dd, yyyy				%B %d, %Y
   02 January 2006					dd MMMM yyyy				%d %B %Y
   02-Jan-2006						dd-MMM-yyyy					%d-%b-%Y
   01/02/06						MM/dd/yy					%D						US
   01/02/2006						MM/dd/yyyy					%m/%d/%Y				US
   010206							MMddyy						%m%d%y					US
   Jan-02-06						MMM-dd-yy					%b-%d-%y				US
   Jan-02-2006						MMM-dd-yyyy					%b-%d-%Y				US
   06								yy							%y
   Mon								EEE							%a
   Monday							EEEE						%A
   Jan-06							MMM-yy						%b-%y
   15:04							HH:mm						%R
   15:04:05						HH:mm:ss					%T						ISO 8601
   3:04 PM							K:mm a						%l:%M %p				US
   03:04:05 PM						KK:mm:ss a					%r						US
   2006-01-02T15:04:05				yyyy-MM-dd'T'HH:mm:ss		%FT%T					ISO 8601
   2006-01-02T15:04:05-0700		yyyy-MM-dd'T'HH:mm:ssZ		%FT%T%z					ISO 8601
   2 Jan 2006 15:04:05				d MMM yyyy HH:mm:ss			%e %b %Y %T
   2 Jan 2006 15:04				d MMM yyyy HH:mm			%e %b %Y %R
   Mon, 2 Jan 2006 15:04:05 MST	EEE, d MMM yyyy HH:mm:ss z	%a, %e %b %Y %T %Z		RFC 1123 RFC 822
*/

var TimeFormats = []string{
	"2006-01-02 15:04:05.999999999 -0700 MST",	// time.String() format
	"Mon 02 Jan 2006 15:04:05 MST",	// OSX date
	"2006-01-02T15:04:05",
	"2006-01-02 15:04:05",
	"2006-01-02 15:04",
	"2006-01-02 15",
	"2006-01-02",
	"15:04:05",
	time.ANSIC,
	time.UnixDate,
	time.RubyDate,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.RFC1123,
	time.RFC1123Z,
	time.RFC3339,
	time.RFC3339Nano,
	time.Kitchen,
	time.Stamp,
	time.StampMilli,
	time.StampMicro,
	time.StampNano,
	time.Layout,
}

/*
Sat 01 Jul 1967 09:42:42 AEST

Layout      = "01/02 03:04:05PM '06 -0700"
ANSIC       = "Mon Jan _2 15:04:05 2006"
UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
RFC822      = "02 Jan 06 15:04 MST"
RFC822Z     = "02 Jan 06 15:04 -0700"
RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700"
RFC3339     = "2006-01-02T15:04:05Z07:00"
RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
Kitchen     = "3:04PM"
Stamp      = "Jan _2 15:04:05"
StampMilli = "Jan _2 15:04:05.000"
StampMicro = "Jan _2 15:04:05.000000"
StampNano  = "Jan _2 15:04:05.000000000"

*/


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
