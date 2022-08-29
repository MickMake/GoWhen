package cmd

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdConfig"
	"GoWhen/Unify/cmdHelp"
	"GoWhen/cmd/cal"
	"fmt"
	"github.com/spf13/cobra"
	"time"
)


//goland:noinspection GoNameStartsWithPackageName
type CmdAdd struct {
	Error   error
	cmd     *cobra.Command
	SelfCmd *cobra.Command
}


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
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               nil,
			RunE:                  cmds.CmdAdd,
			Args:                  cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(w.SelfCmd)
		w.SelfCmd.Example = cmdHelp.PrintExamples(w.SelfCmd, "30s", "7w", "1m", "-- '-1y 12M -1w +7d -2h 120m -5s'")

	}

	return w.SelfCmd
}

func (cs *Cmds) CmdAdd(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		args = cmdConfig.FillArray(1, args)
		var arg []string
		arg, args = cs.PopArgs(1, args)
		if cs.Data.Date == nil {
			cs.Data.SetDate(time.Now())
		}
		// ######################################## //


		var d cal.Duration
		d, cs.Error = cal.ParseDuration(arg[0])
		if cs.Error != nil {
			break
		}

		t := cs.Data.Date.AddDate(int(d.Years), int(d.Months), 0)
		t = t.Add(d.Time)
		cs.Data.SetDate(t)


		// ######################################## //
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}
