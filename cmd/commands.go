package cmd

import (
	"GoWhen/cmd/cal"
	"GoWhen/defaults"
	"github.com/MickMake/GoUnify/Only"
	"github.com/MickMake/GoUnify/Unify"
	"github.com/spf13/cobra"
)


type Cmds struct {
	Unify    *Unify.Unify
	Google   *CmdGoogle
	Parse    *CmdParse
	Add      *CmdAdd
	Format   *CmdFormat
	Timezone *CmdTimezone
	Round    *CmdRound
	Is       *CmdIs
	Diff     *CmdDiff
	Range    *CmdRange
	Convert  *CmdConvert
	Alias    *CmdAlias

	reparse  bool
	last  bool
	Data  cal.Data
	Error error
}

//goland:noinspection GoNameStartsWithPackageName
type CmdDefault struct {
	Error   error
	cmd     *cobra.Command
	SelfCmd *cobra.Command
}


var cmds Cmds


func init() {
	for range Only.Once {
		cmds.Unify = Unify.New(
			Unify.Options{
				Description:   defaults.Description,
				BinaryName:    defaults.BinaryName,
				BinaryVersion: defaults.BinaryVersion,
				SourceRepo:    defaults.SourceRepo,
				BinaryRepo:    defaults.BinaryRepo,
				EnvPrefix:     defaults.EnvPrefix,
				HelpSummary:   defaults.HelpSummary,
				ReadMe:        defaults.Readme,
				Examples:      defaults.Examples,
			},
			Unify.Flags{},
		)

		cmdRoot := cmds.Unify.GetCmd()

		cmds.Parse = NewCmdParse()
		cmds.Parse.AttachCommand(cmdRoot)

		cmds.Add = NewCmdAdd()
		cmds.Add.AttachCommand(cmdRoot)

		cmds.Format = NewCmdFormat()
		cmds.Format.AttachCommand(cmdRoot)

		cmds.Timezone = NewCmdTimezone()
		cmds.Timezone.AttachCommand(cmdRoot)

		cmds.Round = NewCmdRound()
		cmds.Round.AttachCommand(cmdRoot)

		cmds.Is = NewCmdIs()
		cmds.Is.AttachCommand(cmdRoot)

		cmds.Diff = NewCmdDiff()
		cmds.Diff.AttachCommand(cmdRoot)

		cmds.Range = NewCmdRange()
		cmds.Range.AttachCommand(cmdRoot)

		cmds.Convert = NewCmdConvert()
		cmds.Convert.AttachCommand(cmdRoot)

		cmds.Alias = NewCmdAlias()
		cmds.Alias.AttachCommand(cmdRoot)


		cmds.Data.GoFormat = true
		cmds.Data.CppFormat = false
		cmds.Data.JavaFormat = false


		cmds.AttachFlags(cmdRoot, cmds.Unify.GetViper())

		// cmds.Google = NewCmdGoogle()
		// cmds.Google.AttachCommands(cmdRoot)
	}
}

func Execute() error {
	var err error

	for range Only.Once {
		// Execute adds all child commands to the root command and sets flags appropriately.
		// This is called by main.main(). It only needs to happen once to the rootCmd.
		err = cmds.Unify.Execute()
		if err != nil {
			break
		}
	}

	return err
}
