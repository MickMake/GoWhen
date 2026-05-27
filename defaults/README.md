# GoWhen - CLI based Date/Time manipulation written in GoLang.

This tool came about because I needed a cross-platform way of performing date and time manipulations within scripts.

This tool does several things:
- parse - Parse a date/time string.
- add - Add a date/time duration to a date/time.
- timezone - Convert between timezones.
- round - Rounding of date/time.
- format - Print date/time in a user selectable format.
- keep - Keep output only when the working date matches a selector.
- drop - Drop output when the working date matches a selector.
- alias - Define reusable command aliases.
- is dst - Is date/time within DST or not.
- is leap - Is date/time a leap-year or not.
- is weekend - Is date/time a weekend or not.
- is weekday - Is date/time a weekday or not.
- is before - Is date/time before a specified date/time.
- is after - Is date/time after a specified date/time.
- diff - Return date/time duration from a specified date/time.
- cal - Produce a traditional calendar in multiple formats.
- range - Produce a range of dates with variable duration span between.
- Automatically process piped stdin, one working date per input line.
- Support for more parse formats, (Java and C), using a simple JSON mapping file.
- Can run as an interactive shell.

Also, since it's based on my Unify package, it has support for self-updating.

## Command summary
Note: all commands are stackable. Except `format` and `is` - doesn't make any sense to make them stackable.

### Date input
	% GoWhen parse <format> <date/time>

When stdin is piped, each non-empty line is automatically parsed as the working date before the command chain runs.

	% cat dates.txt | GoWhen format 2006-01-02

### Date modify
	% GoWhen add <duration>

	% GoWhen timezone <zone>
	% GoWhen tz <zone>

	% GoWhen round up <duration>
	% GoWhen round down <duration>

### Date filters
	% GoWhen keep <selector>
	% GoWhen drop <selector>

### Aliases
	% GoWhen alias add <name> <cmd> ...
	% GoWhen alias list
	% GoWhen alias del <name>

Aliases are saved to `aliases.json` under the GoWhen config directory and loaded before command execution.

Examples:

	% GoWhen alias add christmas parse . 2026-12-25
	% GoWhen christmas format 2006-01-02
	% GoWhen alias list
	% GoWhen alias del christmas

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

## Filters and piped stdin

GoWhen can act as a shell filter. When stdin is piped, each non-empty input line is automatically parsed as the working date before the command chain runs.

```sh
cat dates.txt | GoWhen format 2006-01-02
```

This is equivalent to running the command chain once for each line, with a hidden parse step first:

```text
parse . <stdin-line> format 2006-01-02
```

Commands still keep their normal argument rules. Piped stdin supplies the working date; it does not fill missing command arguments.

For example, this compares each input date against a fixed command-line date:

```sh
cat dates.txt | GoWhen diff . 2026-06-01
```

But this is still invalid, because `diff` requires both arguments:

```sh
cat dates.txt | GoWhen diff .
```

### keep and drop

`keep` and `drop` filter output based on the current working date.

```sh
GoWhen keep <selector>
GoWhen drop <selector>
```

`keep` prints only dates matching the selector.

```sh
GoWhen parse . 2026-01-01 range . 2026-02-01 1d keep monday
```

`drop` suppresses dates matching the selector.

```sh
GoWhen parse . 2026-01-01 range . 2026-02-01 1d drop weekend
```

They also work with piped input:

```sh
printf '%s\n' 2026-05-30 2026-05-31 2026-06-01 | GoWhen drop weekend format 2006-01-02
```

Output:

```text
2026-06-01
```

### Selectors

Selector names are case-insensitive. Spaces, underscores, and hyphens are treated the same, so these are equivalent:

```text
month end
month_end
month-end
```

Current selectors are:

```text
weekday | weekdays | workday | workdays | business-day | business-days
weekend | weekends

mon | monday
tue | tuesday
wed | wednesday
thu | thursday
fri | friday
sat | saturday
sun | sunday

bom | month-start | month-begin
eom | month-end
boy | year-start | year-begin
eoy | year-end
quarter-start | quarter-begin
quarter-end

day-1 ... day-31

first-mon ... first-sun
first-monday ... first-sunday
second-mon ... second-sun
second-monday ... second-sunday
third-mon ... third-sun
third-monday ... third-sunday
fourth-mon ... fourth-sun
fourth-monday ... fourth-sunday
last-mon ... last-sun
last-monday ... last-sunday
```

`business-day` and `business-days` currently mean Monday-Friday only. Public holidays are not considered.

### Filter examples

Keep weekdays from piped input:

```sh
cat dates.txt | GoWhen keep weekday format 2006-01-02
```

Drop weekends from piped input:

```sh
cat dates.txt | GoWhen drop weekend format 2006-01-02
```

Keep Mondays from a generated range:

```sh
GoWhen parse . 2026-01-01 range . 2026-02-01 1d keep mon
```

Drop Saturdays from a generated range:

```sh
GoWhen parse . 2026-01-01 range . 2026-02-01 1d drop saturday
```

Keep month-end dates from a generated range:

```sh
GoWhen parse . 2026-01-01 range . 2026-12-31 1d keep "month end"
```

Keep quarter-end dates from a generated range:

```sh
GoWhen parse . 2026-01-01 range . 2026-12-31 1d keep "quarter end"
```

Keep the 15th day of each month from a generated range:

```sh
GoWhen parse . 2026-01-01 range . 2026-12-31 1d keep day-15
```

Keep the last Friday of each month from a generated range:

```sh
GoWhen parse . 2026-01-01 range . 2026-12-31 1d keep "last friday"
```

Keep today only if today is a weekday:

```sh
GoWhen keep weekday
```

Drop today if today is a weekend:

```sh
GoWhen drop weekend
```

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
	.            = Best guess input string.

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
The full conversion table is maintained in `README.md` and can be refreshed into this embedded copy with:

```sh
go generate ./...
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
