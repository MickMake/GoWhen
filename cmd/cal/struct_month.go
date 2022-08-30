package cal

import (
	"GoWhen/Unify/Only"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"reflect"
)


// Month have Week data to represent month.
type Month []Week

func (m *Month) Print() {
	x := (*m)[1][1]
	fmt.Printf("| %s\n", x.Format("Jan 2006"))
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"S", "M", "T", "W", "T", "F", "S"})
	table.SetBorder(true)
	for _, month := range *m {
		var row []string
		for _, week := range month {
			row = append(row, fmt.Sprintf("%d", week.Day()))
		}
		table.Append(row)
	}
	table.Render()
}


// CurrentMonth returns Now().Month()
func CurrentMonth() Month {
	return Now().Month()
}

// NextMonth returns Now().NextMonth()
func NextMonth() Month {
	return Now().NextMonth()
}

// PreviousMonth returns Now().PreviousMonth()
func PreviousMonth() Month {
	return Now().PreviousMonth()
}


// Month returns Month regarding current date.
func (c *Calendar) Month() Month {
	var month Month
	for range Only.Once {
		beginningOfMonth := c.Now.BeginningOfMonth()
		endOfMonth := c.Now.EndOfMonth()
		week := New(beginningOfMonth).Week()
		lastWeek := New(endOfMonth).Week()

		for !reflect.DeepEqual(lastWeek, week) {
			month = append(month, week)
			week = week.Next()
		}
		month = append(month, lastWeek)
	}
	return month
}

// NextMonth returns next Month regarding current date.
func (c *Calendar) NextMonth() Month {
	return c.NextCalendar().Month()
}

// PreviousMonth returns previous Month regarding current date.
func (c *Calendar) PreviousMonth() Month {
	return c.PreviousCalendar().Month()
}
