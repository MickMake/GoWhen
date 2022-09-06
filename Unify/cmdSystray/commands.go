package cmdSystray

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdHelp"
	"github.com/spf13/cobra"
)


const Group = "Systray"

func (c *Config) AttachCommands(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		c.cmd = cmd

		// ******************************************************************************** //
		c.SelfCmd = &cobra.Command{
			Use:                   "systray",
			Aliases:               []string{},
			Short:                 "Run program as systray.",
			Long:                  "Run program as systray.",
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               c.InitArgs,
			RunE:                  c.CmdSystray,
			Args:                  cobra.RangeArgs(0, 1),
		}
		cmd.AddCommand(c.SelfCmd)
		c.SelfCmd.Example = cmdHelp.PrintExamples(c.SelfCmd, "run")
		c.SelfCmd.Annotations = map[string]string{"group": Group}

		// ******************************************************************************** //
		var cmdSystrayRun = &cobra.Command{
			Use:                   "run",
			Aliases:               []string{},
			Short:                 "Run program as systray.",
			Long:                  "Run program as systray.",
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               c.InitArgs,
			RunE:                  c.CmdSystrayRun,
			Args:                  cobra.RangeArgs(0, 1),
		}
		c.SelfCmd.AddCommand(cmdSystrayRun)
		cmdSystrayRun.Example = cmdHelp.PrintExamples(cmdSystrayRun, "")
		cmdSystrayRun.Annotations = map[string]string{"group": Group}
	}

	return c.SelfCmd
}

func (c *Config) InitArgs(_ *cobra.Command, _ []string) error {
	var err error
	for range Only.Once {
		//
	}
	return err
}

func (c *Config) CmdSystray(cmd *cobra.Command, _ []string) error {
	for range Only.Once {
		c.Error = cmd.Help()
	}

	return c.Error
}

func (c *Config) CmdSystrayRun(_ *cobra.Command, _ []string) error {
	for range Only.Once {
		c.Error = c.Run()
		if c.Error != nil {
			break
		}
	}

	return c.Error
}
