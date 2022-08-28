package cmd

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdConfig"
	"GoWhen/Unify/cmdHelp"
	"fmt"
	"github.com/spf13/cobra"
	"time"
)


//goland:noinspection GoNameStartsWithPackageName
type CmdParse struct {
	Error   error
	cmd     *cobra.Command
	SelfCmd *cobra.Command
}


func NewCmdParse() *CmdParse {
	var ret *CmdParse

	for range Only.Once {
		ret = &CmdParse{
			Error:   nil,
			cmd:     nil,
			SelfCmd: nil,
		}
	}

	return ret
}

func (w *CmdParse) AttachCommand(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		w.cmd = cmd

		// ******************************************************************************** //
		w.SelfCmd = &cobra.Command{
			Use:                   "parse",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Parse"},
			Short:                 fmt.Sprintf("Parse date or time."),
			Long:                  fmt.Sprintf("Parse date or time."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE: func(cmd *cobra.Command, args []string) error {
				// err := cmds.Git.InitArgs(cmd, args)
				// if err != nil {
				// 	return err
				// }
				return nil
			},
			RunE: func(cmd *cobra.Command, args []string) error {
				args = cmdConfig.FillArray(2, args)
				return cmds.CmdParse(args...)
			},
			Args: cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(w.SelfCmd)
		w.SelfCmd.Example = cmdHelp.PrintExamples(w.SelfCmd, "date \"Sat 31 Jul 1967 09:42:42 AEST\"", "date now", "date today")

		var CmdParseDate = &cobra.Command{
			Use:                   "date <date/time>",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Parse"},
			Short:                 fmt.Sprintf("Parse a date."),
			Long:                  fmt.Sprintf("Parse a date."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               nil,
			RunE:                  cmds.CmdParseDate,
			Args:                  cobra.MinimumNArgs(1),
		}
		w.SelfCmd.AddCommand(CmdParseDate)
		CmdParseDate.Example = cmdHelp.PrintExamples(CmdParseDate, "\"Sat 31 Jul 1967 09:42:42 AEST\"", "now", "today")

		// ******************************************************************************** //
		var CmdParseFormat = &cobra.Command{
			Use:                   "format <date/time> <format>",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Parse"},
			Short:                 fmt.Sprintf("Parse a date with custom format."),
			Long:                  fmt.Sprintf("Parse a date with custom format."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               nil,
			RunE:                  cmds.CmdParseFormat,
			Args:                  cobra.MinimumNArgs(2),
		}
		w.SelfCmd.AddCommand(CmdParseFormat)
		CmdParseFormat.Example = cmdHelp.PrintExamples(CmdParseFormat, "\"1967-07-01 09:42:42\" \"2006-01-02 15:04:05\"", "\"1967-07-01 09:42:42\" epoch")

		// ******************************************************************************** //
		var CmdParseEpoch = &cobra.Command{
			Use:                   "epoch <epoch>",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Parse"},
			Short:                 fmt.Sprintf("Parse a date as epoch."),
			Long:                  fmt.Sprintf("Parse a date as epoch."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               nil,
			RunE:                  cmds.CmdParseEpoch,
			Args:                  cobra.MinimumNArgs(1),
		}
		w.SelfCmd.AddCommand(CmdParseEpoch)
		CmdParseEpoch.Example = cmdHelp.PrintExamples(CmdParseEpoch, "1661585565")

	}

	return w.SelfCmd
}

func (cs *Cmds) CmdParse(_ ...string) error {
	for range Only.Once {
		_ = cs.Parse.SelfCmd.Help()
	}

	return cs.Error
}

func (cs *Cmds) CmdParseDate(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		args = cmdConfig.FillArray(1, args)
		var arg string
		arg, args = cs.PopArgs(1, args)
		if cs.Data.Date == nil {
			cs.Data.SetDate(time.Now())
		}
		// ######################################## //


		switch arg {
			case "":
				fallthrough
			case "now":
				fallthrough
			case "today":
				arg = cs.Data.Date.Format("Mon 02 Jan 2006 15:04:05 MST")
		}

		for _, f := range TimeFormats {
			var t time.Time
			t, cs.Error = time.Parse(f, arg)
			if cs.Error == nil {
				cs.Data.SetDate(t)
				break
			}
		}
		if cs.Error != nil {
			break
		}


		// ######################################## //
		// if cs.IsLastArg(args) {
		// 	fmt.Printf("%s\n", cs.Data.Date.Format(time.RFC3339Nano))
		// 	break
		// }
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}

func (cs *Cmds) CmdParseFormat(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		args = cmdConfig.FillArray(2, args)
		var arg1 string
		arg1, args = cs.PopArgs(1, args)
		var arg2 string
		arg2, args = cs.PopArgs(1, args)
		if cs.Data.Date == nil {
			cs.Data.SetDate(time.Now())
		}
		// ######################################## //


		arg2 = StrToFormat(arg2)
		var t time.Time
		t, cs.Error = time.Parse(arg2, arg1)
		if cs.Error == nil {
			fmt.Printf("%s\n", t.Format(time.RFC3339Nano))
			break
		}


		// ######################################## //
		// if cs.IsLastArg(args) {
		// 	fmt.Printf("%s\n", cs.Data.Date.Format(time.RFC3339Nano))
		// 	break
		// }
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}

func (cs *Cmds) CmdParseEpoch(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		args = cmdConfig.FillArray(1, args)
		var arg string
		arg, args = cs.PopArgs(1, args)
		if cs.Data.Date == nil {
			cs.Data.SetDate(time.Now())
		}
		// ######################################## //


		var t time.Time
		for _, f := range TimeFormats {
			// fmt.Printf("Trying '%s'\n", f)
			t, cs.Error = time.Parse(f, arg)
			if cs.Error == nil {
				fmt.Printf("%s\n", t.Format(time.RFC3339Nano))
				break
			}
		}


		// ######################################## //
		// if cs.IsLastArg(args) {
		// 	fmt.Printf("%s\n", cs.Data.Date.Format(time.RFC3339Nano))
		// 	break
		// }
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}
