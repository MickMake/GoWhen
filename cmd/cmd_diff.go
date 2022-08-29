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
type CmdDiff struct {
	Error   error
	cmd     *cobra.Command
	SelfCmd *cobra.Command
}


func NewCmdDiff() *CmdDiff {
	var ret *CmdDiff

	for range Only.Once {
		ret = &CmdDiff{
			Error:   nil,
			cmd:     nil,
			SelfCmd: nil,
		}
	}

	return ret
}

func (w *CmdDiff) AttachCommand(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		w.cmd = cmd

		// ******************************************************************************** //
		w.SelfCmd = &cobra.Command{
			Use:                   "diff <date/time> <format>",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Diff"},
			Short:                 fmt.Sprintf("Diff date or time."),
			Long:                  fmt.Sprintf("Diff date or time."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               nil,
			RunE:                  cmds.CmdDiffFormat,
			Args:                  cobra.MinimumNArgs(2),
		}
		cmd.AddCommand(w.SelfCmd)
		w.SelfCmd.Example = cmdHelp.PrintExamples(w.SelfCmd, "\"Sat 01 Jul 1967 09:42:42 AEST\" \"\"", "now \"\"", "today \"\"", "\"Sat Jul  1 09:42:42 UTC 1967\" UnixDate")

	}

	return w.SelfCmd
}

func (cs *Cmds) CmdDiffFormat(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		args = cmdConfig.FillArray(2, args)
		var arg []string
		arg, args = cs.PopArgs(2, args)
		if cs.Data.Date == nil {
			cs.Data.SetDate(time.Now())
		}
		// ######################################## //


		var t time.Time
		t, cs.Error = cs.Data.Parse(arg[1], arg[0])
		if cs.Error != nil {
			break
		}
		d := DateDiff(*cs.Data.Date, t)
		cs.Data.SetDiff(d)


		// ######################################## //
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}

func DateDiff(a, b time.Time) Diff {
	var d Diff

	for range Only.Once {
		var year, month, day, hour, min, sec int

		if a.Location() != b.Location() {
			b = b.In(a.Location())
		}
		if a.After(b) {
			a, b = b, a
		}
		y1, M1, d1 := a.Date()
		y2, M2, d2 := b.Date()

		h1, m1, s1 := a.Clock()
		h2, m2, s2 := b.Clock()

		year = int(y2 - y1)
		month = int(M2 - M1)
		day = int(d2 - d1)
		hour = int(h2 - h1)
		min = int(m2 - m1)
		sec = int(s2 - s1)

		// Normalize negative values
		if sec < 0 {
			sec += 60
			min--
		}
		if min < 0 {
			min += 60
			hour--
		}
		if hour < 0 {
			hour += 24
			day--
		}
		if day < 0 {
			// days in month:
			t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
			day += 32 - t.Day()
			month--
		}
		if month < 0 {
			month += 12
			year--
		}

		d = Diff{
			Year:   year,
			Month:  month,
			Day:    day,
			Hour:   hour,
			Minute: min,
			Second: sec,
		}
	}

	return d
}


// func CalcMonthDiff(now time.Time, then time.Time) int {
// 	var m int
//
// 	for range Only.Once {
// 		diffYears := now.Year() - then.Year()
// 		if diffYears == 0 {
// 			m = int(now.Month() - then.Month())
// 			break
// 		}
//
// 		if diffYears == 1 {
// 			m = monthsTillEndOfYear(then) + int(now.Month())
// 			break
// 		}
//
// 		yearsInMonths := (now.Year() - then.Year() - 1) * 12
// 		m = yearsInMonths + monthsTillEndOfYear(then) + int(now.Month())
// 	}
//
// 	return m
// }
//
// func monthsTillEndOfYear(then time.Time) int {
// 	return int(12 - then.Month())
// }
//
// /*
// Each month has a different length (28/29/30/31).
// Something like this should be exact.
//
// func diffMonths(now time.Time, then time.Time) int {
// 	diffYears := now.Year() - then.Year()
// 	if diffYears == 0 {
// 		return int(now.Month() - then.Month())
// 	}
//
// 	if diffYears == 1 {
// 		return monthsTillEndOfYear(then) + int(now.Month())
// 	}
//
// 	yearsInMonths := (now.Year() - then.Year() - 1) * 12
// 	return yearsInMonths + monthsTillEndOfYear(then) + int(now.Month())
// }
//
// func monthsTillEndOfYear(then time.Time) int {
// 	return int(12 - then.Month())
// }
//
//
// fmt.Println("Days : ", RoundTime(diff.Seconds()/86400))
// fmt.Println("Weeks : ", RoundTime(diff.Seconds()/604800))
// fmt.Println("Months : ", RoundTime(diff.Seconds()/2600640))
// fmt.Println("Years : ", RoundTime(diff.Seconds()/31207680))
//
//  */