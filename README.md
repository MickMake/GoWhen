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

What is planned for future releases:
- Support for more parse formats, (Java and C).

Also, since it's based on my Unify package, it has support for self-updating.


## Command summary
Note: all commands are stackable. Except `format` and `is` - doesn't make any sense to make them stackable.

### Parsing.
	% GoWhen parse <format | .> <date/time>

### Adding
	% GoWhen add <duration>

### Timezones
	% GoWhen timezone <zone>
	% GoWhen tz <zone>

### Rounding
	% GoWhen round up <duration>
	% GoWhen round down <duration>

### Formatting
	% GoWhen format <format | cal-year | cal-month | cal-week | .>

### Conditionals
	% GoWhen is dst
	% GoWhen is leap
	% GoWhen is weekday
	% GoWhen is weekend
	% GoWhen is before <format | .> <date/time>
	% GoWhen is after <format | .> <date/time>

### Difference
	% GoWhen diff <format | .> <date/time>

### Ranging
	% GoWhen range <format | .> <to date/time> <duration>


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
    Epoch       = Unix epoch
    Week        = Week number of the year.

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
Special date strings.

    "" / now today - Today's date/time.
    tomorrow - 
    yesterday - 
    next-week - 
    last-week - 
    epoch - UNIX epoch, (1970-01-01 00:00:00).


## Examples:

### Parsing.
Parse the date string `Sat 01 Jul 1967 09:42:42 AEST`.

	% GoWhen parse "Sat 31 Jul 1967 09:42:42 AEST" ""
    1967-07-31T09:42:42+10:00

Parse the date string `1967-07-01 09:42:42` with custom format `2006-01-02 15:04:05`.

	% GoWhen parse "1967-07-01 09:42:42" "2006-01-02 15:04:05"
    1967-07-01T09:42:42Z

Print UNIX epoch

    % GoWhen parse "epoch" ""
    1970-01-01T00:00:00Z


### Adding
Parse today's date and add `20 days`.

	% GoWhen parse "today" "" add "20d"
    2022-09-18T06:35:07+10:00

Parse today's date and add `2000 microseconds`.

	% GoWhen parse "now" "" add "2000us"

Parse today's date and add `-1 year, +12 months, -1 week, +7 days, -2 hours, +120 minutes, -2 seconds, +2000 mS`.

	% GoWhen add -- "-1y 12M -1w +7d -2h 120m -2s +2000ms"
    2022-08-29T06:38:36.73375+10:00


### Timezones
Convert `1967-07-01 09:42:42` to timezone `Australia/Sydney`.

	% GoWhen parse "1967-07-01 09:42:42" "" timezone "Australia/Sydney"
    1967-07-01T19:42:42+10:00

Convert `1967-07-01 09:42:42` to timezone `UTC`.

	% GoWhen parse "1967-07-01 09:42:42" "" timezone "UTC"
    1967-07-01T09:42:42Z

Convert `1967-07-01 09:42:42` to timezone `America/Chicago`.

	% GoWhen parse "1967-07-01 09:42:42" "" timezone "America/Chicago"
    1967-07-01T04:42:42-05:00

Convert `1967-07-01 09:42:42` to timezone `Iceland`.

	% GoWhen parse "1967-07-01 09:42:42" "" timezone "Iceland"
    1967-07-01T09:42:42Z


### Rounding
Round `1967-07-01 09:42:42` down to the nearest `5 minutes`.

	% GoWhen parse "1967-07-01 09:42:42" "" round down 5m
    1967-07-01T09:40:00Z

Round `1967-07-01 09:42:42` up to the nearest `1 hour`.

	% GoWhen parse "1967-07-01 09:42:42" "" round down 1h
    1967-07-01T10:00:00Z


### Differences
Show difference between `tomorrow` and `yesterday`.

	% GoWhen parse "yesterday" . diff "tomorrow" .
    2d

Show difference between `tomorrow` and `yesterday`.

	% GoWhen parse "last-week" . diff "today" .
    7d

Show difference between `now` and `2022-02-01 00:00:00`.

    % GoWhen parse "now" . diff "2022-02-01 00:00:00" .
    6M 28d 7h 7m 55s

Show difference between "2022-02-01 00:00:00" and "now".

    % GoWhen parse "2020-01-01 00:00:00" . diff now .
    2y 7M 28d 8h 19m 36s


### Formatting
Format `1967-07-01 09:42:42` as `Mon Jan _2 15:04:05 MST 2006`.

	% GoWhen parse "1967-07-01 09:42:42" . format UnixDate
    Sat Jul  1 09:42:42 UTC 1967

Format `1967-07-01 09:42:42` as `2006-01-02 15:04:05`.

	% GoWhen parse "1967-07-01 09:42:42" . format "2006-01-02 15:04:05"
    1967-07-01 09:42:42

Print current date/time as UNIX epoch, (in seconds).

    % GoWhen parse now . format epoch
    1661754986

Print today's date as a week number.

    % parse today "" format week
    35

Print a calendar for the year of `1967-07-01`.

    % GoWhen parse 1967-07-01 . format cal-year
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

    % GoWhen parse 1967-07 . format cal-month
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

    % GoWhen parse 1967-07-01 . format cal-week
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

    % GoWhen parse . "2022-01-01 00:00:00" range . "2022-01-01 10:00:00" 1h
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

    % GoWhen parse . "2000-01-01" range . "2022-01-01" 365d
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
Parse the date `Sat 31 Jul 1967 09:42:42 AEST`, add `20 days` and print as format `20060102/20060102_150405-webcam.jpg`.

	% GoWhen parse "Sat 31 Jul 1967 09:42:42 AEST" ""  add "20d"  format "20060102/20060102_150405-webcam.jpg"
    19670820/19670820_094242-webcam.jpg

Parse the date `Sat 31 Jul 1967 09:42:42 AEST`, convert to timezone `Iceland`, add `1 day`, round down to every `5 minutes` and print as format `2006-01-02 15:04:05`.

	% GoWhen parse "1967-07-01 09:42:42" ""  timezone Iceland  add 1d  round down 5m  format "2006-01-02 15:04:05"
    1967-07-02 09:40:00

Is the date `1967-07-07 09:42:42` after `1967-07-01 09:42:42`

    % GoWhen parse "1967-07-01 09:42:42" ""  is after "1967-07-07 09:42:42" ""
    NO

Is the date `1967-07-01 09:42:42` before `1967-07-07 09:42:42`

    % GoWhen parse "1967-07-01 09:42:42" ""  is before "1967-07-07 09:42:42" ""
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
