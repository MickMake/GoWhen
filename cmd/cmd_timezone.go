package cmd

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdHelp"
	"fmt"
	"github.com/spf13/cobra"
)


//goland:noinspection GoNameStartsWithPackageName
type CmdTimezone CmdDefault


func NewCmdTimezone() *CmdTimezone {
	var ret *CmdTimezone

	for range Only.Once {
		ret = &CmdTimezone{
			Error:   nil,
			cmd:     nil,
			SelfCmd: nil,
		}
	}

	return ret
}

func (w *CmdTimezone) AttachCommand(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		w.cmd = cmd

		// ******************************************************************************** //
		w.SelfCmd = &cobra.Command{
			Use:                   "timezone <zone>",
			Aliases:               []string{"tz"},
			Annotations:           map[string]string{"group": "Timezone"},
			Short:                 fmt.Sprintf("Adjust date/time by timezone."),
			Long:                  fmt.Sprintf("Adjust date/time by timezone."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRun:                func(cmd *cobra.Command, args []string) { cmds.Data.SetDateIfNil() },
			RunE:                  cmds.CmdTimezone,
			Args:                  cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(w.SelfCmd)
		w.SelfCmd.Example = cmdHelp.PrintExamples(w.SelfCmd, "timezone Australia/Sydney", "tz UTC")

	}

	return w.SelfCmd
}

func (cs *Cmds) CmdTimezone(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		var arg string
		arg, args = cs.PopArg(args)
		// ######################################## //


		cs.Error = cs.Data.DateTimezone(arg)
		if cs.Error != nil {
			break
		}


		// ######################################## //
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}
