package cmd

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdHelp"
	"fmt"
	"github.com/spf13/cobra"
)


//goland:noinspection GoNameStartsWithPackageName
type CmdRound CmdDefault


func NewCmdRound() *CmdRound {
	var ret *CmdRound

	for range Only.Once {
		ret = &CmdRound{
			Error:   nil,
			cmd:     nil,
			SelfCmd: nil,
		}
	}

	return ret
}

func (w *CmdRound) AttachCommand(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		w.cmd = cmd

		// ******************************************************************************** //
		w.SelfCmd = &cobra.Command{
			Use:                   "round <up | down> <duration>",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Round"},
			Short:                 fmt.Sprintf("Round up/down date/time."),
			Long:                  fmt.Sprintf("Round up/down date/time."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRun:                func(cmd *cobra.Command, args []string) { cmds.Data.SetDateIfNil() },
			RunE:                  cmds.CmdRound,
			Args:                  cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(w.SelfCmd)
		w.SelfCmd.Example = cmdHelp.PrintExamples(w.SelfCmd, "up 30s", "down 1w")

		// ******************************************************************************** //
		var CmdRoundUp = &cobra.Command{
			Use:                   "up <duration>",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Round"},
			Short:                 fmt.Sprintf("Round up date/time."),
			Long:                  fmt.Sprintf("Round up date/time."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRun:                func(cmd *cobra.Command, args []string) { cmds.Data.SetDateIfNil() },
			RunE:                  cmds.CmdRoundUp,
			Args:                  cobra.MinimumNArgs(1),
		}
		w.SelfCmd.AddCommand(CmdRoundUp)
		CmdRoundUp.Example = cmdHelp.PrintExamples(CmdRoundUp, "30s", "6h")

		// ******************************************************************************** //
		var CmdRoundDown = &cobra.Command{
			Use:                   "down <duration>",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Round"},
			Short:                 fmt.Sprintf("Round down date/time."),
			Long:                  fmt.Sprintf("Round down date/time."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRun:                func(cmd *cobra.Command, args []string) { cmds.Data.SetDateIfNil() },
			RunE:                  cmds.CmdRoundDown,
			Args:                  cobra.MinimumNArgs(1),
		}
		w.SelfCmd.AddCommand(CmdRoundDown)
		CmdRoundDown.Example = cmdHelp.PrintExamples(CmdRoundDown, "2d", "1w")

	}

	return w.SelfCmd
}

func (cs *Cmds) CmdRound(_ *cobra.Command, _ []string) error {
	for range Only.Once {
		_ = cs.Round.SelfCmd.Help()
	}

	return cs.Error
}

func (cs *Cmds) CmdRoundUp(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		var arg string
		arg, args = cs.PopArg(args)
		// ######################################## //


		cs.Error = cs.Data.DateRound(arg)
		if cs.Error != nil {
			break
		}


		// ######################################## //
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}

func (cs *Cmds) CmdRoundDown(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		var arg string
		arg, args = cs.PopArg(args)
		// ######################################## //


		cs.Error = cs.Data.DateTruncate(arg)
		if cs.Error != nil {
			break
		}


		// ######################################## //
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}
