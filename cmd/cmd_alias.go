package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/MickMake/GoUnify/Only"
	"github.com/MickMake/GoUnify/cmdExec"
	"github.com/MickMake/GoUnify/cmdHelp"
	"github.com/spf13/cobra"
)

const aliasFileName = "aliases.json"

type AliasMap map[string][]string

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
			Use:                   "alias <add | del | list> <name> ...",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Alias"},
			Short:                 "Build up command aliases.",
			Long:                  "Build up command aliases.",
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdAlias,
			Args:                  cobra.MinimumNArgs(0),
		}
		cmd.AddCommand(w.SelfCmd)
		w.SelfCmd.Example = cmdHelp.PrintExamples(w.SelfCmd,
			"add christmas parse . 2026-12-25",
			"list",
			"del christmas",
		)

		// ******************************************************************************** //
		var CmdAliasList = &cobra.Command{
			Use:                   "list",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Alias"},
			Short:                 "Show defined aliases.",
			Long:                  "Show defined aliases.",
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
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Alias"},
			Short:                 "Add an alias.",
			Long:                  "Add an alias.",
			DisableFlagParsing:    true,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdAliasAdd,
			Args:                  cobra.MinimumNArgs(2),
		}
		w.SelfCmd.AddCommand(CmdAliasAdd)
		CmdAliasAdd.Example = cmdHelp.PrintExamples(CmdAliasAdd,
			"christmas parse . 2026-12-25",
			"workday drop weekend",
		)

		// ******************************************************************************** //
		var CmdAliasDelete = &cobra.Command{
			Use:                   "del <name>",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Alias"},
			Short:                 "Delete an alias.",
			Long:                  "Delete an alias.",
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdAliasDelete,
			Args:                  cobra.MinimumNArgs(1),
		}
		w.SelfCmd.AddCommand(CmdAliasDelete)
		CmdAliasDelete.Example = cmdHelp.PrintExamples(CmdAliasDelete,
			"christmas",
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
		var nameArgs []string
		nameArgs, args = cmdExec.PopArgs(1, args)
		name := nameArgs[0]
		if len(args) == 0 {
			cs.Error = errors.New("alias command chain can not be empty")
			break
		}

		if cs.AliasCommandExists(name) {
			cs.Error = fmt.Errorf("alias %q conflicts with an existing command", name)
			break
		}

		aliases, err := cs.LoadAliases()
		if err != nil {
			cs.Error = err
			break
		}

		aliases[name] = append([]string{}, args...)
		cs.Error = cs.SaveAliases(aliases)
		if cs.Error != nil {
			break
		}

		cs.Error = cs.AttachRuntimeAlias(cs.Alias.cmd, name, args)
		if cs.Error != nil {
			break
		}

		fmt.Printf("%s = %s\n", name, strings.Join(args, " "))
	}

	return cs.Error
}

func (cs *Cmds) CmdAliasDelete(_ *cobra.Command, args []string) error {
	for range Only.Once {
		var nameArgs []string
		nameArgs, _ = cmdExec.PopArgs(1, args)
		name := nameArgs[0]

		aliases, err := cs.LoadAliases()
		if err != nil {
			cs.Error = err
			break
		}

		if _, ok := aliases[name]; !ok {
			cs.Error = fmt.Errorf("alias %q does not exist", name)
			break
		}

		delete(aliases, name)
		cs.Error = cs.SaveAliases(aliases)
		if cs.Error != nil {
			break
		}

		fmt.Printf("Deleted alias %q\n", name)
	}

	return cs.Error
}

func (cs *Cmds) CmdAliasList(_ *cobra.Command, _ []string) error {
	for range Only.Once {
		aliases, err := cs.LoadAliases()
		if err != nil {
			cs.Error = err
			break
		}

		if len(aliases) == 0 {
			fmt.Println("No aliases defined.")
			break
		}

		names := make([]string, 0, len(aliases))
		for name := range aliases {
			names = append(names, name)
		}
		sort.Strings(names)

		for _, name := range names {
			fmt.Printf("%s = %s\n", name, strings.Join(aliases[name], " "))
		}
	}

	return cs.Error
}

func (cs *Cmds) AttachStoredAliases(root *cobra.Command) error {
	aliases, err := cs.LoadAliases()
	if err != nil {
		return err
	}

	names := make([]string, 0, len(aliases))
	for name := range aliases {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		err = cs.AttachRuntimeAlias(root, name, aliases[name])
		if err != nil {
			return err
		}
	}

	return nil
}

func (cs *Cmds) AttachRuntimeAlias(root *cobra.Command, name string, chain []string) error {
	if root == nil {
		return errors.New("can not attach alias without root command")
	}

	if name == "" {
		return errors.New("alias name can not be empty")
	}

	if len(chain) == 0 {
		return fmt.Errorf("alias %q command chain can not be empty", name)
	}

	if cs.AliasCommandExists(name) {
		return fmt.Errorf("alias %q conflicts with an existing command", name)
	}

	aliasName := name
	aliasChain := append([]string{}, chain...)

	root.AddCommand(&cobra.Command{
		Use:                   aliasName,
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "Alias"},
		Short:                 "Alias for: " + strings.Join(aliasChain, " "),
		Long:                  "Alias for: " + strings.Join(aliasChain, " "),
		DisableFlagParsing:    true,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.InitArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			next := append([]string{}, aliasChain...)
			next = append(next, args...)

			cs.last, cs.Error = cmdExec.ReparseArgs(cmd, next)
			cs.LastPrint()
			return cs.Error
		},
		Args: cobra.ArbitraryArgs,
	})

	return nil
}

func (cs *Cmds) AliasCommandExists(name string) bool {
	if cs.Alias == nil || cs.Alias.cmd == nil {
		return false
	}

	for _, command := range cs.Alias.cmd.Commands() {
		if command.Name() == name || command.HasAlias(name) {
			return true
		}
	}

	return false
}

func (cs *Cmds) AliasFile() string {
	return filepath.Join(cs.Unify.Commands.CmdConfig.Dir, aliasFileName)
}

func (cs *Cmds) LoadAliases() (AliasMap, error) {
	aliases := AliasMap{}
	content, err := os.ReadFile(cs.AliasFile())
	if err != nil {
		if os.IsNotExist(err) {
			return aliases, nil
		}
		return aliases, err
	}

	if len(content) == 0 {
		return aliases, nil
	}

	err = json.Unmarshal(content, &aliases)
	if err != nil {
		return aliases, err
	}

	return aliases, nil
}

func (cs *Cmds) SaveAliases(aliases AliasMap) error {
	content, err := json.MarshalIndent(aliases, "", "  ")
	if err != nil {
		return err
	}
	content = append(content, '\n')

	err = os.MkdirAll(filepath.Dir(cs.AliasFile()), 0700)
	if err != nil {
		return err
	}

	return os.WriteFile(cs.AliasFile(), content, 0600)
}
