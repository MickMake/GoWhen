package cmd

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdConfig"
	"GoWhen/Unify/cmdHelp"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
	"time"
)


//goland:noinspection GoNameStartsWithPackageName
type CmdAdd struct {
	Error   error
	cmd     *cobra.Command
	SelfCmd *cobra.Command
}


func NewCmdAdd() *CmdAdd {
	var ret *CmdAdd

	for range Only.Once {
		ret = &CmdAdd{
			Error:   nil,
			cmd:     nil,
			SelfCmd: nil,
		}
	}

	return ret
}

func (w *CmdAdd) AttachCommand(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		w.cmd = cmd

		// ******************************************************************************** //
		w.SelfCmd = &cobra.Command{
			Use:                   "add <duration>",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Parse"},
			Short:                 fmt.Sprintf("Add duration to date."),
			Long:                  fmt.Sprintf("Add duration to date."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               nil,
			RunE:                  cmds.CmdAdd,
			Args:                  cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(w.SelfCmd)
		w.SelfCmd.Example = cmdHelp.PrintExamples(w.SelfCmd, "30s", "7w", "1m", "-- '-1y 12M -1w +7d -2h 120m -5s'")

	}

	return w.SelfCmd
}

func (cs *Cmds) CmdAdd(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		args = cmdConfig.FillArray(1, args)
		var arg string
		arg, args = cs.PopArgs(1, args)
		if cs.Data.Date == nil {
			cs.Data.SetDate(time.Now())
		}
		// ######################################## //


		var d Duration
		d, cs.Error = ParseDuration(arg)
		if cs.Error != nil {
			break
		}

		t := cs.Data.Date.AddDate(int(d.Years), int(d.Months), 0)
		t = t.Add(d.Time)
		cs.Data.SetDate(t)


		// ######################################## //
		// if cs.IsLastArg(args) {
		// 	fmt.Printf("%s\n", cs.Data.Date.Format(time.RFC3339Nano))
		// 	break
		// }
		cs.Error = cs.ReparseArgs(cmd, args)
	}

	return cs.Error
}


type Duration struct {
	Time time.Duration
	Years int64
	Months int64
	// Weeks int	// Handled by classic time.Duration
	// Days int	// Handled by classic time.Duration
}

func ParseDuration(s string) (Duration, error) {
	var duration Duration
	var err error

	for range Only.Once {
		times := strings.Split(s, " ")

		for _, ds := range times {
			ds = strings.TrimSpace(ds)
			if ds == "" {
				continue
			}

			var d time.Duration
			d, err = time.ParseDuration(ds)
			if err == nil {
				duration.Time += d
				continue
			}

			//
			// neg := false
			// c := ds[0]
			// if c == '-' || c == '+' {
			// 	neg = c == '-'
			// 	ds = ds[1:]
			// }

			lb := ds[len(ds)-1]

			switch lb {
				case 'Y':
					fallthrough
				case 'y':
					// Using DateAdd type duration.
					var lbv int64
					lbv, err = strconv.ParseInt(ds[:len(ds)-1], 10, 64)
					if err != nil {
						break
					}
					duration.Years += lbv

				case 'M':
					// Using DateAdd type duration.
					var lbv int64
					lbv, err = strconv.ParseInt(ds[:len(ds)-1], 10, 64)
					if err != nil {
						break
					}
					duration.Months += lbv

				case 'W':
					fallthrough
				case 'w':
					// Straight-forward conversion.
					var lbv float64
					lbv, err = strconv.ParseFloat(ds[:len(ds)-1], 10)
					if err != nil {
						break
					}
					v := float64(int64(time.Hour) * 168) * lbv
					duration.Time += time.Duration(v)

				case 'D':
					fallthrough
				case 'd':
					// Straight-forward conversion.
					var lbv float64
					lbv, err = strconv.ParseFloat(ds[:len(ds)-1], 10)
					if err != nil {
						break
					}
					v := float64(int64(time.Hour) * 24) * lbv
					duration.Time += time.Duration(v)

				default:
					err = errors.New("time: invalid duration " + ds)
					break
			}
		}
	}

	// for range Only.Once {
	// 	var fv int64
	// 	for _, v := range duration.Time {
	// 		fv += int64(v)
	// 	}
	// 	duration.Time = []time.Duration{time.Duration(fv)}
	//
	// 	for _, v := range duration.Time {
	// 		fv += int64(v)
	// 	}
	// 	duration.Years = []int64{fv}
	//
	// 	for _, v := range duration.Time {
	// 		fv += int64(v)
	// 	}
	// 	duration.Months = []int64{fv}
	// }

	return duration, err
}
