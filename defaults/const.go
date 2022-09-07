package defaults

import _ "embed"

// Need to execute `go generate -v -x defaults/const.go` OR `go generate -v -x ./...`
//go:generate cp ../README.md README.md
//go:generate cp ../EXAMPLES.md EXAMPLES.md

//go:embed README.md
var Readme string

//go:embed EXAMPLES.md
var Examples string

const (
	Description   = "GoWhen - CLI based Date/Time manipulation written in GoLang"
	BinaryName    = "GoWhen"
	BinaryVersion = "1.0.3"
	SourceRepo    = "github.com/MickMake/" + BinaryName
	BinaryRepo    = "github.com/MickMake/" + BinaryName

	EnvPrefix = "GOWHEN"

	// HelpTemplate Extended help...
	HelpTemplate = `
DefaultBinaryName - Extended help.

### Input
	% GoWhen parse <format> <date/time>

### Modify
	% GoWhen add <duration>

	% GoWhen timezone <zone>
	% GoWhen tz <zone>

	% GoWhen round up <duration>
	% GoWhen round down <duration>

### Output
	% GoWhen format <format | cal-year | cal-month | cal-week>

	% GoWhen is dst
	% GoWhen is leap
	% GoWhen is weekday
	% GoWhen is weekend
	% GoWhen is before <format> <date/time>
	% GoWhen is after <format> <date/time>

	% GoWhen diff <format> <date/time>

	% GoWhen range <format> <to date/time> <duration>


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

`
)
