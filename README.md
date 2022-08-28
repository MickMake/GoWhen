# GoWhen - CLI based Date/Time manipulation written in GoLang.

This tool came about because I needed a cross-platform way of performing date and time manipulations within scripts.

This tool does several things:
- parse date - Parse a date/time string in multiple formats.
- parse format - Provide user selectable date/time parse format.
- add - Add a date/time duration to a date/time.
- timezone - Convert between timezones.
- round - Rounding of date/time.
- format - Print date/time in a user selectable format.
- is dst - Is date/time within DST or not.
- is leap - Is date/time a leap-year or not.
- is weekend - Is date/time a weekend or not.
- is weekday - Is date/time a weekday or not.

What is planned for future releases:
- before - Is date/time before a specified date/time.
- after - Is date/time after a specified date/time.
- epoch - Format date/time as UNIX epoch.
- iso-week - Format date/time as ISO 8601.
- diff - Return date/time duration from a specified date/time.
- cal - Produce a traditional calendar in multiple formats between two dates.
- Support for more parse formats.

Also, since it's based on my Unify package, it also has support for self-updating.


## Use case example:
Note: all commands are stackable.


### Parsing.
Parse the date string `Sat 01 Jul 1967 09:42:42 AEST`.

	% parse date "Sat 31 Jul 1967 09:42:42 AEST"
    1967-07-31T09:42:42+10:00

Parse the date string `1967-07-01 09:42:42`.

	% parse date "1967-07-01 09:42:42"
    1967-07-01T09:42:42Z


### Adding
Parse today's date and add `20 days`.

	% parse date "today" add "20d"
    2022-09-18T06:35:07+10:00

Parse today's date and add `2000 microseconds`.

	% parse date "now" add "2000us"

Parse today's date and add `-1 year, +12 months, -1 week, +7 days, -2 hours, +120 minutes, -2 seconds, +2000 mS`.

	% add -- '-1y 12M -1w +7d -2h 120m -2s +2000ms'
    2022-08-29T06:38:36.73375+10:00


### Timezones
Convert `1967-07-01 09:42:42` to timezone `Australia/Sydney`.

	% parse date '1967-07-01 09:42:42' timezone 'Australia/Sydney'
    1967-07-01T19:42:42+10:00

Convert `1967-07-01 09:42:42` to timezone `UTC`.

	% parse date '1967-07-01 09:42:42' timezone 'UTC'
    1967-07-01T09:42:42Z

Convert `1967-07-01 09:42:42` to timezone `America/Chicago`.

	% parse date '1967-07-01 09:42:42' timezone 'America/Chicago'
    1967-07-01T04:42:42-05:00

Convert `1967-07-01 09:42:42` to timezone `Iceland`.

	% parse date '1967-07-01 09:42:42' timezone 'Iceland'
    1967-07-01T09:42:42Z


### Rounding
Round `1967-07-01 09:42:42` down to the nearest `5 minutes`.

	% parse date '1967-07-01 09:42:42' round down 5m
    1967-07-01T09:40:00Z

Round `1967-07-01 09:42:42` up to the nearest `1 hour`.

	% parse date '1967-07-01 09:42:42' round down 1h
    1967-07-01T10:00:00Z


### Formatting
Format `1967-07-01 09:42:42` as `Mon Jan _2 15:04:05 MST 2006`.

	% parse date '1967-07-01 09:42:42' format UnixDate
    Sat Jul  1 09:42:42 UTC 1967

Format `1967-07-01 09:42:42` as `2006-01-02 15:04:05`.

	% parse date '1967-07-01 09:42:42' format '2006-01-02 15:04:05'
    1967-07-01 09:42:42


### Stacking
Parse the date `Sat 31 Jul 1967 09:42:42 AEST`, add `20 days` and print as format `20060102/20060102_150405-webcam.jpg`.

	% parse date "Sat 31 Jul 1967 09:42:42 AEST" add "20d" format "20060102/20060102_150405-webcam.jpg"
    19670820/19670820_094242-webcam.jpg

Parse the date `Sat 31 Jul 1967 09:42:42 AEST`, convert to timezone `Iceland`, add `1 day`, round down to every `5 minutes` and print as format `2006-01-02 15:04:05`.

	% parse date '1967-07-01 09:42:42' timezone Iceland add 1d round down 5m format '2006-01-02 15:04:05'
    1967-07-02 09:40:00


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
