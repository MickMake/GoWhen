package cal

import "time"


const (
	stdLongMonth             = "January"
	stdMonth                 = "Jan"
	stdNumMonth              = "1"
	stdZeroMonth             = "01"
	stdLongWeekDay           = "Monday"
	stdWeekDay               = "Mon"
	stdDay                   = "2"
	stdUnderDay              = "_2"
	stdZeroDay               = "02"
	stdUnderYearDay          = "__2"
	stdZeroYearDay           = "002"
	stdHour                  = "15"
	stdHour12                = "3"
	stdZeroHour12            = "03"
	stdMinute                = "4"
	stdZeroMinute            = "04"
	stdSecond                = "5"
	stdZeroSecond            = "05"
	stdLongYear              = "2006"
	stdYear                  = "06"
	stdPM                    = "PM"
	stdpm                    = "pm"
	stdTZ                    = "MST"
	stdISO8601TZ             = "Z0700"
	stdISO8601SecondsTZ      = "Z070000"
	stdISO8601ShortTZ        = "Z07"
	stdISO8601ColonTZ        = "Z07:00"
	stdISO8601ColonSecondsTZ = "Z07:00:00"
	stdNumTZ                 = "-0700"
	stdNumSecondsTz          = "-070000"
	stdNumShortTZ            = "-07"
	stdNumColonTZ            = "-07:00"
	stdNumColonSecondsTZ     = "-07:00:00"
	stdFracSecond0           = ".0"
	stdFracSecond9           = ".9"
)

var LayoutOptions = map[string]string{
	"stdLongMonth": stdLongMonth,
	"stdMonth": stdMonth,
	"stdNumMonth": stdNumMonth,
	"stdZeroMonth": stdZeroMonth,
	"stdLongWeekDay": stdLongWeekDay,
	"stdWeekDay": stdWeekDay,
	"stdDay": stdDay,
	"stdUnderDay": stdUnderDay,
	"stdZeroDay": stdZeroDay,
	"stdUnderYearDay": stdUnderYearDay,
	"stdZeroYearDay": stdZeroYearDay,
	"stdHour": stdHour,
	"stdHour12": stdHour12,
	"stdZeroHour12": stdZeroHour12,
	"stdMinute": stdMinute,
	"stdZeroMinute": stdZeroMinute,
	"stdSecond": stdSecond,
	"stdZeroSecond": stdZeroSecond,
	"stdLongYear": stdLongYear,
	"stdYear": stdYear,
	"stdPM": stdPM,
	"stdpm": stdpm,
	"stdTZ": stdTZ,
	"stdISO8601TZ": stdISO8601TZ,
	"stdISO8601SecondsTZ": stdISO8601SecondsTZ,
	"stdISO8601ShortTZ": stdISO8601ShortTZ,
	"stdISO8601ColonTZ": stdISO8601ColonTZ,
	"stdISO8601ColonSecondsTZ": stdISO8601ColonSecondsTZ,
	"stdNumTZ": stdNumTZ,
	"stdNumSecondsTz": stdNumSecondsTz,
	"stdNumShortTZ": stdNumShortTZ,
	"stdNumColonTZ": stdNumColonTZ,
	"stdNumColonSecondsTZ": stdNumColonSecondsTZ,
	"stdFracSecond0": stdFracSecond0,
	"stdFracSecond9": stdFracSecond9,
}

var Layouts = map[string]string{
	"ANSIC": time.ANSIC,
	"UnixDate": time.UnixDate,
	"RubyDate": time.RubyDate,
	"RFC822": time.RFC822,
	"RFC822Z": time.RFC822Z,
	"RFC850": time.RFC850,
	"RFC1123": time.RFC1123,
	"RFC1123Z": time.RFC1123Z,
	"RFC3339": time.RFC3339,
	"RFC3339Nano": time.RFC3339Nano,
	"Kitchen": time.Kitchen,
	"Stamp": time.Stamp,
	"StampMilli": time.StampMilli,
	"StampMicro": time.StampMicro,
	"StampNano": time.StampNano,
	"Layout": time.Layout,
}

