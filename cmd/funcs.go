package cmd

import (
	"strings"
	"time"
)

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

func StrToDate(str string) string {
	s := strings.ToLower(str)
	switch {
		case s == "":
			fallthrough
		case s == "now":
			fallthrough
		case s == "today":
			str = time.Now().Format(time.RFC3339)

		case s == "tomorrow":
			str = time.Now().Add(time.Hour * 24).Format(time.RFC3339)

		case s == "next-week":
			str = time.Now().Add(time.Hour * 168).Format(time.RFC3339)

		case s == "yesterday":
			str = time.Now().Add(time.Hour * -24).Format(time.RFC3339)

		case s == "last-week":
			str = time.Now().Add(time.Hour * -168).Format(time.RFC3339)

		case s == "epoch":
			str = "1970-01-01 00:00:00"

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
	return str
}
