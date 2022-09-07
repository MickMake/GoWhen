package cmdHelp

import (
	"GoWhen/Unify/Only"
	"fmt"
	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"syscall"
	"unsafe"
)


const Group = "Help"

const (
	cmdHelpFlags = "flags"
	cmdHelpReadMe = "readme"
	cmdHelpExamples = "examples"
)

func (h *Help) AttachCommands(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		h.cmd = cmd

		// ******************************************************************************** //
		h.SelfCmd = &cobra.Command{
			Use:                   "help",
			Aliases:               []string{},
			Short:                 "Extended help",
			Long:                  "Extended help",
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               h.InitArgs,
			RunE:                  h.CmdHelp,
			Args:                  cobra.MinimumNArgs(0),
		}
		cmd.AddCommand(h.SelfCmd)
		h.SelfCmd.Example = PrintExamples(h.SelfCmd, "")
		h.SelfCmd.Annotations = map[string]string{"group": Group}

		// ******************************************************************************** //
		CmdAllHelp := &cobra.Command{
			Use:                   cmdHelpFlags,
			Aliases:               []string{},
			Short:                 "Flag help",
			Long:                  "Flag help",
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               h.InitArgs,
			RunE:                  h.CmdHelp,
			Args:                  cobra.MinimumNArgs(0),
		}
		h.SelfCmd.AddCommand(CmdAllHelp)
		CmdAllHelp.Example = PrintExamples(CmdAllHelp, "")
		CmdAllHelp.Annotations = map[string]string{"group": Group}

		// ******************************************************************************** //
		CmdReadMe := &cobra.Command{
			Use:                   cmdHelpReadMe,
			Aliases:               []string{},
			Short:                 "README",
			Long:                  "README",
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               h.InitArgs,
			RunE:                  h.CmdHelp,
			Args:                  cobra.RangeArgs(0, 0),
		}
		h.SelfCmd.AddCommand(CmdReadMe)
		CmdReadMe.Example = PrintExamples(CmdReadMe, "")
		CmdReadMe.Annotations = map[string]string{"group": Group}

		// ******************************************************************************** //
		CmdExamples := &cobra.Command{
			Use:                   cmdHelpExamples,
			Aliases:               []string{},
			Short:                 "Extended examples",
			Long:                  "Extended examples",
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               h.InitArgs,
			RunE:                  h.CmdHelp,
			Args:                  cobra.RangeArgs(0, 0),
		}
		h.SelfCmd.AddCommand(CmdExamples)
		CmdExamples.Example = PrintExamples(CmdExamples, "")
		CmdExamples.Annotations = map[string]string{"group": Group}

		h.cmd.SetHelpTemplate(DefaultHelpTemplate)
		h.cmd.SetUsageTemplate(DefaultUsageTemplate)
	}

	return h.SelfCmd
}

func (h *Help) InitArgs(_ *cobra.Command, _ []string) error {
	var err error
	for range Only.Once {
		//
	}
	return err
}

const DHT = `{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}`

const DUT1 = `{{if .HasAvailableSubCommands}}
{{HeadingStyle "Help on Commands:"}}{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad (CommandStyle .Name) (sum .NamePadding 12)}} {{ with (index .Annotations "group") }}{{ . }}	- {{ end }}{{.Short}}{{end}}{{end}}{{end}}`

const DUT2 = `{{if .HasAvailableSubCommands}}
{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  help {{rpad (CommandStyle .Name) (sum .NamePadding 12)}} {{ with (index .Annotations "group") }}{{ . }}	- {{ end }}{{.Short}}{{end}}{{end}}{{end}}

{{if .HasAvailableSubCommands}}
Use "{{ExecStyle .CommandPath}} help [command]" for more information about a command.{{end}}`

func (h *Help) CmdHelp(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		if cmd.Name() == cmdHelpFlags {
			h.PrintConfig(h.cmd)
			break
		}

		if cmd.Name() == cmdHelpReadMe {
			w := getWidth()
			t := isTerminal()
			if !t {
				// Determine max width.
				for _, l := range strings.Split(h.ReadMe, "\n") {
					if len(l) > w {
						w = len(l)
					}
				}
			}

			result := markdown.Render(h.ReadMe, w, 6)
			fmt.Printf("%s", result)
			break
		}

		if cmd.Name() == cmdHelpExamples {
			w := getWidth()
			t := isTerminal()
			if !t {
				// Determine max width.
				for _, l := range strings.Split(h.Examples, "\n") {
					if len(l) > w {
						w = len(l)
					}
				}
			}

			result := markdown.Render(h.Examples, w, 6)
			fmt.Printf("%s", result)
			break
		}

		if len(args) == 0 {
			// h.cmd.SetVersionTemplate("")

			cmd.SetHelpTemplate(DHT)
			h.cmd.SetUsageTemplate(DUT1)
			_ = h.cmd.Help()
			h.cmd.SetUsageTemplate(DUT2)
			_ = cmd.Help()
			break
		}

		var hc *cobra.Command
		hc, _, h.Error = h.cmd.Find(args)
		if h.Error != nil {
			fmt.Printf("ERROR: %s\n", h.Error)
			h.Error = nil

			cmd.SetHelpTemplate(DHT)
			h.cmd.SetUsageTemplate(DUT1)
			_ = h.cmd.Help()
			h.cmd.SetUsageTemplate(DUT2)
			_ = cmd.Help()
			break
		}
		_ = hc.Help()
	}

	return h.Error
}


type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func getWidth() int {
	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		panic(errno)
	}
	return int(ws.Col)
}

func isTerminal() bool {
	var yes bool
	for range Only.Once {
		o, _ := os.Stdout.Stat()
		if (o.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
			// Is terminal
			yes = true
			break
		}
		// Is pipe
		yes = false
	}
	return yes
}

func isPipe() bool {
	return !isTerminal()
}
