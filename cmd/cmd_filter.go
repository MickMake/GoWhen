package cmd

import (
	"github.com/MickMake/GoUnify/Only"
	"github.com/MickMake/GoUnify/cmdExec"
	"github.com/MickMake/GoUnify/cmdHelp"
	"github.com/spf13/cobra"
)

func AttachFilterCommands(cmd *cobra.Command) {
	for range Only.Once {
		if cmd == nil {
			break
		}

		cmdKeep := &cobra.Command{
			Use:                   "keep <selector>",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Filter"},
			Short:                 "Keep dates matching selector.",
			Long:                  "Keep dates matching selector.",
			DisableFlagParsing:    true,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdKeep,
			Args:                  cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(cmdKeep)
		cmdKeep.Example = cmdHelp.PrintExamples(cmdKeep,
			"weekday",
			"weekend",
			"mon",
			"monday",
		)

		cmdDrop := &cobra.Command{
			Use:                   "drop <selector>",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Filter"},
			Short:                 "Drop dates matching selector.",
			Long:                  "Drop dates matching selector.",
			DisableFlagParsing:    true,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdDrop,
			Args:                  cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(cmdDrop)
		cmdDrop.Example = cmdHelp.PrintExamples(cmdDrop,
			"weekday",
			"weekend",
			"fri",
			"friday",
		)
	}
}

func (cs *Cmds) CmdKeep(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		var arg string
		arg, args = cmdExec.PopArg(args)

		cs.Error = cs.Data.AddKeepSelector(arg)
		if cs.Error != nil {
			break
		}
		cs.last = true

		cs.last, cs.Error = cmdExec.ReparseArgs(cmd, args)
		cs.LastPrint()
	}

	return cs.Error
}

func (cs *Cmds) CmdDrop(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		var arg string
		arg, args = cmdExec.PopArg(args)

		cs.Error = cs.Data.AddDropSelector(arg)
		if cs.Error != nil {
			break
		}
		cs.last = true

		cs.last, cs.Error = cmdExec.ReparseArgs(cmd, args)
		cs.LastPrint()
	}

	return cs.Error
}
