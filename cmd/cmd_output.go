package cmd

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdHelp"
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
			DisableFlagParsing:    true,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdFormat,
			Args:                  cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(w.SelfCmd)
		w.SelfCmd.Example = cmdHelp.PrintExamples(w.SelfCmd,
			"format \"2006-01-02T15:04:05\"",
			"format \"Mon 02 Jan 15:04:05 2006\"",
			)

	}

	return w.SelfCmd
}

func (cs *Cmds) CmdFormat(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		var arg string
		arg, args = cs.PopArg(args)
		// ######################################## //


		cs.Data.ConvertFormat(arg)
		cs.last = true


		// ######################################## //
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}


//goland:noinspection GoNameStartsWithPackageName
type CmdDiff CmdDefault

func NewCmdDiff() *CmdDiff {
	var ret *CmdDiff

	for range Only.Once {
		ret = &CmdDiff{
			Error:   nil,
			cmd:     nil,
			SelfCmd: nil,
		}
	}

	return ret
}

func (w *CmdDiff) AttachCommand(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		w.cmd = cmd

		// ******************************************************************************** //
		w.SelfCmd = &cobra.Command{
			Use:                   "diff <format> <date/time>",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Diff"},
			Short:                 fmt.Sprintf("Diff date or time."),
			Long:                  fmt.Sprintf("Diff date or time."),
			DisableFlagParsing:    true, 
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdDiffFormat,
			Args:                  cobra.MinimumNArgs(2),
		}
		cmd.AddCommand(w.SelfCmd)
		w.SelfCmd.Example = cmdHelp.PrintExamples(w.SelfCmd,
			". \"Sat 01 Jul 1967 09:42:42 AEST\"",
			". now",
			". today",
			"UnixDate \"Sat Jul  1 09:42:42 UTC 1967\"",
			)

	}

	return w.SelfCmd
}

func (cs *Cmds) CmdDiffFormat(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		var arg []string
		arg, args = cs.PopArgs(2, args)
		// ######################################## //


		cs.Error = cs.Data.DateDiff(arg[0], arg[1])
		if cs.Error != nil {
			break
		}
		cs.last = true


		// ######################################## //
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}


//goland:noinspection GoNameStartsWithPackageName
type CmdRange CmdDefault

func NewCmdRange() *CmdRange {
	var ret *CmdRange

	for range Only.Once {
		ret = &CmdRange{
			Error:   nil,
			cmd:     nil,
			SelfCmd: nil,
		}
	}

	return ret
}

func (w *CmdRange) AttachCommand(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		w.cmd = cmd

		// ******************************************************************************** //
		w.SelfCmd = &cobra.Command{
			Use:                   "range <format> <to date/time>",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Range"},
			Short:                 fmt.Sprintf("Produce a range of dates."),
			Long:                  fmt.Sprintf("Produce a range of dates."),
			DisableFlagParsing:    true,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdRange,
			Args:                  cobra.MinimumNArgs(2),
		}
		cmd.AddCommand(w.SelfCmd)
		w.SelfCmd.Example = cmdHelp.PrintExamples(w.SelfCmd,
			". tomorrow 5m",
			"\"2006-01-02 - Monday\" \"now\" 1y",
			". \"2022-01-01 00:00:00\" \"1y 2M 3w 4d 5h 6m 7s\"",
			"'%F %T' '1967-07-01 09:00:00'",
			"'yyyy-MM-dd HH:mm:ss' '2022-12-31 09:00:00'\"",
			)

	}

	return w.SelfCmd
}