var TimeFormats = []string{
	"2006-01-02 15:04:05.999999999 -0700 MST",	// time.String() format
	"Mon 02 Jan 2006 15:04:05 MST",	// OSX date
	"2006-01-02T15:04:05",
	"2006-01-02 15:04:05",
	"2006-01-02 15:04",
	"2006-01-02 15",
	"2006-01-02",
	"15:04:05",
	time.ANSIC,			// "Mon Jan _2 15:04:05 2006"
	time.UnixDate,		// "Mon Jan _2 15:04:05 MST 2006"
	time.RubyDate,		// "Mon Jan 02 15:04:05 -0700 2006"
	time.RFC822,		// "02 Jan 06 15:04 MST"
	time.RFC822Z,		// "02 Jan 06 15:04 -0700"
	time.RFC850,		// "Monday, 02-Jan-06 15:04:05 MST"
	time.RFC1123,		// "Mon, 02 Jan 2006 15:04:05 MST"
	time.RFC1123Z,		// "Mon, 02 Jan 2006 15:04:05 -0700"
	time.RFC3339,		// "2006-01-02T15:04:05Z07:00"
	time.RFC3339Nano,	// "2006-01-02T15:04:05.999999999Z07:00"
	time.Kitchen,		// "3:04PM"
	time.Stamp,			// "Jan _2 15:04:05"
	time.StampMilli,	// "Jan _2 15:04:05.000"
	time.StampMicro,	// "Jan _2 15:04:05.000000"
	time.StampNano,		// "Jan _2 15:04:05.000000000"
	time.Layout,		// "01/02 03:04:05PM '06 -0700"
}


/*
   Alternative formats - https://programming.guide/go/format-parse-string-time-date-example.html

   Go layout						Java notation				C notation				Notes
   20060102							yyyyMMdd					%Y%m%d					ISO 8601
   January 02, 2006					MMMM dd, yyyy				%B %d, %Y
   02 January 2006					dd MMMM yyyy				%d %B %Y
   02-Jan-2006						dd-MMM-yyyy					%d-%b-%Y
   01/02/2006						MM/dd/yyyy					%m/%d/%Y				US
   010206							MMddyy						%m%d%y					US
   Jan-02-06						MMM-dd-yy					%b-%d-%y				US
   Jan-02-2006						MMM-dd-yyyy					%b-%d-%Y				US

   3:04 PM							K:mm a						%l:%M %p				US
   2006-01-02T15:04:05				yyyy-MM-dd'T'HH:mm:ss		%FT%T					ISO 8601
   2006-01-02T15:04:05-0700			yyyy-MM-dd'T'HH:mm:ssZ		%FT%T%z					ISO 8601
   2 Jan 2006 15:04:05				d MMM yyyy HH:mm:ss			%e %b %Y %T
   2 Jan 2006 15:04					d MMM yyyy HH:mm			%e %b %Y %R
   Mon, 2 Jan 2006 15:04:05 MST	EEE, d MMM yyyy HH:mm:ss z		%a, %e %b %Y %T %Z		RFC 1123 RFC 822

   01/02/06							MM/dd/yy					%D	- equivalent to "%m/%d/%y"
   15:04							HH:mm						%R	- equivalent to "%H:%M"
   15:04:05							HH:mm:ss					%T	- equivalent to "%H:%M:%S" (the ISO 8601 time format)
   03:04:05 PM						KK:mm:ss a					%r	- writes localized 12-hour clock time (locale dependent)
   2006-01-02						yyyy-MM-dd					%F	- equivalent to "%Y-%m-%d" (the ISO 8601 date format)
   ?								?							%c	- writes standard date and time string, e.g. Sun Oct 17 04:41:13 2010 (locale dependent)
   ?								?							%x	- writes localized date representation (locale dependent)
   ?								?							%X	- writes localized time representation, e.g. 18:40:20 or 6:40:20 PM (locale dependent)
   2-Jan-2006						d-MMM-YYYY					%v	- is equivalent to "%e-%b-%Y"



   2006								yyyy						%Y	- writes year as a decimal number, e.g. 2017
   06								yy							%y	- writes last 2 digits of year as a decimal number (range [00,99])
   January							MMMM						%B	- writes full month name, e.g. October (locale dependent)
   Jan								MMM							%b %h	- writes abbreviated month name, e.g. Oct (locale dependent)
   01								MM							%m	- writes month as a decimal number (range [01,12])
   02								dd							%d	- writes day of the month as a decimal number (range [01,31])
   2								d							%e	- writes day of the month as a decimal number (range [1,31]).
   Mon								EEE							%a	- writes abbreviated weekday name, e.g. Fri (locale dependent)
   Monday							EEEE						%A	- writes full weekday name, e.g. Friday (locale dependent)

   ?								?							%j	- Day of the year (001-366).
   ?								?							%U	- Week number with the first Sunday as the first day of week one (00-53).
   ?								?							%W	- Week number with the first Monday as the first day of week one (00-53).
   ?								?							%w	- Weekday as a decimal number with Sunday as 0 (0-6).
   ?								?							%u	- writes weekday as a decimal number, where Monday is 1 (ISO 8601 format) (range [1-7])

   15								HH							%H	- writes hour as a decimal number, 24 hour clock (range [00-23])
   3								K							%l	-
   03								KK							%I	- writes hour as a decimal number, 12 hour clock (range [01,12])
   PM								a							%p	- writes localized a.m. or p.m. (locale dependent)
   04								mm							%M	- writes minute as a decimal number (range [00,59])
   05								ss							%S	- writes second as a decimal number (range [00,60])
   -0700							Z							%z	- writes offset from UTC in the ISO 8601 format (e.g. -0430), or no characters if the time zone information is not available
   EEE								z							%Z	- writes locale-dependent time zone name or abbreviation, or no characters if the time zone information is not available

*/

