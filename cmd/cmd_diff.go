package cmd

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdConfig"
	"GoWhen/Unify/cmdHelp"
	"GoWhen/cmd/cal"
	"fmt"
	"github.com/spf13/cobra"
	"time"
)


//goland:noinspection GoNameStartsWithPackageName
type CmdDiff struct {
	Error   error
	cmd     *cobra.Command
	SelfCmd *cobra.Command
}


func NewCmdDiff() *CmdDiff {
	var ret *CmdDiff

	for range Only.Once {
		ret = &CmdDiff{
			Error:   nil,
			cmd:     nil,
			SelfCmd: nil,
		}
	}

	return ret
}

func (w *CmdDiff) AttachCommand(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		w.cmd = cmd

		// ******************************************************************************** //
		w.SelfCmd = &cobra.Command{
			Use:                   "diff <date/time> <format>",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Diff"},
			Short:                 fmt.Sprintf("Diff date or time."),
			Long:                  fmt.Sprintf("Diff date or time."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               nil,
			RunE:                  cmds.CmdDiffFormat,
			Args:                  cobra.MinimumNArgs(2),
		}
		cmd.AddCommand(w.SelfCmd)
		w.SelfCmd.Example = cmdHelp.PrintExamples(w.SelfCmd, "\"Sat 01 Jul 1967 09:42:42 AEST\" \"\"", "now \"\"", "today \"\"", "\"Sat Jul  1 09:42:42 UTC 1967\" UnixDate")

	}

	return w.SelfCmd
}

func (cs *Cmds) CmdDiffFormat(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		args = cmdConfig.FillArray(2, args)
		var arg []string
		arg, args = cs.PopArgs(2, args)
		if cs.Data.Date == nil {
			cs.Data.SetDate(time.Now())
		}
		// ######################################## //


		var t time.Time
		t, cs.Error = cs.Data.Parse(arg[1], arg[0])
		if cs.Error != nil {
			break
		}
		d := cal.DateDiff(*cs.Data.Date, t)
		cs.Data.SetDiff(d)


		// ######################################## //
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}
