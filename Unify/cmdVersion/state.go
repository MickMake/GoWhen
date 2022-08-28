package cmdVersion

import (
	"GoWhen/Unify/Only"
	"errors"
	"fmt"
	"github.com/logrusorgru/aurora"
	"os"
)

type State struct {
	error   error
	warning error
	ok      error
}

func (s *State) IsOk() bool {
	if s.ok == nil {
		return true
	}
	if s.ok.Error() == "" {
		return true
	}
	return false
}
func (s *State) IsNotOk() bool {
	return !s.IsOk()
}

func (s *State) IsError() bool {
	if s == nil {
		return false
	}
	if s.error == nil {
		return false
	}
	if s.error.Error() == "" {
		return false
	}
	return true
}
func (s *State) IsNotError() bool {
	return !s.IsError()
}

func (s *State) SetError(format string, args ...interface{}) {
	s.error = errors.New(fmt.Sprintf(format, args...))
	s.warning = nil
	s.ok = nil
}
func (s *State) GetError() error {
	return s.error
}

func (s *State) SetWarning(format string, args ...interface{}) {
	s.error = nil
	s.warning = errors.New(fmt.Sprintf(format, args...))
	s.ok = nil
}
func (s *State) GetWarning() error {
	return s.warning
}

func (s *State) SetOk(format string, args ...interface{}) {
	s.error = nil
	s.warning = nil
	s.ok = errors.New(fmt.Sprintf(format, args...))
}
func (s *State) GetOk() error {
	return s.ok
}

func (s *State) GetState() error {
	if s.ok != nil {
		// fmt.Println(s.warning)
		return nil
	}
	if s.warning != nil {
		// fmt.Println(s.warning)
		return nil
	}
	return s.error
}

// ******************************************************************************** //

type Ux struct {
}
type typeColours struct {
	Ref           aurora.Aurora
	Defined       bool
	Name          string
	EnableColours bool
	// TemplateRef   *template.Template
	// TemplateFuncs template.FuncMap
	Prefix string
}

var colours typeColours

func Open(name string, enable bool) (*typeColours, error) {
	var err error

	for range Only.Once {
		if name == "" {
			// name = defaults.BinaryVersion
			name = "Unify"
		}
		name += ": "

		colours.Ref = aurora.NewAurora(enable)
		colours.Name = name
		colours.EnableColours = enable
		colours.Defined = true
		colours.Prefix = fmt.Sprintf("%s", aurora.BrightCyan(colours.Name).Bold())

		//err = termui.Init();
		//if err != nil {
		//      fmt.Printf("failed to initialize termui: %v", err)
		//      break
		//}

		// err = CreateTemplate()
	}

	return &colours, err
}

func Close() {
	if colours.Defined {
		//termui.Close()
	}
}

// func (u *Ux) PrintflnBlue(format string, args ...interface{}) {
// 	log.Printf(format, args...)
// }

func (u *Ux) PrintfWhite(format string, args ...interface{}) {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightWhite(inline))
	}
	fmt.Printf(inline)
}

func (u *Ux) PrintfCyan(format string, args ...interface{}) {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightCyan(inline))
	}
	fmt.Printf(inline)
}

func (u *Ux) PrintfYellow(format string, args ...interface{}) {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightYellow(inline))
	}
	fmt.Printf(inline)
}

func (u *Ux) PrintfRed(format string, args ...interface{}) {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightRed(inline))
	}
	fmt.Printf(inline)
}

func (u *Ux) PrintfGreen(format string, args ...interface{}) {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightGreen(inline))
	}
	fmt.Printf(inline)
}

func (u *Ux) PrintfBlue(format string, args ...interface{}) {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightBlue(inline))
	}
	fmt.Printf(inline)
}

func (u *Ux) PrintfMagenta(format string, args ...interface{}) {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightMagenta(inline))
	}
	fmt.Printf(inline)
}

func (u *Ux) PrintflnWhite(format string, args ...interface{}) {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightWhite(inline))
	}
	fmt.Printf(inline + "\n")
}

