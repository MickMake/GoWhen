package cmd

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdConfig"
	"GoWhen/Unify/cmdHelp"
	"fmt"
	"github.com/spf13/cobra"
	"time"
)


//goland:noinspection GoNameStartsWithPackageName
type CmdRound struct {
	Error   error
	cmd     *cobra.Command
	SelfCmd *cobra.Command
}


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
			PreRunE:               nil,
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
			PreRunE:               nil,
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
			PreRunE:               nil,
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
		args = cmdConfig.FillArray(1, args)
		var arg []string
		arg, args = cs.PopArgs(1, args)
		if cs.Data.Date == nil {
			cs.Data.SetDate(time.Now())
		}
		// ######################################## //


		var d Duration
		d, cs.Error = ParseDuration(arg[0])
		if cs.Error != nil {
			break
		}

		t := cs.Data.Date.Round(d.Time)
		cs.Data.SetDate(t)


		// ######################################## //
		// if cs.IsLastArg(args) {
		// 	fmt.Printf("%s\n", cs.Data.Date.Format(time.RFC3339Nano))
		// 	break
		// }
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}

func (cs *Cmds) CmdRoundDown(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		args = cmdConfig.FillArray(1, args)
		var arg []string
		arg, args = cs.PopArgs(1, args)
		if cs.Data.Date == nil {
			cs.Data.SetDate(time.Now())
		}
		// ######################################## //


		var d Duration
		d, cs.Error = ParseDuration(arg[0])
		if cs.Error != nil {
			break
		}

		t := cs.Data.Date.Truncate(d.Time)
		cs.Data.SetDate(t)
		cs.Data.Date.ISOWeek()


		// ######################################## //
		// if cs.IsLastArg(args) {
		// 	fmt.Printf("%s\n", cs.Data.Date.Format(time.RFC3339Nano))
		// 	break
		// }
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}
