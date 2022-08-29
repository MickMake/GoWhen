package cmd

import "time"

/*
Examples:
parse date "Sat 01 Jul 1967 09:42:42 AEST" add "20d" format "2006-01-02T15:04:05"
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

const (
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
)