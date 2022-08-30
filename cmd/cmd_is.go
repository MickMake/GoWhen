package cmd

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdHelp"
	"fmt"
	"github.com/spf13/cobra"
)


//goland:noinspection GoNameStartsWithPackageName
type CmdIs CmdDefault


func NewCmdIs() *CmdIs {
	var ret *CmdIs

	for range Only.Once {
		ret = &CmdIs{
			Error:   nil,
			cmd:     nil,
			SelfCmd: nil,
		}
	}

	return ret
}

func (w *CmdIs) AttachCommand(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		w.cmd = cmd

		// ******************************************************************************** //
		w.SelfCmd = &cobra.Command{
			Use:                   "is",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Is"},
			Short:                 fmt.Sprintf("Is date or time."),
			Long:                  fmt.Sprintf("Is date or time."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               nil,
			RunE:                  cmds.CmdIs,
			Args:                  cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(w.SelfCmd)
		w.SelfCmd.Example = cmdHelp.PrintExamples(w.SelfCmd, "dst", "leap", "before . \"1967-07-07 09:42:42\"")

		var CmdIsDst = &cobra.Command{
			Use:                   "dst",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Is"},
			Short:                 fmt.Sprintf("Is date a within daylight savins time?"),
			Long:                  fmt.Sprintf("Is date a within daylight savins time?"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRun:                func(cmd *cobra.Command, args []string) { cmds.Data.SetDateIfNil() },
			RunE:                  cmds.CmdIsDst,
			Args:                  cobra.MinimumNArgs(0),
		}
		w.SelfCmd.AddCommand(CmdIsDst)
		CmdIsDst.Example = cmdHelp.PrintExamples(CmdIsDst, "")

		// ******************************************************************************** //
		var CmdIsLeap = &cobra.Command{
			Use:                   "leap",
			Aliases:               []string{"leap-year"},
			Annotations:           map[string]string{"group": "Is"},
			Short:                 fmt.Sprintf("Is date a leap year?"),
			Long:                  fmt.Sprintf("Is date a leap year?"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRun:                func(cmd *cobra.Command, args []string) { cmds.Data.SetDateIfNil() },
			RunE:                  cmds.CmdIsLeap,
			Args:                  cobra.MinimumNArgs(0),
		}
		w.SelfCmd.AddCommand(CmdIsLeap)
		CmdIsLeap.Example = cmdHelp.PrintExamples(CmdIsLeap, "")

		// ******************************************************************************** //
		var CmdIsWeekday = &cobra.Command{
			Use:                   "weekday",
			Aliases:               []string{"workday"},
			Annotations:           map[string]string{"group": "Is"},
			Short:                 fmt.Sprintf("Is date a weekday?"),
			Long:                  fmt.Sprintf("Is date a weekday?"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRun:                func(cmd *cobra.Command, args []string) { cmds.Data.SetDateIfNil() },
			RunE:                  cmds.CmdIsWeekday,
			Args:                  cobra.MinimumNArgs(0),
		}
		w.SelfCmd.AddCommand(CmdIsWeekday)
		CmdIsWeekday.Example = cmdHelp.PrintExamples(CmdIsWeekday, "")

		// ******************************************************************************** //
		var CmdIsWeekend = &cobra.Command{
			Use:                   "weekend",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Is"},
			Short:                 fmt.Sprintf("Is date a weekday?"),
			Long:                  fmt.Sprintf("Is date a weekday?"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRun:                func(cmd *cobra.Command, args []string) { cmds.Data.SetDateIfNil() },
			RunE:                  cmds.CmdIsWeekend,
			Args:                  cobra.MinimumNArgs(0),
		}
		w.SelfCmd.AddCommand(CmdIsWeekend)
		CmdIsWeekend.Example = cmdHelp.PrintExamples(CmdIsWeekend, "")

		// ******************************************************************************** //
		var CmdIsBefore = &cobra.Command{
			Use:                   "before <format> <date/time>",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Is"},
			Short:                 fmt.Sprintf("Is parsed date before specified date?"),
			Long:                  fmt.Sprintf("Is parsed date before specified date?"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRun:                func(cmd *cobra.Command, args []string) { cmds.Data.SetDateIfNil() },
			RunE:                  cmds.CmdIsBefore,
			Args:                  cobra.MinimumNArgs(2),
		}
		w.SelfCmd.AddCommand(CmdIsBefore)
		CmdIsBefore.Example = cmdHelp.PrintExamples(CmdIsBefore, ". \"1967-07-07 09:42:42\"")

		// ******************************************************************************** //
		var CmdIsAfter = &cobra.Command{
			Use:                   "after <format> <date/time>",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Is"},
			Short:                 fmt.Sprintf("Is parsed date after specified date?"),
			Long:                  fmt.Sprintf("Is parsed date after specified date?"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRun:                func(cmd *cobra.Command, args []string) { cmds.Data.SetDateIfNil() },
			RunE:                  cmds.CmdIsAfter,
			Args:                  cobra.MinimumNArgs(2),
		}
		w.SelfCmd.AddCommand(CmdIsAfter)
		CmdIsAfter.Example = cmdHelp.PrintExamples(CmdIsAfter, "UnixDate \"Sat Jul  1 09:42:42 UTC 1967\"")

	}

	return w.SelfCmd
}

func (cs *Cmds) CmdIs(_ *cobra.Command, _ []string) error {
	for range Only.Once {
		_ = cs.Is.SelfCmd.Help()
	}

	return cs.Error
}

func (cs *Cmds) CmdIsDst(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		// ######################################## //


		if cs.Data.IsDateDST() {
			fmt.Println(True)
		} else {
			fmt.Println(False)
		}
		cs.last = true
		cs.Data.Clear()


		// ######################################## //
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}

func (cs *Cmds) CmdIsLeap(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		// ######################################## //


		if cs.Data.IsDateLeap() {
			fmt.Println(True)
		} else {
			fmt.Println(False)
		}
		cs.last = true
		cs.Data.Clear()


		// ######################################## //
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}

func (cs *Cmds) CmdIsWeekday(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		// ######################################## //


		if cs.Data.IsDateWeekday() {
			fmt.Println(True)
		} else {
			fmt.Println(False)
		}
		cs.last = true
		cs.Data.Clear()


		// ######################################## //
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}

func (cs *Cmds) CmdIsWeekend(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		// ######################################## //


		if cs.Data.IsDateWeekend() {
			fmt.Println(True)
		} else {
			fmt.Println(False)
		}
		cs.last = true
		cs.Data.Clear()


		// ######################################## //
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}

func (cs *Cmds) CmdIsBefore(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		var arg []string
		arg, args = cs.PopArgs(2, args)
		// ######################################## //


		if cs.Data.IsDateBefore(arg[0], arg[1]) {
			fmt.Println(True)
		} else {
			fmt.Println(False)
		}
		cs.last = true
		cs.Data.Clear()


		// ######################################## //
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}

func (cs *Cmds) CmdIsAfter(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		var arg []string
		arg, args = cs.PopArgs(2, args)
		// ######################################## //


		if cs.Data.IsDateAfter(arg[0], arg[1]) {
			fmt.Println(True)
		} else {
			fmt.Println(False)
		}
		cs.last = true
		cs.Data.Clear()


		// ######################################## //
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}
