package cmd

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdHelp"
	"GoWhen/cmd/cal"
	"fmt"
	"github.com/spf13/cobra"
)


//goland:noinspection GoNameStartsWithPackageName
type CmdFormat CmdDefault


func NewCmdFormat() *CmdFormat {
	var ret *CmdFormat

	for range Only.Once {
		ret = &CmdFormat{
			Error:   nil,
			cmd:     nil,
			SelfCmd: nil,
		}
	}

	return ret
}

func (w *CmdFormat) AttachCommand(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		w.cmd = cmd

		// ******************************************************************************** //
		w.SelfCmd = &cobra.Command{
			Use:                   "format <format>",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Format"},
			Short:                 fmt.Sprintf("Format date or time."),
			Long:                  fmt.Sprintf("Format date or time."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRun:                func(cmd *cobra.Command, args []string) { cmds.Data.SetDateIfNil() },
			RunE:                  cmds.CmdFormat,
			Args:                  cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(w.SelfCmd)
		w.SelfCmd.Example = cmdHelp.PrintExamples(w.SelfCmd, "format \"2006-01-02T15:04:05\"", "format \"Mon 02 Jan 15:04:05 2006\"")

	}

	return w.SelfCmd
}

func (cs *Cmds) CmdFormat(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		var arg string
		arg, args = cs.PopArg(args)
		// ######################################## //


		cs.Data.Format = cal.StrToFormat(arg)
		cs.last = true


		// ######################################## //
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}
