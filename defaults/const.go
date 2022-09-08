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
	BinaryVersion = "1.0.5"
	SourceRepo    = "github.com/MickMake/" + BinaryName
	BinaryRepo    = "github.com/MickMake/" + BinaryName

	EnvPrefix = "GOWHEN"

	HelpSummary = `
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

`
)
