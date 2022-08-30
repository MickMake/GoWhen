package cmd

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdHelp"
	"fmt"
	"github.com/spf13/cobra"
)


//goland:noinspection GoNameStartsWithPackageName
type CmdAdd CmdDefault


func NewCmdAdd() *CmdAdd {
	var ret *CmdAdd

	for range Only.Once {
		ret = &CmdAdd{
			Error:   nil,
			cmd:     nil,
			SelfCmd: nil,
		}
	}

	return ret
}

func (w *CmdAdd) AttachCommand(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		w.cmd = cmd

		// ******************************************************************************** //
		w.SelfCmd = &cobra.Command{
			Use:                   "add <duration>",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Parse"},
			Short:                 fmt.Sprintf("Add duration to date."),
			Long:                  fmt.Sprintf("Add duration to date."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRun:                func(cmd *cobra.Command, args []string) { cmds.Data.SetDateIfNil() },
			RunE:                  cmds.CmdAdd,
			Args:                  cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(w.SelfCmd)
		w.SelfCmd.Example = cmdHelp.PrintExamples(w.SelfCmd, "30s", "7w", "1m", "-- '-1y 12M -1w +7d -2h 120m -5s'")

	}

	return w.SelfCmd
}

func (cs *Cmds) CmdAdd(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		var arg string
		arg, args = cs.PopArg(args)
		// ######################################## //


		cs.Error = cs.Data.DateAdd(arg)
		if cs.Error != nil {
			break
		}


		// ######################################## //
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}
