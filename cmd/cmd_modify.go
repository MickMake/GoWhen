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
			DisableFlagParsing:    true, 
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdAdd,
			Args:                  cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(w.SelfCmd)
		w.SelfCmd.Example = cmdHelp.PrintExamples(w.SelfCmd,
			"30s",
			"7w",
			"1m",
			"'-1y 12M -1w +7d -2h 120m -5s'",
			)

	}

	return w.SelfCmd
}

// CmdAdd - Add duration to FromDate and stores in FromDate
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
			DisableFlagParsing:    true,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdTimezone,
			Args:                  cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(w.SelfCmd)
		w.SelfCmd.Example = cmdHelp.PrintExamples(w.SelfCmd,
			"timezone Australia/Sydney",
			"tz UTC",
			)

	}

	return w.SelfCmd
}

// CmdTimezone - Apply a timezone to FromDate
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
			DisableFlagParsing:    true,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdRound,
			Args:                  cobra.MinimumNArgs(2),
		}
		cmd.AddCommand(w.SelfCmd)
		w.SelfCmd.Example = cmdHelp.PrintExamples(w.SelfCmd,
			"up 30s",
			"up '6h 30m'",
			"down 1w",
			"down '1y 6M'",
			)

		// ******************************************************************************** //
		var CmdRoundUp = &cobra.Command{
			Use:                   "up <duration>",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Round"},
			Short:                 fmt.Sprintf("Round up date/time."),
			Long:                  fmt.Sprintf("Round up date/time."),
			DisableFlagParsing:    true,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdRoundUp,
			Args:                  cobra.MinimumNArgs(1),
		}
		w.SelfCmd.AddCommand(CmdRoundUp)
		CmdRoundUp.Example = cmdHelp.PrintExamples(CmdRoundUp,
			"up 30s",
			"up '6h 30m'",
			)

		// ******************************************************************************** //
		var CmdRoundDown = &cobra.Command{
			Use:                   "down <duration>",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Round"},
			Short:                 fmt.Sprintf("Round down date/time."),
			Long:                  fmt.Sprintf("Round down date/time."),
			DisableFlagParsing:    true,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdRoundDown,
			Args:                  cobra.MinimumNArgs(1),
		}
		w.SelfCmd.AddCommand(CmdRoundDown)
		CmdRoundDown.Example = cmdHelp.PrintExamples(CmdRoundDown,
			"down 1w",
			"down '1y 6M'",
			)

	}

	return w.SelfCmd
}

func (cs *Cmds) CmdRound(_ *cobra.Command, _ []string) error {
	for range Only.Once {
		_ = cs.Round.SelfCmd.Help()
	}

	return cs.Error
}

// CmdRoundUp - Round up FromDate
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

// CmdRoundDown - Round down FromDate
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
