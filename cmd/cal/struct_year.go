package cal

import (
	"github.com/MickMake/GoUnify/Only"
	"fmt"
	"github.com/jinzhu/now"
)


// Year have Month data to represent year.
type Year [12]Month

func (y *Year) Print() {
	x := (*y)[1][1][1]
	fmt.Printf("|-------------- %s --------------|\n", x.Format("2006"))
	for _, m := range *y {
		m.Print()
		fmt.Println()
	}
}


// CurrentYear returns Now().Year()
func CurrentYear() Year {
	return Now().Year()
}

// NextYear returns Now().NextMonth()
func NextYear() Year {
	return Now().NextYear()
}

// PreviousYear returns Now().PreviousMonth()
func PreviousYear() Year {
	return Now().PreviousYear()
}


// Year returns Year regarding current date.
func (c *Calendar) Year() Year {
	var year Year
	for range Only.Once {
		var days [12]*Calendar
		day := c.Now.BeginningOfYear()
		for i := 0; i < 12; i++ {
			days[i] = &Calendar{
				Now: now.New(day),
			}
			day = day.AddDate(0, 1, 0)
		}
		for i, cal := range days {
			year[i] = cal.Month()
		}
	}
	return year
}

// NextYear returns next Month regarding current date.
func (c *Calendar) NextYear() Year {
	return c.NextYearCalendar().Year()
}

// PreviousYear returns previous Month regarding current date.
func (c *Calendar) PreviousYear() Year {
	return c.PreviousYearCalendar().Year()
}
