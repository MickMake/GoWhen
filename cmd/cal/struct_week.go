package cal

import (
	"GoWhen/Unify/Only"
	"fmt"
	"github.com/jinzhu/now"
	"github.com/olekukonko/tablewriter"
	"os"
	"time"
)


// Week have time.Time data to represent week.
type Week []time.Time

func (w *Week) Print() {
	x := (*w)[1]
	fmt.Printf("| %s\n", x.Format("Jan 2006"))
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"S", "M", "T", "W", "T", "F", "S"})
	table.SetBorder(true)
	var row []string
	for _, week := range *w {
		row = append(row, fmt.Sprintf("%d", week.Day()))
	}
	table.Append(row)
	table.Render()
}


// CurrentWeek returns Now().Week()
func CurrentWeek() Week {
	return Now().Week()
}

// NextWeek returns Now().NextWeek()
func NextWeek() Week {
	return Now().NextWeek()
}

// PreviousWeek returns Now().PreviousWeek()
func PreviousWeek() Week {
	return Now().PreviousWeek()
}


// Week returns Week regarding current date.
func (c *Calendar) Week() Week {
	var week Week
	for range Only.Once {
		beginningOfWeek := c.Now.BeginningOfWeek()

		for i := 0; i < 7; i++ {
			week = append(week, beginningOfWeek)
			beginningOfWeek = beginningOfWeek.AddDate(0, 0, 1)
		}
	}
	return week
}

// NextWeek returns next Week regarding current date.
// It doesn't have side effect.
func (c *Calendar) NextWeek() Week {
	var week Week
	for range Only.Once {
		newDate := c.Now.AddDate(0, 0, 7)
		c.Now = now.New(newDate)
		//goland:noinspection GoDeferInLoop
		defer func() {
			c.Now = now.New(newDate.AddDate(0, 0, -7))
		}()

		week = c.Week()
	}
	return week
}

// PreviousWeek returns previous Week regarding current date.
// It doesn't have side effect.
func (c *Calendar) PreviousWeek() Week {
	var week Week
	for range Only.Once {
		newDate := c.Now.AddDate(0, 0, -7)
		c.Now = now.New(newDate)
		//goland:noinspection GoDeferInLoop
		defer func() {
			c.Now = now.New(newDate.AddDate(0, 0, 7))
		}()

		week = c.Week()
	}
	return week
}
