package cmd

import (
	"fmt"
	"time"

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

		if cs.Data.Format == "epoch" {
			fmt.Printf("%d\n", t.Unix())
			break
		}

		if cs.Data.Format == "week" {
			_, w := t.ISOWeek()
			fmt.Printf("%d\n", w)
			break
		}

		if cs.Data.Format == "list" || cs.Data.Format == "cal-week" {
			m := calNew(t).Week()
			m.Print()
			break
		}

		if cs.Data.Format == "cal-month" {
			m := calNew(t).Month()
			m.Print()
			break
		}

		if cs.Data.Format == "cal-year" {
			y := calNew(t).Year()
			y.Print()
			break
		}

		if cs.Data.Format == "" {
			cs.Data.Format = time.RFC3339Nano
		}
		fmt.Printf("%s\n", t.Format(cs.Data.Format))
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
					fmt.Println(t.Format(cs.Data.Format))
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
				fmt.Println(t.Format(cs.Data.Format))
			}
			t = t.AddDate(int(cs.Data.Range.Years), int(cs.Data.Range.Months), 0).Add(cs.Data.Range.Time)
			if lt == t {
				break
			}
			lt = t
		}
	}
}
