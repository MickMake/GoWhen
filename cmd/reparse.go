package cmd

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdCron"
	"github.com/spf13/cobra"
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

func (cs *Cmds) PopArgs1(args []string) (string, []string) {
	if len(args) == 0 {
		return "", args
	}
	return (args)[0], (args)[1:]
}

func (cs *Cmds) PopArgs(cull int, args []string) ([]string, []string) {
	if cull > len(args) {
		return []string{}, args
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
