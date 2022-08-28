package cmd

import (
	"GoWhen/Unify"
	"GoWhen/Unify/Only"
	"GoWhen/defaults"
	"time"
)


type Cmds struct {
	Unify    *Unify.Unify
	Google   *CmdGoogle
	Parse    *CmdParse
	Add      *CmdAdd
	Format   *CmdFormat
	Timezone *CmdTimezone
	Round    *CmdRound

	Data  Data
	Error error
}

var cmds Cmds


type Data struct {
	Date *time.Time
	Duration *time.Duration
}

func (d *Data) SetDate(t time.Time) {
	d.Date = &t
	d.Duration = nil
}

func (d *Data) SetDuration(t time.Duration) {
	d.Date = nil
	d.Duration = &t
}


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
				HelpTemplate:  defaults.HelpTemplate,
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