// CmdRange - Output only module.
func (cs *Cmds) CmdRange(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		var arg []string
		arg, args = cs.PopArgs(3, args)
		// ######################################## //


		cs.Error = cs.Data.DateRange(arg[0], arg[1], arg[2])
		if cs.Error != nil {
			break
		}
		cs.last = true


		// ######################################## //
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}


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
			DisableFlagParsing:    true,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdIs,
			Args:                  cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(w.SelfCmd)
		w.SelfCmd.Example = cmdHelp.PrintExamples(w.SelfCmd,
			"dst",
			"leap",
			"before . \"1967-07-07 09:42:42\"",
			"after '%F %T' '1967-07-01 09:00:00'",
			"before 'yyyy-MM-dd HH:mm:ss' '2022-12-31 09:00:00'\"",
			)

		var CmdIsDst = &cobra.Command{
			Use:                   "dst",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Is"},
			Short:                 fmt.Sprintf("Is date a within daylight savins time?"),
			Long:                  fmt.Sprintf("Is date a within daylight savins time?"),
			DisableFlagParsing:    true,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdIsDst,
			Args:                  cobra.MinimumNArgs(0),
		}
		w.SelfCmd.AddCommand(CmdIsDst)
		CmdIsDst.Example = cmdHelp.PrintExamples(CmdIsDst,
			"",
		)

		// ******************************************************************************** //
		var CmdIsLeap = &cobra.Command{
			Use:                   "leap",
			Aliases:               []string{"leap-year"},
			Annotations:           map[string]string{"group": "Is"},
			Short:                 fmt.Sprintf("Is date a leap year?"),
			Long:                  fmt.Sprintf("Is date a leap year?"),
			DisableFlagParsing:    true,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdIsLeap,
			Args:                  cobra.MinimumNArgs(0),
		}
		w.SelfCmd.AddCommand(CmdIsLeap)
		CmdIsLeap.Example = cmdHelp.PrintExamples(CmdIsLeap,
			"",
		)

		// ******************************************************************************** //
		var CmdIsWeekday = &cobra.Command{
			Use:                   "weekday",
			Aliases:               []string{"workday"},
			Annotations:           map[string]string{"group": "Is"},
			Short:                 fmt.Sprintf("Is date a weekday?"),
			Long:                  fmt.Sprintf("Is date a weekday?"),
			DisableFlagParsing:    true,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdIsWeekday,
			Args:                  cobra.MinimumNArgs(0),
		}
		w.SelfCmd.AddCommand(CmdIsWeekday)
		CmdIsWeekday.Example = cmdHelp.PrintExamples(CmdIsWeekday,
			"",
		)

		// ******************************************************************************** //
		var CmdIsWeekend = &cobra.Command{
			Use:                   "weekend",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Is"},
			Short:                 fmt.Sprintf("Is date a weekday?"),
			Long:                  fmt.Sprintf("Is date a weekday?"),
			DisableFlagParsing:    true,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdIsWeekend,
			Args:                  cobra.MinimumNArgs(0),
		}
		w.SelfCmd.AddCommand(CmdIsWeekend)
		CmdIsWeekend.Example = cmdHelp.PrintExamples(CmdIsWeekend,
			"",
			)

		// ******************************************************************************** //
		var CmdIsBefore = &cobra.Command{
			Use:                   "before <format> <date/time>",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Is"},
			Short:                 fmt.Sprintf("Is parsed date before specified date?"),
			Long:                  fmt.Sprintf("Is parsed date before specified date?"),
			DisableFlagParsing:    true,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdIsBefore,
			Args:                  cobra.MinimumNArgs(2),
		}
		w.SelfCmd.AddCommand(CmdIsBefore)
		CmdIsBefore.Example = cmdHelp.PrintExamples(CmdIsBefore,
			". \"1967-07-07 09:42:42\"",
			"'%F %T' '1967-07-01 09:00:00'",
			"'yyyy-MM-dd HH:mm:ss' '2022-12-31 09:00:00'\"",
			)

		// ******************************************************************************** //
		var CmdIsAfter = &cobra.Command{
			Use:                   "after <format> <date/time>",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Is"},
			Short:                 fmt.Sprintf("Is parsed date after specified date?"),
			Long:                  fmt.Sprintf("Is parsed date after specified date?"),
			DisableFlagParsing:    true,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdIsAfter,
			Args:                  cobra.MinimumNArgs(2),
		}
		w.SelfCmd.AddCommand(CmdIsAfter)
		CmdIsAfter.Example = cmdHelp.PrintExamples(CmdIsAfter,
			"UnixDate \"Sat Jul  1 09:42:42 UTC 1967\"",
			"'%F %T' '1967-07-01 09:00:00'",
			"'yyyy-MM-dd HH:mm:ss' '2022-12-31 09:00:00'\"",
			)

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
