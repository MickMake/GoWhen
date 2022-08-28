package cmd

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdHelp"
	"fmt"
	"github.com/spf13/cobra"
	"time"
)


//goland:noinspection GoNameStartsWithPackageName
type CmdIs struct {
	Error   error
	cmd     *cobra.Command
	SelfCmd *cobra.Command
}


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
		w.SelfCmd.Example = cmdHelp.PrintExamples(w.SelfCmd, "date \"Sat 31 Jul 1967 09:42:42 AEST\"", "date now", "date today")

		var CmdIsDst = &cobra.Command{
			Use:                   "dst",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Is"},
			Short:                 fmt.Sprintf("Is date a within daylight savins time?"),
			Long:                  fmt.Sprintf("Is date a within daylight savins time?"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               nil,
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
			PreRunE:               nil,
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
			PreRunE:               nil,
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
			PreRunE:               nil,
			RunE:                  cmds.CmdIsWeekend,
			Args:                  cobra.MinimumNArgs(0),
		}
		w.SelfCmd.AddCommand(CmdIsWeekend)
		CmdIsWeekend.Example = cmdHelp.PrintExamples(CmdIsWeekend, "")

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
		// args = cmdConfig.FillArray(1, args)
		// var arg string
		// arg, args = cs.PopArgs(1, args)
		if cs.Data.Date == nil {
			cs.Data.SetDate(time.Now())
		}
		// ######################################## //


		if cs.Data.IsDST() {
			fmt.Println(True)
		} else {
			fmt.Println(False)
		}
		cs.Data.Clear()


		// ######################################## //
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}

func (cs *Cmds) CmdIsLeap(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		// args = cmdConfig.FillArray(1, args)
		// var arg string
		// arg, args = cs.PopArgs(1, args)
		if cs.Data.Date == nil {
			cs.Data.SetDate(time.Now())
		}
		// ######################################## //


		if cs.Data.IsLeap() {
			fmt.Println(True)
		} else {
			fmt.Println(False)
		}
		cs.Data.Clear()


		// ######################################## //
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}

func (cs *Cmds) CmdIsWeekday(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		// args = cmdConfig.FillArray(1, args)
		// var arg string
		// arg, args = cs.PopArgs(1, args)
		if cs.Data.Date == nil {
			cs.Data.SetDate(time.Now())
		}
		// ######################################## //


		if cs.Data.IsWeekday() {
			fmt.Println(True)
		} else {
			fmt.Println(False)
		}
		cs.Data.Clear()


		// ######################################## //
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}

func (cs *Cmds) CmdIsWeekend(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		// args = cmdConfig.FillArray(1, args)
		// var arg string
		// arg, args = cs.PopArgs(1, args)
		if cs.Data.Date == nil {
			cs.Data.SetDate(time.Now())
		}
		// ######################################## //


		if cs.Data.IsWeekend() {
			fmt.Println(True)
		} else {
			fmt.Println(False)
		}
		cs.Data.Clear()


		// ######################################## //
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}
