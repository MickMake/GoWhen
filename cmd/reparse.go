package cmd

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdCron"
	"fmt"
	"github.com/spf13/cobra"
	"time"
)


func (cs *Cmds) ReparseArgs(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		// args = args[cull:]
		if len(args) == 0 {
			switch {
				case cs.Data.Date != nil:
					fmt.Printf("%s\n", cs.Data.Date.Format(time.RFC3339Nano))
				case cs.Data.Duration != nil:
					fmt.Printf("%s\n", cs.Data.Duration.String())
			}
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

func (cs *Cmds) PopArgs(cull int, args []string) (string, []string) {
	if cull > len(args) {
		return "", args
	}
	if len(args) == 0 {
		return "", args
	}
	return (args)[0], (args)[cull:]
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
