# GoWhen - CLI based Date/Time manipulation written in GoLang.

This tool came about because I needed a cross-platform way of performing date and time manipulations within scripts.

This tool does several things:
- parse - Parse a date/time string.
- add - Add a date/time duration to a date/time.
- timezone - Convert between timezones.
- round - Rounding of date/time.
- format - Print date/time in a user selectable format.
- is dst - Is date/time within DST or not.
- is leap - Is date/time a leap-year or not.
- is weekend - Is date/time a weekend or not.
- is weekday - Is date/time a weekday or not.
- is before - Is date/time before a specified date/time.
- is after - Is date/time after a specified date/time.
- diff - Return date/time duration from a specified date/time.
- cal - Produce a traditional calendar in multiple formats.
- range - Produce a range of dates with variable duration span between.
- Support for more parse formats, (Java and C), using a simple JSON mapping file.
- Can run as an interactive shell.

Also, since it's based on my Unify package, it has support for self-updating.

Planned enhancements:
- Ability to define command aliases. EG: `epoch` = `parse . epoch` or `christmas` = `parse . now diff . "2022-12-25 00:00:00"`


## Command summary
Note: all commands are stackable. Except `format` and `is` - doesn't make any sense to make them stackable.

### Date input
	% GoWhen parse <format> <date/time>

### Date modify
	% GoWhen add <duration>

	% GoWhen timezone <zone>
	% GoWhen tz <zone>

	% GoWhen round up <duration>
	% GoWhen round down <duration>

### Output
	% GoWhen format <format | cal-year | cal-month | cal-week | .>

	% GoWhen is dst
	% GoWhen is leap
	% GoWhen is weekday
	% GoWhen is weekend
	% GoWhen is before <format> <date/time>
	% GoWhen is after <format> <date/time>

	% GoWhen diff <format> <date/time>

	% GoWhen range <format> <to date/time> <duration>


## Quick start
Impatient? OK, so am I. Here's some things you can do with `GoWhen`.

Parse today's date and add `-1 year, +12 months, -1 week, +7 days, -2 hours, +120 minutes, -2 seconds, +2000 mS`.

	% GoWhen add "-1y 12M -1w +7d -2h 120m -2s +2000ms"
    2022-08-29T06:38:36.73375+10:00

Show difference between "2022-02-01 00:00:00" and "now".

    % GoWhen parse . "2020-01-01 00:00:00"  diff . now
    2y 7M 28d 8h 19m 36s

Produce a reverse set of dates from `2010-01-01 00:00:00` to `2000-01-01 00:00:00` in increments of `1 year, 2 months, 3 weeks, 4 days, 5 hours, 6 minutes and 7 seconds`.

    % GoWhen parse . "2010-01-01 00:00:00"  range . "2000-01-01 00:00:00" "1y 2M 3w 4d 5h 6m 7s"
    2010-01-01 00:00:00
    2008-10-06 18:53:53
    2007-07-12 13:47:46
    2006-04-17 08:41:39
    2005-01-23 03:35:32
    2003-10-28 22:29:25
    2002-08-03 17:23:18
    2001-05-09 12:17:11
    2000-02-13 07:11:04

Parse the date `01 Jul 1967 09:42:42 AEST`, convert to timezone `Iceland`, add `1 day`, round down to every `30 minutes` and print as format `ANSIC`, (`Mon Jan _2 15:04:05 2006`).

	% GoWhen parse . "01 Jul 1967 09:42:42 AEST"   timezone Iceland  add 1d  round down 30m  format ANSIC
    1967-07-02 09:40:00

Show the difference between two dates, `1920-09-08` and `2023-07-01`.

    % GoWhen parse . 1920-09-08 diff . 2023-07-01
    102y 9M 23d

Or in reverse...

    % GoWhen parse . 1920-09-08 add "102y 9M 23d"
    2023-07-01

Is the year 2000 a leap year?

    % GoWhen parse . 2000 is leap
    YES

Is the date `1967-07-01` a weekend?

    % GoWhen parse . 1967-07-01 is weekend
    YES

Produce a list of files with names based on `%Y%m%d_%H%M%S-webcam.jpg` from `01 Jul 1967 09:42:42 AEST` until `1 week` in `half day` increments.

    % GoWhen parse . "01 Jul 1967 09:42:42 AEST"   add "1w"  range "%Y%m%d_%H%M%S-webcam.jpg" . .5d
    19670701_094242-webcam.jpg
    19670701_214242-webcam.jpg
    19670801_094242-webcam.jpg
    19670801_214242-webcam.jpg
    19670802_094242-webcam.jpg
    19670802_214242-webcam.jpg
    19670803_094242-webcam.jpg
    19670803_214242-webcam.jpg
    19670804_094242-webcam.jpg
    19670804_214242-webcam.jpg
    19670805_094242-webcam.jpg
    19670805_214242-webcam.jpg
    19670806_094242-webcam.jpg
    19670806_214242-webcam.jpg


