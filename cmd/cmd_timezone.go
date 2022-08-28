package cmd

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdConfig"
	"GoWhen/Unify/cmdHelp"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"time"
)


//goland:noinspection GoNameStartsWithPackageName
type CmdTimezone struct {
	Error   error
	cmd     *cobra.Command
	SelfCmd *cobra.Command
}


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
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               nil,
			RunE:                  cmds.CmdTimezone,
			Args:                  cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(w.SelfCmd)
		w.SelfCmd.Example = cmdHelp.PrintExamples(w.SelfCmd, "timezone Australia/Sydney", "tz UTC")

	}

	return w.SelfCmd
}

func (cs *Cmds) CmdTimezone(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		args = cmdConfig.FillArray(1, args)
		var arg string
		arg, args = cs.PopArgs(1, args)
		if cs.Data.Date == nil {
			cs.Data.SetDate(time.Now())
		}
		// ######################################## //


		var loc *time.Location
		loc, cs.Error = time.LoadLocation(arg)
		if cs.Error != nil {
			cs.Error = errors.New("unknown timezone '" + arg + "'")
			break
		}

		// fmt.Printf("Location: %s\n", loc.String())
		t := cs.Data.Date.In(loc)
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
