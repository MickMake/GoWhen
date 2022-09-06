package cmd

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdHelp"
	"fmt"
	"github.com/spf13/cobra"
)


//goland:noinspection GoNameStartsWithPackageName
type CmdParse CmdDefault

func NewCmdParse() *CmdParse {
	var ret *CmdParse

	for range Only.Once {
		ret = &CmdParse{
			Error:   nil,
			cmd:     nil,
			SelfCmd: nil,
		}
	}

	return ret
}

func (w *CmdParse) AttachCommand(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		w.cmd = cmd

		// ******************************************************************************** //
		w.SelfCmd = &cobra.Command{
			Use:                   "parse <format> <date/time>",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Parse"},
			Short:                 fmt.Sprintf("Parse date or time."),
			Long:                  fmt.Sprintf("Parse date or time."),
			DisableFlagParsing:    true,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdParse,
			Args:                  cobra.MinimumNArgs(2),
		}
		cmd.AddCommand(w.SelfCmd)
		w.SelfCmd.Example = cmdHelp.PrintExamples(w.SelfCmd,
			"2006-01-02T15:04:05 now",
			"RFC3339 .",
			". .",
			"epoch tomorrow",
			"\"2006-01-02 15:04:05\" yesterday",
			"UnixDate \"Sat Jul  1 09:42:42 UTC 1967\"",
			". \"1967-07-01\"",
			"week last-week",
			"'%F %T' '1967-07-01 09:00:00'",
			"'yyyy-MM-dd HH:mm:ss' '2022-12-31 09:00:00'\"",
			". 2022-09-08 diff . 2023-07-01",
			)

	}

	return w.SelfCmd
}

func (cs *Cmds) CmdParse(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		var arg []string
		arg, args = cs.PopArgs(2, args)
		// ######################################## //


		cs.Error = cs.Data.DateParse(arg[0], arg[1])
		if cs.Error != nil {
			break
		}


		// ######################################## //
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}
