package cmdConfig

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdHelp"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const Group = "Config"

func (c *Config) AttachCommands(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		c.cmd = cmd

		// ******************************************************************************** //
		c.SelfCmd = &cobra.Command{
			Use:                   "config",
			Short:                 "Create, update or show config file.",
			Long:                  "Create, update or show config file.",
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               c.InitArgs,
			RunE:                  c.CmdConfig,
			Args:                  cobra.RangeArgs(0, 1),
		}
		cmd.AddCommand(c.SelfCmd)
		c.SelfCmd.Example = cmdHelp.PrintExamples(c.SelfCmd, "read", "write", "write --timeout=60s")
		c.SelfCmd.Annotations = map[string]string{"group": Group}

		// ******************************************************************************** //
		var cmdConfigWrite = &cobra.Command{
			Use:                   "write",
			Short:                 "Update config file from CLI args.",
			Long:                  "Update config file from CLI args.",
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               c.InitArgs,
			RunE:                  c.CmdWrite,
			Args:                  cobra.RangeArgs(0, 1),
		}
		c.SelfCmd.AddCommand(cmdConfigWrite)
		cmdConfigWrite.Example = cmdHelp.PrintExamples(cmdConfigWrite, "", "write --timeout=60s", "--debug=true")
		cmdConfigWrite.Annotations = map[string]string{"group": Group}

		// ******************************************************************************** //
		var cmdConfigRead = &cobra.Command{
			Use:                   "read",
			Short:                 "Read config file.",
			Long:                  "Read config file.",
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               c.InitArgs,
			RunE:                  c.CmdRead,
			Args:                  cobra.RangeArgs(0, 1),
		}
		c.SelfCmd.AddCommand(cmdConfigRead)
		cmdConfigRead.Example = cmdHelp.PrintExamples(cmdConfigRead, "")
		cmdConfigRead.Annotations = map[string]string{"group": Group}
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

func (c *Config) CmdConfig(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		_, _ = fmt.Fprintf(os.Stderr, "Using config file '%s'\n", c.viper.ConfigFileUsed())
		if len(args) == 0 {
			_ = cmd.Help()
		}
	}

	return c.Error
}

func (c *Config) CmdWrite(_ *cobra.Command, args []string) error {
	for range Only.Once {
		if len(args) == 1 {
			c.File = args[0]
			c.SetFile(c.File)
			c.Error = c.Open()
			if c.Error != nil {
				break
			}
		}

		c.Error = c.UpdateFlags()
		if c.Error != nil {
			break
		}

		_, _ = fmt.Fprintf(os.Stderr, "Using config file '%s'\n", c.viper.ConfigFileUsed())
		fmt.Println("New config:")
		cmdHelp.PrintConfig(c.cmd, c.EnvPrefix)

		c.Error = c.Write()
		if c.Error != nil {
			break
		}
	}

	return c.Error
}

func (c *Config) CmdRead(_ *cobra.Command, args []string) error {
	for range Only.Once {
		if len(args) == 1 {
			c.File = args[0]
			c.SetFile(c.File)

			c.Error = c.Open()
			if c.Error != nil {
				break
			}
		}

		c.Error = c.UpdateFlags()
		if c.Error != nil {
			break
		}

		_, _ = fmt.Fprintf(os.Stderr, "Using config file '%s'\n", c.viper.ConfigFileUsed())

		cmdHelp.PrintConfig(c.cmd, c.EnvPrefix) // rootCmd
	}

	return c.Error
}