## Further Examples
[EXAMPLES](https://github.com/MickMake/GoWhen/blob/master/EXAMPLES.md)


## Formats

### Print / Parse formats
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
	Stamp       = "Jan _2 15:04:05"
	StampMilli  = "Jan _2 15:04:05.000"
	StampMicro  = "Jan _2 15:04:05.000000"
	StampNano   = "Jan _2 15:04:05.000000000"

### Additional print formats
    epoch       = Unix epoch
    week        = Week number of the year.
    cal-week    = Produce a week long calendar.
    cal-month   = Produce a monthly calendar.
    cal-year    = Produce a full year calendar.

### Additional parse formats
	.			= Best guess input string.

### Add/round durations
	ns - Nanosecond
	us - microsecond
	ms - Millisecond
	s - Second
	m - Minute
	h - Hour
	d - Day
	w - Week
	M - Month
	y - Year

### Date parsing
Special date entry strings.

    epoch       = Sets the date to `1970-01-01 00:00:00`.
    now         = Today's date.
    today       = Today's date.
    tomorrow    = 
    yesterday   = 
    last-week   = 
    next-week   = 


## Date/time format conversion
This tool now supports date/time formats for various languages. It uses a simple JSON file to build up maps of conversion rules.
The default mapping is as follows.

To select a different date/time layout. Do one of the following:

    % GoWhen config write --format=go
    % GoWhen config write --format=cpp
    % GoWhen config write --format=java

```
Conversion table:
+------------------------------+----------------------------+--------------------+--------------------------------+
|        GOLANG LAYOUT         |       JAVA NOTATION        |   C/CPP NOTATION   |             NOTES              |
+------------------------------+----------------------------+--------------------+--------------------------------+
| 20060102                     | yyyyMMdd                   | %Y%m%d             | ISO 8601                       |
| January 02, 2006             | MMMM dd, yyyy              | %B %d, %Y          |                                |
| 02 January 2006              | dd MMMM yyyy               | %d %B %Y           |                                |
| 02-Jan-2006                  | dd-MMM-yyyy                | %d-%b-%Y           |                                |
| 01/02/2006                   | MM/dd/yyyy                 | %m/%d/%Y           | US                             |
| 010206                       | MMddyy                     | %m%d%y             | US                             |
| Jan-02-06                    | MMM-dd-yy                  | %b-%d-%y           | US                             |
| Jan-02-2006                  | MMM-dd-yyyy                | %b-%d-%Y           | US                             |
| 3:04 PM                      | K:mm a                     | %l:%M %p           | US                             |
| 2006-01-02T15:04:05          | yyyy-MM-dd'T'HH:mm:ss      | %FT%T              | ISO 8601                       |
| 2006-01-02T15:04:05-0700     | yyyy-MM-dd'T'HH:mm:ssZ     | %FT%T%z            | ISO 8601                       |
| 2 Jan 2006 15:04:05          | d MMM yyyy HH:mm:ss        | %e %b %Y %T        |                                |
| 2 Jan 2006 15:04             | d MMM yyyy HH:mm           | %e %b %Y %R        |                                |
| Mon, 2 Jan 2006 15:04:05 MST | EEE, d MMM yyyy HH:mm:ss z | %a, %e %b %Y %T %Z | RFC 1123 RFC 822               |
| 01/02/06                     | MM/dd/yy                   | %D                 | equivalent to "%m/%d/%y"       |
| 15:04                        | HH:mm                      | %R                 | equivalent to "%H:%M"          |
| 15:04:05                     | HH:mm:ss                   | %T                 | equivalent to "%H:%M:%S" (the  |
|                              |                            |                    | ISO 8601 time format)          |
| 03:04:05 PM                  | KK:mm:ss a                 | %r                 | writes localized 12-hour clock |
|                              |                            |                    | time (locale dependent)        |
| 2006-01-02                   | yyyy-MM-dd                 | %F                 | equivalent to "%Y-%m-%d" (the  |
|                              |                            |                    | ISO 8601 date format)          |
| Jan 02 15:04:05 2006         |                            | %c                 | writes standard date and       |
|                              |                            |                    | time string, e.g. Sun Oct      |
|                              |                            |                    | 17 04:41:13 2010 (locale       |
|                              |                            |                    | dependent)                     |
|                              |                            | %C                 | Year divided by 100 and        |
|                              |                            |                    | truncated to integer (00-99)   |
|                              |                            | %x                 | writes localized date          |
|                              |                            |                    | representation (locale         |
|                              |                            |                    | dependent)                     |
| 6:40:20 PM                   |                            | %X                 | writes localized time          |
|                              |                            |                    | representation, e.g. 18:40:20  |
|                              |                            |                    | or 6:40:20 PM (locale          |
|                              |                            |                    | dependent)                     |
| 2-Jan-2006                   | d-MMM-YYYY                 | %v                 | is equivalent to "%e-%b-%Y"    |
| 2006                         | yyyy                       | %Y                 | writes year as a decimal       |
|                              |                            |                    | number, e.g. 2017              |
| 06                           | yy                         | %y                 | writes last 2 digits of year   |
|                              |                            |                    | as a decimal number (range     |
|                              |                            |                    | [00,99])                       |
| January                      | MMMM                       | %B                 | writes full month name, e.g.   |
|                              |                            |                    | October (locale dependent)     |
| Jan                          | MMM                        | %b                 | writes abbreviated month name, |
|                              |                            |                    | e.g. Oct (locale dependent)    |
| Jan                          | MMM                        | %h                 | writes abbreviated month name, |
|                              |                            |                    | e.g. Oct (locale dependent)    |
| 01                           | MM                         | %m                 | writes month as a decimal      |
|                              |                            |                    | number (range [01,12])         |
| 02                           | dd                         | %d                 | writes day of the month as a   |
|                              |                            |                    | decimal number (range [01,31]) |
| 2                            | d                          | %e                 | writes day of the month as a   |
|                              |                            |                    | decimal number (range [1,31]). |
| Mon                          | EEE                        | %a                 | writes abbreviated weekday     |
|                              |                            |                    | name, e.g. Fri (locale         |
|                              |                            |                    | dependent)                     |
| Monday                       | EEEE                       | %A                 | writes full weekday name, e.g. |
|                              |                            |                    | Friday (locale dependent)      |
|                              |                            | %j                 | Day of the year (001-366).     |
|                              |                            | %U                 | Week number with the first     |
|                              |                            |                    | Sunday as the first day of     |
|                              |                            |                    | week one (00-53).              |
|                              |                            | %W                 | Week number with the first     |
|                              |                            |                    | Monday as the first day of     |
|                              |                            |                    | week one (00-53).              |
|                              |                            | %w                 | Weekday as a decimal number    |
|                              |                            |                    | with Sunday as 0 (0-6).        |
|                              |                            | %u                 | writes weekday as a decimal    |
|                              |                            |                    | number, where Monday is 1 (ISO |
|                              |                            |                    | 8601 format) (range [1-7])     |
| 15                           | HH                         | %H                 | writes hour as a decimal       |
|                              |                            |                    | number, 24 hour clock (range   |
|                              |                            |                    | [00-23])                       |
| 3                            | K                          | %l                 | -                              |
| 03                           | KK                         | %I                 | writes hour as a decimal       |
|                              |                            |                    | number, 12 hour clock (range   |
|                              |                            |                    | [01,12])                       |
| PM                           | a                          | %p                 | writes localized a.m. or p.m.  |
|                              |                            |                    | (locale dependent)             |
| 04                           | mm                         | %M                 | writes minute as a decimal     |
|                              |                            |                    | number (range [00,59])         |
| 05                           | ss                         | %S                 | writes second as a decimal     |
|                              |                            |                    | number (range [00,60])         |
| -0700                        | Z                          | %z                 | writes offset from UTC in the  |
|                              |                            |                    | ISO 8601 format (e.g. -0430),  |
|                              |                            |                    | or no characters if the        |
|                              |                            |                    | time zone information is not   |
|                              |                            |                    | available                      |
| MST                          | z                          | %Z                 | writes locale-dependent time   |
|                              |                            |                    | zone name or abbreviation, or  |
|                              |                            |                    | no characters if the time zone |
|                              |                            |                    | information is not available   |
+------------------------------+----------------------------+--------------------+--------------------------------+
```


## Config file.
```
+-----------+------------+----------------+-------------------------------+---------------------+
|   FLAG    | SHORT FLAG |  ENVIRONMENT   |          DESCRIPTION          | VALUE (* = DEFAULT) |
+-----------+------------+----------------+-------------------------------+---------------------+
| --config  |            | GOWHEN_CONFIG  | GoWhen: config file.          |  *                  |
| --debug   |            | GOWHEN_DEBUG   | GoWhen: Debug mode.           | false *             |
| --quiet   |            | GOWHEN_QUIET   | GoWhen: Silence all messages. | false *             |
| --timeout |            | GOWHEN_TIMEOUT | Web timeout.                  | 30s *               |
+-----------+------------+----------------+-------------------------------+---------------------+
```
