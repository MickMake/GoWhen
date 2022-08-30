// Package Unify - This package contains common functionality that's used across multiple binaries.
// It's an easy way to include some important functionality into every binary.
// Currently, it provides:
// - Cron scheduler.
// - Daemonizing a process.
// - Logging.
// - Version control and self-update.
// - Cobra/Viper integration.
package Unify

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdConfig"
	"GoWhen/Unify/cmdCron"
	"GoWhen/Unify/cmdDaemon"
	"GoWhen/Unify/cmdHelp"
	"GoWhen/Unify/cmdVersion"
	"GoWhen/defaults"
	"errors"
	"fmt"
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"time"
)

// New - Create new Unify instance.
func New(options Options, flags Flags) *Unify {
	var unify Unify

	for range Only.Once {
		unify.Options = options
		unify.Flags = flags

		if unify.Options.EnvPrefix == "" {
			unify.Options.EnvPrefix = cmdVersion.GetEnvPrefix()
		}

		unify.Error = unify.InitCmds()
		if unify.Error != nil {
			break
		}

		unify.Error = unify.InitFlags()
		if unify.Error != nil {
			break
		}
	}

	return &unify
}

// InitCmds -
func (u *Unify) InitCmds() error {

	for range Only.Once {
		// ******************************************************************************** //
		u.Commands.CmdRoot = &cobra.Command{
			Use:              u.Options.BinaryName,
			Short:            fmt.Sprintf("%s - %s", u.Options.BinaryName, u.Options.Description),
			Long:             fmt.Sprintf("%s - %s", u.Options.BinaryName, u.Options.Description),
			RunE:             CmdRoot,
			TraverseChildren: true,
		}
		u.Commands.CmdRoot.Example = cmdHelp.PrintExamples(u.Commands.CmdRoot, "")

		u.Commands.CmdVersion = cmdVersion.New(u.Options.BinaryName, u.Options.BinaryVersion, false)
		u.Commands.CmdVersion.SetBinaryRepo(u.Options.BinaryRepo)
		u.Commands.CmdVersion.SetSourceRepo(u.Options.SourceRepo)

		u.Commands.CmdDaemon = cmdDaemon.New()

		u.Commands.CmdCron = cmdCron.New()

		u.Commands.CmdConfig = cmdConfig.New(u.Options.BinaryName)

		u.Commands.CmdHelp = cmdHelp.New()
		u.Commands.CmdHelp.SetCommand(u.Options.BinaryName)
		u.Commands.CmdHelp.SetExtendedHelpTemplate(u.Options.HelpTemplate)
		u.Commands.CmdHelp.SetEnvPrefix(u.Options.EnvPrefix)
	}

	return u.Error
}

// InitFlags -
func (u *Unify) InitFlags() error {

	for range Only.Once {
		// SelfCmd.PersistentFlags().StringVarP(&Cmd.WebHost, flagWebHost, "", defaultHost, fmt.Sprintf("Web Host."))
		// Cmd.CmdConfig.SetDefault(flagWebHost, defaultHost)
		// SelfCmd.PersistentFlags().StringVarP(&Cmd.WebPort, flagWebPort, "", defaultPort, fmt.Sprintf("Web Port."))
		// Cmd.CmdConfig.SetDefault(flagWebPort, defaultPort)
		// SelfCmd.PersistentFlags().StringVarP(&Cmd.WebUsername, flagWebUsername, "u", defaultUsername, fmt.Sprintf("Web username."))
		// Cmd.CmdConfig.SetDefault(flagWebUsername, defaultUsername)
		// SelfCmd.PersistentFlags().StringVarP(&Cmd.WebPassword, flagWebPassword, "p", defaultPassword, fmt.Sprintf("Web password."))
		// Cmd.CmdConfig.SetDefault(flagWebPassword, defaultPassword)
		// SelfCmd.PersistentFlags().StringVarP(&Cmd.WebPrefix, flagWebPrefix, "", defaultPrefix, fmt.Sprintf("Web password."))
		// Cmd.CmdConfig.SetDefault(flagWebPrefix, defaultPrefix)

		u.Commands.CmdRoot.PersistentFlags().StringVar(&u.Flags.ConfigFile, cmdConfig.ConfigFileFlag, defaultConfig, fmt.Sprintf("%s: config file.", defaults.BinaryName))
		// _ = rootCmd.PersistentFlags().MarkHidden(flagConfigFile)
		u.Commands.CmdRoot.PersistentFlags().BoolVarP(&u.Flags.Debug, flagDebug, "", defaultDebug, fmt.Sprintf("%s: Debug mode.", defaults.BinaryName))
		u.Commands.CmdConfig.SetDefault(flagDebug, false)
		u.Commands.CmdRoot.PersistentFlags().BoolVarP(&u.Flags.Quiet, flagQuiet, "", defaultQuiet, fmt.Sprintf("%s: Silence all messages.", defaults.BinaryName))
		u.Commands.CmdConfig.SetDefault(flagQuiet, false)
		u.Commands.CmdRoot.PersistentFlags().DurationVarP(&u.Flags.Timeout, flagTimeout, "", defaultTimeout, fmt.Sprintf("Web timeout."))
		u.Commands.CmdConfig.SetDefault(flagTimeout, defaultTimeout)

		u.Commands.CmdRoot.PersistentFlags().SortFlags = false
		u.Commands.CmdRoot.Flags().SortFlags = false

		// cobra.OnInitialize(initConfig)	// Bound to rootCmd now.
		cobra.EnableCommandSorting = false
	}

	return u.Error
}

