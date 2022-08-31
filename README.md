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

Also, since it's based on my Unify package, it has support for self-updating.


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


## Examples:

### Parsing.
Parse the date string `Sat 01 Jul 1967 09:42:42 AEST`.

	% GoWhen parse . "Sat 01 Jul 1967 09:42:42 AEST"
    1967-07-01T09:42:42+10:00

Parse the date string `1967-07-01 09:42:42` with custom format `2006-01-02 15:04:05`.

	% GoWhen parse "2006-01-02 15:04:05" "1967-07-01 09:42:42"
    1967-07-01T09:42:42Z

Print UNIX epoch

    % GoWhen parse . epoch
    1970-01-01T00:00:00Z


### Adding
Parse today's date and add `20 days`.

	% GoWhen parse . today  add "20d"
    2022-09-18T06:35:07+10:00

Parse today's date and add `2000 microseconds`.

	% GoWhen parse . now  add "2000us"

Parse today's date and add `-1 year, +12 months, -1 week, +7 days, -2 hours, +120 minutes, -2 seconds, +2000 mS`.

	% GoWhen add "-1y 12M -1w +7d -2h 120m -2s +2000ms"
    2022-08-29T06:38:36.73375+10:00


### Timezones
Convert `1967-07-01 09:42:42` to timezone `Australia/Sydney`.

	% GoWhen parse . "1967-07-01 09:42:42"  timezone "Australia/Sydney"
    1967-07-01T19:42:42+10:00

Convert `1967-07-01 09:42:42` to timezone `UTC`.

	% GoWhen parse . "1967-07-01 09:42:42"  timezone "UTC"
    1967-07-01T09:42:42Z

Convert `1967-07-01 09:42:42` to timezone `America/Chicago`.

	% GoWhen parse . "1967-07-01 09:42:42"  timezone "America/Chicago"
    1967-07-01T04:42:42-05:00

Convert `1967-07-01 09:42:42` to timezone `Iceland`.

	% GoWhen parse . "1967-07-01 09:42:42"  timezone "Iceland"
    1967-07-01T09:42:42Z


### Rounding
Round `1967-07-01 09:42:42` down to the nearest `5 minutes`.

	% GoWhen parse . "1967-07-01 09:42:42"  round down 5m
    1967-07-01T09:40:00Z

Round `1967-07-01 09:42:42` up to the nearest `1 hour`.

	% GoWhen parse . "1967-07-01 09:42:42"  round down 1h
    1967-07-01T10:00:00Z


### Differences
Show difference between `tomorrow` and `yesterday`.

	% GoWhen parse . yesterday  diff . tomorrow
    2d

Show difference between `tomorrow` and `yesterday`.

	% GoWhen parse . last-week  diff . today
    7d

Show difference between `now` and `2022-02-01 00:00:00`.

    % GoWhen parse . now diff . "2022-02-01 00:00:00"
    6M 28d 7h 7m 55s

Show difference between "2022-02-01 00:00:00" and "now".

    % GoWhen parse . "2020-01-01 00:00:00"  diff . now
    2y 7M 28d 8h 19m 36s


### Conditionals
Is the date `1967-07-07 09:42:42` after `1967-07-01 09:42:42`

    % GoWhen parse . "1967-07-01 09:42:42"   is after . "1967-07-07 09:42:42"
    NO

Is the date `1967-07-01 09:42:42` before `1967-07-07 09:42:42`

    % GoWhen parse . "1967-07-01 09:42:42"   is before . "1967-07-07 09:42:42"
    YES


### Formatting
Format `1967-07-01 09:42:42` as `Mon Jan _2 15:04:05 MST 2006`.

	% GoWhen parse . "1967-07-01 09:42:42"  format UnixDate
    Sat Jul  1 09:42:42 UTC 1967

Format `1967-07-01 09:42:42` as `2006-01-02 15:04:05`.

	% GoWhen parse . "1967-07-01 09:42:42"  format "2006-01-02 15:04:05"
    1967-07-01 09:42:42

Print current date/time as UNIX epoch, (in seconds).

    % GoWhen parse . now  format epoch
    1661754986

Print today's date as a week number.

    % parse . today  format week
    35

Print a calendar for the year of `1967-07-01`.

    % GoWhen parse . 1967-07-01  format cal-year
```
|-------------- 1967 --------------|
| Jan 1967
+----+----+----+----+----+----+----+
| S  | M  | T  | W  | T  | F  | S  |
+----+----+----+----+----+----+----+
|  1 |  2 |  3 |  4 |  5 |  6 |  7 |
|  8 |  9 | 10 | 11 | 12 | 13 | 14 |
| 15 | 16 | 17 | 18 | 19 | 20 | 21 |

...


| 19 | 20 | 21 | 22 | 23 | 24 | 25 |
| 26 | 27 | 28 | 29 | 30 |  1 |  2 |
+----+----+----+----+----+----+----+
| Dec 1967
+----+----+----+----+----+----+----+
| S  | M  | T  | W  | T  | F  | S  |
+----+----+----+----+----+----+----+
| 26 | 27 | 28 | 29 | 30 |  1 |  2 |
|  3 |  4 |  5 |  6 |  7 |  8 |  9 |
| 10 | 11 | 12 | 13 | 14 | 15 | 16 |
| 17 | 18 | 19 | 20 | 21 | 22 | 23 |
| 24 | 25 | 26 | 27 | 28 | 29 | 30 |
| 31 |  1 |  2 |  3 |  4 |  5 |  6 |
+----+----+----+----+----+----+----+
```

Print a calendar for the month of `1967-07`.

    % GoWhen parse . 1967-07  format cal-month
