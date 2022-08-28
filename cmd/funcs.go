package cmd

import (
	"strings"
)

func StrToFormat(str string) string {
	switch strings.ToLower(str) {
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
	}
	return str
}
