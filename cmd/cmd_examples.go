package cmd

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdHelp"
	"fmt"
	"github.com/spf13/cobra"
)


//goland:noinspection GoNameStartsWithPackageName
type CmdExamples struct {
	Error   error
	cmd     *cobra.Command
	SelfCmd *cobra.Command
}


func NewCmdExamples() *CmdExamples {
	var ret *CmdExamples

	for range Only.Once {
		ret = &CmdExamples{
			Error:   nil,
			cmd:     nil,
			SelfCmd: nil,
		}
	}

	return ret
}

func (w *CmdExamples) AttachCommand(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		w.cmd = cmd

		// ******************************************************************************** //
		w.SelfCmd = &cobra.Command{
			Use:                   "examples <date/time> <format>",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Examples"},
			Short:                 fmt.Sprintf("Examples."),
			Long:                  fmt.Sprintf("Examples."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               nil,
			RunE:                  cmds.CmdExamples,
			Args:                  cobra.MinimumNArgs(0),
		}
		cmd.AddCommand(w.SelfCmd)
		w.SelfCmd.Example = cmdHelp.PrintExamples(w.SelfCmd, "")

	}

	return w.SelfCmd
}

func (cs *Cmds) CmdExamples(_ *cobra.Command, _ []string) error {
	for range Only.Once {
		fmt.Println(Examples)
	}

	return cs.Error
}