```
| Jul 1967
+----+----+----+----+----+----+----+
| S  | M  | T  | W  | T  | F  | S  |
+----+----+----+----+----+----+----+
| 25 | 26 | 27 | 28 | 29 | 30 |  1 |
|  2 |  3 |  4 |  5 |  6 |  7 |  8 |
|  9 | 10 | 11 | 12 | 13 | 14 | 15 |
| 16 | 17 | 18 | 19 | 20 | 21 | 22 |
| 23 | 24 | 25 | 26 | 27 | 28 | 29 |
| 30 | 31 |  1 |  2 |  3 |  4 |  5 |
+----+----+----+----+----+----+----+
```

Print a calendar for the week of `1967-07-01`.

    % GoWhen parse . 1967-07-01  format cal-week
```
| Jun 1967
+----+----+----+----+----+----+---+
| S  | M  | T  | W  | T  | F  | S |
+----+----+----+----+----+----+---+
| 25 | 26 | 27 | 28 | 29 | 30 | 1 |
+----+----+----+----+----+----+---+
```


### Ranging
Produce a set of dates from `2022-01-01 00:00:00` to `2022-01-01 10:00:00` in 1h increments.

    % GoWhen parse . "2022-01-01 00:00:00"  range . "2022-01-01 10:00:00" 1h
    2022-01-01 00:00:00
    2022-01-01 01:00:00
    2022-01-01 02:00:00
    2022-01-01 03:00:00
    2022-01-01 04:00:00
    2022-01-01 05:00:00
    2022-01-01 06:00:00
    2022-01-01 07:00:00
    2022-01-01 08:00:00
    2022-01-01 09:00:00

Produce a set of dates from `2022-01-01` to `2022-01-01` in 1h increments.

    % GoWhen parse . "2000-01-01"  range . "2022-01-01" 365d
    2000-01-01
    2000-12-31
    2001-12-31
    2002-12-31
    2003-12-31
    2004-12-30
    2005-12-30
    2006-12-30
    2007-12-30
    2008-12-29
    2009-12-29
    2010-12-29
    2011-12-29
    2012-12-28
    2013-12-28
    2014-12-28
    2015-12-28
    2016-12-27
    2017-12-27
    2018-12-27
    2019-12-27
    2020-12-26
    2021-12-26

Produce a reverse set of dates from `2022-01-01 00:00:00` to `2000-01-01 00:00:00` in increments of `1 year, 2 months, 3 weeks, 4 days, 5 hours, 6 minutes and 7 seconds`.

    % GoWhen parse . "2022-01-01 00:00:00"  range . "2000-01-01 00:00:00" "1y 2M 3w 4d 5h 6m 7s"
    2022-01-01 00:00:00
    2020-10-06 18:53:53
    2019-07-12 13:47:46
    2018-04-17 08:41:39
    2017-01-23 03:35:32
    2015-10-28 22:29:25
    2014-08-03 17:23:18
    2013-05-09 12:17:11
    2012-02-13 07:11:04
    2010-11-18 02:04:57
    2009-08-23 20:58:50
    2008-05-29 15:52:43
    2007-03-04 10:46:36
    2005-12-10 05:40:29
    2004-09-15 00:34:22
    2003-06-19 19:28:15
    2002-03-25 14:22:08
    2000-12-31 09:16:01

Produce a reverse set of dates from `2000-07-01` to `UNIX epoch` in increments of `1 year` using format `2006-01-02 - Monday`.
IE: Show days that a particular date falls on.

    % GoWhen parse . "2000-07-01"  range "2006-01-02 - Monday" "epoch" 1y
    2000-07-01 - Saturday
    1999-07-01 - Thursday
    1998-07-01 - Wednesday
    1997-07-01 - Tuesday
    1996-07-01 - Monday
    1995-07-01 - Saturday
    1994-07-01 - Friday
    1993-07-01 - Thursday
    1992-07-01 - Wednesday
    1991-07-01 - Monday
    1990-07-01 - Sunday
    1989-07-01 - Saturday
    1988-07-01 - Friday
    1987-07-01 - Wednesday
    1986-07-01 - Tuesday
    1985-07-01 - Monday
    1984-07-01 - Sunday
    1983-07-01 - Friday
    1982-07-01 - Thursday
    1981-07-01 - Wednesday
    1980-07-01 - Tuesday
    1979-07-01 - Sunday
    1978-07-01 - Saturday
    1977-07-01 - Friday
    1976-07-01 - Thursday
    1975-07-01 - Tuesday
    1974-07-01 - Monday
    1973-07-01 - Sunday
    1972-07-01 - Saturday
    1971-07-01 - Thursday
    1970-07-01 - Wednesday


### Stacking
Parse the date `Sat 01 Jul 1967 09:42:42 AEST`, add `20 days` and print as format `20060102/20060102_150405-webcam.jpg`.

	% GoWhen parse . "Sat 01 Jul 1967 09:42:42 AEST"   add "20d"  format "20060102/20060102_150405-webcam.jpg"
    19670820/19670820_094242-webcam.jpg

Parse the date `Sat 01 Jul 1967 09:42:42 AEST`, convert to timezone `Iceland`, add `1 day`, round down to every `5 minutes` and print as format `2006-01-02 15:04:05`.

	% GoWhen parse . "Sat 01 Jul 1967 09:42:42 AEST"   timezone Iceland  add 1d  round down 5m  format "2006-01-02 15:04:05"
    1967-07-02 09:40:00

Is the date `1967-07-07 09:42:42` after `1967-07-01 09:42:42`

    % GoWhen parse . "1967-07-01 09:42:42"   is after . "1967-07-07 09:42:42"
    NO

Is the date `1967-07-01 09:42:42` before `1967-07-07 09:42:42`

    % GoWhen parse . "1967-07-01 09:42:42"   is before . "1967-07-07 09:42:42"
    YES


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
