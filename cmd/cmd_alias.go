package cmd

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdHelp"
	"fmt"
	"github.com/spf13/cobra"
)


//goland:noinspection GoNameStartsWithPackageName
type CmdAlias CmdDefault

func NewCmdAlias() *CmdAlias {
	var ret *CmdAlias

	for range Only.Once {
		ret = &CmdAlias{
			Error:   nil,
			cmd:     nil,
			SelfCmd: nil,
		}
	}

	return ret
}

func (w *CmdAlias) AttachCommand(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		w.cmd = cmd

		// ******************************************************************************** //
		w.SelfCmd = &cobra.Command{
			Use:                   "alias <add | del | *> <name> ...",
			Aliases:               []string{"a"},
			Annotations:           map[string]string{"group": "Alias"},
			Short:                 fmt.Sprintf("Build up command aliases."),
			Long:                  fmt.Sprintf("Build up command aliases."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdAlias,
			Args:                  cobra.MinimumNArgs(0),
		}
		cmd.AddCommand(w.SelfCmd)
		w.SelfCmd.Example = cmdHelp.PrintExamples(w.SelfCmd,
			"add ",
		)

		// ******************************************************************************** //
		var CmdAliasList = &cobra.Command{
			Use:                   "list",
			Aliases:               []string{"ls"},
			Annotations:           map[string]string{"group": "Alias"},
			Short:                 fmt.Sprintf("Show defined aliases."),
			Long:                  fmt.Sprintf("Show defined aliases."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdAliasList,
			Args:                  cobra.MinimumNArgs(0),
		}
		w.SelfCmd.AddCommand(CmdAliasList)
		CmdAliasList.Example = cmdHelp.PrintExamples(CmdAliasList,
			"",
		)

		// ******************************************************************************** //
		var CmdAliasAdd = &cobra.Command{
			Use:                   "add <name> <cmd> ...",
			Aliases:               []string{"create"},
			Annotations:           map[string]string{"group": "Alias"},
			Short:                 fmt.Sprintf("Add an alias."),
			Long:                  fmt.Sprintf("Add an alias."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdAliasAdd,
			Args:                  cobra.MinimumNArgs(2),
		}
		w.SelfCmd.AddCommand(CmdAliasAdd)
		CmdAliasAdd.Example = cmdHelp.PrintExamples(CmdAliasAdd,
			"",
		)

		// ******************************************************************************** //
		var CmdAliasDelete = &cobra.Command{
			Use:                   "del <name>",
			Aliases:               []string{"destroy"},
			Annotations:           map[string]string{"group": "Alias"},
			Short:                 fmt.Sprintf("Delete an alias."),
			Long:                  fmt.Sprintf("Delete an alias."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdAliasDelete,
			Args:                  cobra.MinimumNArgs(1),
		}
		w.SelfCmd.AddCommand(CmdAliasDelete)
		CmdAliasDelete.Example = cmdHelp.PrintExamples(CmdAliasDelete,
			"",
		)

	}

	return w.SelfCmd
}

func (cs *Cmds) CmdAlias(_ *cobra.Command, _ []string) error {
	for range Only.Once {
		cs.Error = cs.Alias.SelfCmd.Help()
	}

	return cs.Error
}

func (cs *Cmds) CmdAliasAdd(_ *cobra.Command, args []string) error {
	for range Only.Once {
		var arg []string
		arg, args = cs.PopArgs(1, args)
		// ######################################## //


		fmt.Println("\n############################################################")
		fmt.Printf("Name: %v\n", arg)
		fmt.Printf("Args: %v\n", args)


		// ######################################## //
		// cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}

func (cs *Cmds) CmdAliasDelete(_ *cobra.Command, args []string) error {
	for range Only.Once {
		var arg []string
		arg, args = cs.PopArgs(1, args)
		// ######################################## //


		fmt.Println("\n############################################################")
		fmt.Printf("Name: %v\n", arg)


		// ######################################## //
		// cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}

func (cs *Cmds) CmdAliasList(_ *cobra.Command, _ []string) error {
	for range Only.Once {
		// var arg []string
		// arg, args = cs.PopArgs(1, args)
		// ######################################## //


		fmt.Println("\n############################################################")


		// ######################################## //
		// cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}
