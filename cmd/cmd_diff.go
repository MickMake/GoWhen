package cmd

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdHelp"
	"fmt"
	"github.com/spf13/cobra"
)


//goland:noinspection GoNameStartsWithPackageName
type CmdDiff CmdDefault


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
			Use:                   "diff <format> <date/time>",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Diff"},
			Short:                 fmt.Sprintf("Diff date or time."),
			Long:                  fmt.Sprintf("Diff date or time."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRun:                func(cmd *cobra.Command, args []string) { cmds.Data.SetDateIfNil() },
			RunE:                  cmds.CmdDiffFormat,
			Args:                  cobra.MinimumNArgs(2),
		}
		cmd.AddCommand(w.SelfCmd)
		w.SelfCmd.Example = cmdHelp.PrintExamples(w.SelfCmd, ". \"Sat 01 Jul 1967 09:42:42 AEST\"", ". now", ". today", "UnixDate \"Sat Jul  1 09:42:42 UTC 1967\"")

	}

	return w.SelfCmd
}

func (cs *Cmds) CmdDiffFormat(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		var arg []string
		arg, args = cs.PopArgs(2, args)
		// ######################################## //


		cs.Error = cs.Data.DateDiff(arg[0], arg[1])
		if cs.Error != nil {
			break
		}


		// ######################################## //
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}
