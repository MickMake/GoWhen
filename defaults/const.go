package defaults

const (
	Description   = "GoWhen - CLI based Date/Time manipulation written in GoLang"
	BinaryName    = "GoWhen"
	BinaryVersion = "1.0.0"
	SourceRepo    = "github.com/MickMake/" + BinaryName
	BinaryRepo    = "github.com/MickMake/" + BinaryName

	EnvPrefix = "GOWHEN"

	// HelpTemplate Extended help...
	HelpTemplate = `
DefaultBinaryName - Extended help.

### Parsing.
	% DefaultBinaryName parse <format | .> <date/time>

### Adding
	% DefaultBinaryName add <duration>

### Timezones
	% DefaultBinaryName timezone <zone>
	% DefaultBinaryName tz <zone>

### Rounding
	% DefaultBinaryName round up <duration>
	% DefaultBinaryName round down <duration>

### Formatting
	% DefaultBinaryName format <format | cal-year | cal-month | cal-week | .>

### Conditionals
	% DefaultBinaryName is dst
	% DefaultBinaryName is leap
	% DefaultBinaryName is weekday
	% DefaultBinaryName is weekend
	% DefaultBinaryName is before <format | .> <date/time>
	% DefaultBinaryName is after <format | .> <date/time>

### Difference
	% DefaultBinaryName diff <format | .> <date/time>

### Ranging
	% DefaultBinaryName range <format | .> <from date/time> <to date/time>


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
