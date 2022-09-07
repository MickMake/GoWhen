package cmdShell

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdHelp"
	"github.com/spf13/cobra"
)


const Group = "Shell"

func (s *Shell) AttachCommands(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		s.cmd = cmd

		// ******************************************************************************** //
		s.SelfCmd = &cobra.Command{
			Use:                   "shell",
			Aliases:               []string{},
			Short:                 "Run commands in a shell.",
			Long:                  "Run commands in a shell.",
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               s.InitArgs,
			RunE:                  s.CmdHelpAll,
			Args:                  cobra.RangeArgs(0, 0),
		}
		cmd.AddCommand(s.SelfCmd)
		s.SelfCmd.Example = cmdHelp.PrintExamples(s.SelfCmd, "")
		s.SelfCmd.Annotations = map[string]string{"group": Group}

	}

	return s.SelfCmd
}

func (s *Shell) InitArgs(_ *cobra.Command, _ []string) error {
	for range Only.Once {
		//
	}
	return s.Error
}

func (s *Shell) CmdHelpAll(_ *cobra.Command, _ []string) error {
	for range Only.Once {
		s.Error = s.RunShell()
	}

	return s.Error
}
