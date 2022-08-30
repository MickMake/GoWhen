package cmd

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdConfig"
	"GoWhen/Unify/cmdCron"
	"github.com/spf13/cobra"
)


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


func (cs *Cmds) ReparseArgs(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		// args = args[cull:]
		if (len(args) == 0) || (cs.last) {
			cs.Data.Print()
			// switch {
			// 	case cs.Data.Date != nil:
			// 		if cs.Data.format == "epoch" {
			// 			break
			// 		}
			// 		if cs.Data.format == "" {
			// 			cs.Data.format = time.RFC3339Nano
			// 		}
			// 		fmt.Printf("%s\n", cs.Data.Date.Format(cs.Data.format))
			// 	case cs.Data.Duration != nil:
			// 		fmt.Printf("%s\n", cs.Data.Duration.String())
			// }
			break
		}

		cmdCron.ResetArgs(args...)

		rootCmd := cmdCron.FindRoot(cmd)
		// rootCmd.SetArgs(os.Args)
		cs.Error = rootCmd.Execute()
		if cs.Error != nil {
			break
		}
	}

	return cs.Error
}

func (cs *Cmds) PopArg(args []string) (string, []string) {
	if len(args) == 0 {
		return "", args
	}
	return (args)[0], (args)[1:]
}

func (cs *Cmds) PopArgs(cull int, args []string) ([]string, []string) {
	if cull > len(args) {
		args = cmdConfig.FillArray(cull, args)
		return args, []string{}
	}
	if len(args) == 0 {
		return []string{}, args
	}
	return (args)[:cull], (args)[cull:]
}

func (cs *Cmds) IsLastArg(args []string) bool {
	if len(args) == 0 {
		return true
	}
	return false
}

// func (cs *Cmds) PopArgs(cull int, args *[]string) string {
// 	if cull > (len(*args)-1) {
// 		return ""
// 	}
// 	if len(*args) == 0 {
// 		return ""
// 	}
// 	pop := (*args)[0]
// 	foo := (*args)[cull:]
// 	args = &foo
// 	return pop
// }
