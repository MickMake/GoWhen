package cmd

import (
	"fmt"
	"time"

	"GoWhen/cmd/cal"
	"github.com/MickMake/GoUnify/Only"
)

func (cs *Cmds) FilteredPrint() {
	for range Only.Once {
		if cs.Data.Range != nil {
			cs.FilteredPrintRange()
			break
		}

		if cs.Data.Diff != nil {
			cs.Data.PrintDiff()
			break
		}

		if cs.Data.ToDate.Time != nil {
			cs.FilteredPrintTime(*cs.Data.ToDate.Time)
			break
		}

		if cs.Data.FromDate.Time != nil {
			cs.FilteredPrintTime(*cs.Data.FromDate.Time)
			break
		}
	}
}

func (cs *Cmds) FilteredPrintTime(t time.Time) {
	for range Only.Once {
		if !cs.Data.ShouldPrint(t) {
			break
		}

		cs.PrintFormattedTime(t)
	}
}

func (cs *Cmds) PrintFormattedTime(t time.Time) {
	for range Only.Once {
		switch cs.Data.Format {
		case "epoch", "unix":
			fmt.Printf("%d\n", t.Unix())
		case "unix-ms":
			fmt.Printf("%d\n", t.UnixNano()/int64(time.Millisecond))
		case "unix-us":
			fmt.Printf("%d\n", t.UnixNano()/int64(time.Microsecond))
		case "unix-ns":
			fmt.Printf("%d\n", t.UnixNano())
		case "week":
			_, w := t.ISOWeek()
			fmt.Printf("%d\n", w)
		case "list", "cal-week":
			m := cal.New(t).Week()
			m.Print()
		case "cal-month":
			m := cal.New(t).Month()
			m.Print()
		case "cal-year":
			y := cal.New(t).Year()
			y.Print()
		case "iso":
			fmt.Printf("%s\n", t.Format(time.RFC3339))
		case "date":
			fmt.Printf("%s\n", t.Format("2006-01-02"))
		case "time":
			fmt.Printf("%s\n", t.Format("15:04:05"))
		case "datetime":
			fmt.Printf("%s\n", t.Format("2006-01-02 15:04:05"))
		default:
			if cs.Data.Format == "" {
				cs.Data.Format = time.RFC3339Nano
			}
			fmt.Printf("%s\n", t.Format(cs.Data.Format))
		}
	}
}

func (cs *Cmds) FilteredPrintRange() {
	for range Only.Once {
		if cs.Data.Range == nil {
			break
		}

		if cs.Data.Format == "" {
			cs.Data.Format = time.RFC3339
		}

		var lt time.Time
		if cs.Data.ToDate.Time.Before(*cs.Data.FromDate.Time) {
			for t := *cs.Data.FromDate.Time; t.After(*cs.Data.ToDate.Time); {
				if cs.Data.ShouldPrint(t) {
					cs.PrintFormattedTime(t)
				}
				t = t.AddDate(-int(cs.Data.Range.Years), -int(cs.Data.Range.Months), 0).Add(-cs.Data.Range.Time)
				if lt == t {
					break
				}
				lt = t
			}
			break
		}

		for t := *cs.Data.FromDate.Time; t.Before(*cs.Data.ToDate.Time); {
			if cs.Data.ShouldPrint(t) {
				cs.PrintFormattedTime(t)
			}
			t = t.AddDate(int(cs.Data.Range.Years), int(cs.Data.Range.Months), 0).Add(cs.Data.Range.Time)
			if lt == t {
				break
			}
			lt = t
		}
	}
}