// Execute -
func (u *Unify) Execute() error {
	var err error
	for range Only.Once {
		u.Commands.CmdVersion.AttachCommands(u.Commands.CmdRoot, true)
		u.Commands.CmdDaemon.AttachCommands(u.Commands.CmdRoot)
		u.Commands.CmdCron.AttachCommands(u.Commands.CmdRoot)
		u.Commands.CmdConfig.AttachCommands(u.Commands.CmdRoot)
		u.Commands.CmdHelp.AttachCommands(u.Commands.CmdRoot)
		u.Commands.CmdConfig.SetDir(u.Flags.ConfigDir)
		u.Commands.CmdConfig.SetFile(u.Flags.ConfigFile)
		u.Commands.CmdRoot.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
			// You can bind cobra and viper in a few locations, but PersistencePreRunE on the root command works well
			return u.Commands.CmdConfig.Init(cmd)
		}

		cc.Init(&cc.Config{
			RootCmd:         u.Commands.CmdRoot,
			Headings:        cc.HiCyan + cc.Bold + cc.Underline,
			Commands:        cc.HiYellow + cc.Bold,
			Example:         cc.Italic,
			ExecName:        cc.Bold,
			Flags:           cc.Bold,
			NoExtraNewlines: true,
			NoBottomNewline: true,
			CmdShortDescr:   0,
			FlagsDataType:   0,
			FlagsDescr:      0,
			Aliases:         0,
		})

		err = u.Commands.Execute()
		if err != nil {
			break
		}
	}
	return err
}

// Execute -
func (c *Commands) Execute() error {
	return c.CmdRoot.Execute()
}

// GetCmd -
func (u *Unify) GetCmd() *cobra.Command {
	return u.Commands.CmdRoot
}

// GetViper -
func (u *Unify) GetViper() *viper.Viper {
	return u.Commands.CmdConfig.GetViper()
}

// WriteConfig -
func (u *Unify) WriteConfig() error {
	return u.Commands.CmdConfig.Write()
}

// ReadConfig -
func (u *Unify) ReadConfig() error {
	return u.Commands.CmdConfig.Read()
}

// CmdRoot -
func CmdRoot(cmd *cobra.Command, args []string) error {
	var err error
	for range Only.Once {
		// _ = cmd.Help()
		err = errors.New(fmt.Sprintf("Unknown command string: %v\n", args))
	}
	return err
}

type Unify struct {
	Options  Options  `json:"options"`
	Flags    Flags    `json:"flags"`
	Commands Commands `json:"commands"`

	Error error `json:"-"`
}

type Options struct {
	Description   string `json:"description"`
	BinaryName    string `json:"binary_name"`
	BinaryVersion string `json:"binary_version"`
	SourceRepo    string `json:"source_repo"`
	BinaryRepo    string `json:"binary_repo"`
	EnvPrefix     string `json:"env_prefix"`
	HelpTemplate  string `json:"help_template"`
}

type Flags struct {
	ConfigFile string        `json:"config_file"`
	ConfigDir  string        `json:"config_dir"`
	CacheDir   string        `json:"cache_dir"`
	Quiet      bool          `json:"quiet"`
	Debug      bool          `json:"debug"`
	Timeout    time.Duration `json:"timeout"`
}

type Commands struct {
	CmdRoot    *cobra.Command
	CmdVersion *cmdVersion.Version
	CmdDaemon  *cmdDaemon.Daemon
	CmdCron    *cmdCron.Cron
	CmdConfig  *cmdConfig.Config
	CmdHelp    *cmdHelp.Help
}

// func (c *Commands) IsValid() error {
// 	for range Only.Once {
// 		if !c.Valid {
// 			c.Error = errors.New("args are not valid")
// 			break
// 		}
// 	}
//
// 	return c.Error
// }
//
// func (c *Commands) ProcessArgs(_ *cobra.Command, _ []string) error {
// 	for range Only.Once {
// 		// ca.Args = args
//
// 		c.Valid = true
// 	}
//
// 	return c.Error
// }
