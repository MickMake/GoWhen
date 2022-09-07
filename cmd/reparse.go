package cmd

import "github.com/MickMake/GoUnify/Only"


/*
Examples:
parse date "Sat 01 Jul 1967 09:42:42 AEST" add "20d" format "2006-01-02T15:04:05"
add -- '-1y 12M -1w +7d -2h 120m -2s +2000ms' format '2006-01-02 15:04:05'
tz "UTC" format '2006-01-02 15:04:05'

*/

const (
	True = "YES"
	False = "NO"
)


func (cs *Cmds) LastPrint() {
	for range Only.Once {
		if cs.last {
			cs.Data.Print()
			break
		}
	}
}