func (u *Ux) PrintflnCyan(format string, args ...interface{}) {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightCyan(inline))
	}
	fmt.Printf(inline + "\n")
}

func (u *Ux) PrintflnYellow(format string, args ...interface{}) {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightYellow(inline))
	}
	fmt.Printf(inline + "\n")
}

func (u *Ux) PrintflnRed(format string, args ...interface{}) {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightRed(inline))
	}
	fmt.Printf(inline + "\n")
}

func (u *Ux) PrintflnGreen(format string, args ...interface{}) {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightGreen(inline))
	}
	fmt.Printf(inline + "\n")
}

func (u *Ux) PrintflnBlue(format string, args ...interface{}) {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightBlue(inline))
	}
	fmt.Printf(inline + "\n")
}

func (u *Ux) PrintflnMagenta(format string, args ...interface{}) {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightMagenta(inline))
	}
	fmt.Printf(inline + "\n")
}

func SprintfWhite(format string, args ...interface{}) string {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightWhite(inline))
	}
	return inline
}

func SprintfCyan(format string, args ...interface{}) string {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightCyan(inline))
	}
	return inline
}

func SprintfYellow(format string, args ...interface{}) string {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightYellow(inline))
	}
	return inline
}

func SprintfRed(format string, args ...interface{}) string {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightRed(inline))
	}
	return inline
}

func SprintfGreen(format string, args ...interface{}) string {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightGreen(inline))
	}
	return inline
}

func SprintfBlue(format string, args ...interface{}) string {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightBlue(inline))
	}
	return inline
}

func SprintfMagenta(format string, args ...interface{}) string {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightMagenta(inline))
	}
	return inline
}

func Sprintf(format string, args ...interface{}) string {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s%s", colours.Prefix, inline)
	}
	return inline
}
func (u *Ux) Printf(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, Sprintf(format, args...))
}

func SprintfNormal(format string, args ...interface{}) string {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = Sprintf("%s", aurora.BrightBlue(inline))
	}
	return inline
}

func (u *Ux) PrintfNormal(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, SprintfNormal(format, args...))
}

func (u *Ux) PrintflnNormal(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, SprintfNormal(format+"\n", args...))
}

func SprintfInfo(format string, args ...interface{}) string {
	inline := fmt.Sprintf(format, args...)
	return Sprintf("%s", aurora.BrightBlue(inline))
}

func (u *Ux) PrintfInfo(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, SprintfInfo(format, args...))
}

func (u *Ux) PrintflnInfo(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, SprintfInfo(format+"\n", args...))
}

func SprintfOk(format string, args ...interface{}) string {
	inline := fmt.Sprintf(format, args...)
	return Sprintf("%s", aurora.BrightGreen(inline))
}

func (u *Ux) PrintfOk(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, SprintfOk(format, args...))
}

func (u *Ux) PrintflnOk(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, SprintfOk(format+"\n", args...))
}

func SprintfDebug(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

func (u *Ux) PrintfDebug(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, fmt.Sprintf(format+"\n", args...))
}

func SprintfWarning(format string, args ...interface{}) string {
	inline := fmt.Sprintf(format, args...)
	return Sprintf("%s", aurora.BrightYellow(inline))
}

func (u *Ux) PrintfWarning(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, SprintfWarning(format, args...))
}

func (u *Ux) PrintflnWarning(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, SprintfWarning(format+"\n", args...))
}

func SprintfError(format string, args ...interface{}) string {
	inline := fmt.Sprintf(format, args...)
	return Sprintf("%s", aurora.BrightRed(inline))
}

func (u *Ux) PrintfError(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, SprintfError(format, args...))
}

func (u *Ux) PrintflnError(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, SprintfError(format+"\n", args...))
}

func SprintError(err error) string {
	var s string

	for range Only.Once {
		if err == nil {
			break
		}

		s = Sprintf("%s%s\n", aurora.BrightRed("ERROR: ").Framed(), aurora.BrightRed(err).Framed().SlowBlink().BgBrightWhite())
	}

	return s
}

func (u *Ux) PrintError(err error) {
	_, _ = fmt.Fprintf(os.Stderr, SprintError(err))
}
