package cmd

import "time"

/*
Examples:
parse date "Sat 31 Jul 1967 09:42:42 AEST" add "20d" format "2006-01-02T15:04:05"
add -- '-1y 12M -1w +7d -2h 120m -2s +2000ms' format '2006-01-02 15:04:05'
tz "UTC" format '2006-01-02 15:04:05'

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

const (
	True = "YES"
	False = "NO"
)
