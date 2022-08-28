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
	Is       *CmdIs

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

func (d *Data) Clear() {
	d.Date = nil
	d.Duration = nil
}

func (d *Data) IsWeekend() bool {
	if d.Date == nil {
		return false
	}
	switch d.Date.Weekday() {
		case time.Sunday:
			return true
		case time.Saturday:
			return true
	}
	return false
}

func (d *Data) IsWeekday() bool {
	return !d.IsWeekend()
}

func (d *Data) IsLeap() bool {
	if d.Date == nil {
		return false
	}
	year := d.Date.Year()
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func (d *Data) IsDST() bool {
	if d.Date == nil {
		return false
	}
	return d.Date.IsDST()
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

		cmds.Is = NewCmdIs()
		cmds.Is.AttachCommand(cmdRoot)

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
