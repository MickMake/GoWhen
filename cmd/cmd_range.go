package cmd

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdHelp"
	"GoWhen/cmd/cal"
	"fmt"
	"github.com/spf13/cobra"
)


//goland:noinspection GoNameStartsWithPackageName
type CmdRange CmdDefault


func NewCmdRange() *CmdRange {
	var ret *CmdRange

	for range Only.Once {
		ret = &CmdRange{
			Error:   nil,
			cmd:     nil,
			SelfCmd: nil,
		}
	}

	return ret
}

func (w *CmdRange) AttachCommand(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		w.cmd = cmd

		// ******************************************************************************** //
		w.SelfCmd = &cobra.Command{
			Use:                   "range <format> <to date/time>",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Range"},
			Short:                 fmt.Sprintf("Produce a range of dates."),
			Long:                  fmt.Sprintf("Produce a range of dates."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRun:                func(cmd *cobra.Command, args []string) { cmds.Data.SetDateIfNil() },
			RunE:                  cmds.CmdRange,
			Args:                  cobra.MinimumNArgs(2),
		}
		cmd.AddCommand(w.SelfCmd)
		w.SelfCmd.Example = cmdHelp.PrintExamples(w.SelfCmd, ". \"\" \"\"", "")

	}

	return w.SelfCmd
}

func (cs *Cmds) CmdRange(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		var arg []string
		arg, args = cs.PopArgs(3, args)
		// ######################################## //


		cs.Data.Format = cal.StrToFormat(arg[0])
		cs.Error = cs.Data.DateRange(cs.Data.Format, arg[1], arg[2])
		if cs.Error != nil {
			break
		}


		// ######################################## //
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}
