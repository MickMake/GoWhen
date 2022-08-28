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
type CmdFormat struct {
	Error   error
	cmd     *cobra.Command
	SelfCmd *cobra.Command
}


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
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               nil,
			RunE:                  cmds.CmdFormat,
			Args:                  cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(w.SelfCmd)
		w.SelfCmd.Example = cmdHelp.PrintExamples(w.SelfCmd, "format \"2006-01-02T15:04:05\"", "format \"Mon 02 Jan 15:04:05 2006\"")

	}

	return w.SelfCmd
}

func (cs *Cmds) CmdFormat(_ *cobra.Command, args []string) error {
	for range Only.Once {
		args = cmdConfig.FillArray(1, args)
		var arg string
		arg, args = cs.PopArgs(1, args)
		if cs.Data.Date == nil {
			cs.Data.SetDate(time.Now())
		}
		// ######################################## //


		arg = StrToFormat(arg)


		// ######################################## //
		if cs.IsLastArg(args) {
			fmt.Printf("%s\n", cs.Data.Date.Format(arg))
			break
		}
		// cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}