const DefaultConvertConfig = `[
  { "go": "20060102", "java": "yyyyMMdd", "c": "%Y%m%d", "notes": "ISO 8601" },
  { "go": "January 02, 2006", "java": "MMMM dd, yyyy", "c": "%B %d, %Y", "notes": "" },
  { "go": "02 January 2006", "java": "dd MMMM yyyy", "c": "%d %B %Y", "notes": "" },
  { "go": "02-Jan-2006", "java": "dd-MMM-yyyy", "c": "%d-%b-%Y", "notes": "" },
  { "go": "01/02/2006", "java": "MM/dd/yyyy", "c": "%m/%d/%Y", "notes": "US" },
  { "go": "010206", "java": "MMddyy", "c": "%m%d%y", "notes": "US" },
  { "go": "Jan-02-06", "java": "MMM-dd-yy", "c": "%b-%d-%y", "notes": "US" },
  { "go": "Jan-02-2006", "java": "MMM-dd-yyyy", "c": "%b-%d-%Y", "notes": "US" },
  { "go": "3:04 PM", "java": "K:mm a", "c": "%l:%M %p", "notes": "US" },
  { "go": "2006-01-02T15:04:05", "java": "yyyy-MM-dd'T'HH:mm:ss", "c": "%FT%T", "notes": "ISO 8601" },
  { "go": "2006-01-02T15:04:05-0700", "java": "yyyy-MM-dd'T'HH:mm:ssZ", "c": "%FT%T%z", "notes": "ISO 8601" },
  { "go": "2 Jan 2006 15:04:05", "java": "d MMM yyyy HH:mm:ss", "c": "%e %b %Y %T", "notes": "" },
  { "go": "2 Jan 2006 15:04", "java": "d MMM yyyy HH:mm", "c": "%e %b %Y %R", "notes": "" },
  { "go": "Mon, 2 Jan 2006 15:04:05 MST", "java": "EEE, d MMM yyyy HH:mm:ss z", "c": "%a, %e %b %Y %T %Z", "notes": "RFC 1123 RFC 822" },
  { "go": "01/02/06", "java": "MM/dd/yy", "c": "%D", "notes": "- equivalent to \"%m/%d/%y\"" },
  { "go": "15:04", "java": "HH:mm", "c": "%R", "notes": "- equivalent to \"%H:%M\"" },
  { "go": "15:04:05", "java": "HH:mm:ss", "c": "%T", "notes": "- equivalent to \"%H:%M:%S\" (the ISO 8601 time format)" },
  { "go": "03:04:05 PM", "java": "KK:mm:ss a", "c": "%r", "notes": "- writes localized 12-hour clock time (locale dependent)" },
  { "go": "2006-01-02", "java": "yyyy-MM-dd", "c": "%F", "notes": "- equivalent to \"%Y-%m-%d\" (the ISO 8601 date format)" },
  { "go": "Jan 02 15:04:05 2006", "java": "", "c": "%c", "notes": "- writes standard date and time string, e.g. Sun Oct 17 04:41:13 2010 (locale dependent)" },
  { "go": "", "java": "", "c": "%x", "notes": "- writes localized date representation (locale dependent)" },
  { "go": "6:40:20 PM", "java": "", "c": "%X", "notes": "- writes localized time representation, e.g. 18:40:20 or 6:40:20 PM (locale dependent)" },
  { "go": "2-Jan-2006", "java": "d-MMM-YYYY", "c": "%v", "notes": "- is equivalent to \"%e-%b-%Y\"" },
  { "go": "2006", "java": "yyyy", "c": "%Y", "notes": "- writes year as a decimal number, e.g. 2017" },
  { "go": "06", "java": "yy", "c": "%y", "notes": "- writes last 2 digits of year as a decimal number (range [00,99])" },
  { "go": "January", "java": "MMMM", "c": "%B", "notes": "- writes full month name, e.g. October (locale dependent)" },
  { "go": "Jan", "java": "MMM", "c": "%b", "notes": "- writes abbreviated month name, e.g. Oct (locale dependent)" },
  { "go": "Jan", "java": "MMM", "c": "%h", "notes": "- writes abbreviated month name, e.g. Oct (locale dependent)" },
  { "go": "01", "java": "MM", "c": "%m", "notes": "- writes month as a decimal number (range [01,12])" },
  { "go": "02", "java": "dd", "c": "%d", "notes": "- writes day of the month as a decimal number (range [01,31])" },
  { "go": "2", "java": "d", "c": "%e", "notes": "- writes day of the month as a decimal number (range [1,31])." },
  { "go": "Mon", "java": "EEE", "c": "%a", "notes": "- writes abbreviated weekday name, e.g. Fri (locale dependent)" },
  { "go": "Monday", "java": "EEEE", "c": "%A", "notes": "- writes full weekday name, e.g. Friday (locale dependent)" },
  { "go": "", "java": "", "c": "%j", "notes": "- Day of the year (001-366)." },
  { "go": "", "java": "", "c": "%U", "notes": "- Week number with the first Sunday as the first day of week one (00-53)." },
  { "go": "", "java": "", "c": "%W", "notes": "- Week number with the first Monday as the first day of week one (00-53)." },
  { "go": "", "java": "", "c": "%w", "notes": "- Weekday as a decimal number with Sunday as 0 (0-6)." },
  { "go": "", "java": "", "c": "%u", "notes": "- writes weekday as a decimal number, where Monday is 1 (ISO 8601 format) (range [1-7])" },
  { "go": "15", "java": "HH", "c": "%H", "notes": "- writes hour as a decimal number, 24 hour clock (range [00-23])" },
  { "go": "3", "java": "K", "c": "%l", "notes": "-" },
  { "go": "03", "java": "KK", "c": "%I", "notes": "- writes hour as a decimal number, 12 hour clock (range [01,12])" },
  { "go": "PM", "java": "a", "c": "%p", "notes": "- writes localized a.m. or p.m. (locale dependent)" },
  { "go": "04", "java": "mm", "c": "%M", "notes": "- writes minute as a decimal number (range [00,59])" },
  { "go": "05", "java": "ss", "c": "%S", "notes": "- writes second as a decimal number (range [00,60])" },
  { "go": "-0700", "java": "Z", "c": "%z", "notes": "- writes offset from UTC in the ISO 8601 format (e.g. -0430), or no characters if the time zone information is not available" },
  { "go": "MST", "java": "z", "c": "%Z", "notes": "- writes locale-dependent time zone name or abbreviation, or no characters if the time zone information is not available" }
]`
