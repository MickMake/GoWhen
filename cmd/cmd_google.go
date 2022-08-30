package cmd

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdHelp"
	"GoWhen/cmd/cal"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


const (
	flagGoogleSheet       = "google-sheet"
	flagGoogleSheetUpdate = "update"
)

type CmdGoogle struct {
	GoogleSheet       string
	GoogleSheetUpdate bool

	OutputType string
	OutputFile string

	Error   error
	cmd     *cobra.Command
	SelfCmd *cobra.Command
}


func NewCmdGoogle() *CmdGoogle {
	var ret *CmdGoogle

	for range Only.Once {
		ret = &CmdGoogle{
			Error:   nil,
			cmd:     nil,
			SelfCmd: nil,
		}
	}

	return ret
}

func (w *CmdGoogle) AttachCommands(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		w.cmd = cmd

		// ******************************************************************************** //
		w.SelfCmd = &cobra.Command{
			Use:                   "google",
			Aliases:               []string{},
			Short:                 fmt.Sprintf("Update and view Google sheets."),
			Long:                  fmt.Sprintf("Update and view Google sheets."),
			DisableFlagParsing:    true, 
			DisableFlagsInUseLine: false,
			PreRunE:               w.InitArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				return w.SelfCmd.Help()
			},
			Args: cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(w.SelfCmd)
		w.SelfCmd.Example = cmdHelp.PrintExamples(w.SelfCmd, "update all", "update users")
		w.SelfCmd.Annotations = map[string]string{"group": "Google"}

		// ******************************************************************************** //
		var CmdGoogleUpdate = &cobra.Command{
			Use:                   "update",
			Aliases:               []string{"refresh"},
			Short:                 fmt.Sprintf("Update Google sheets."),
			Long:                  fmt.Sprintf("Update Google sheets."),
			DisableFlagParsing:    true, 
			DisableFlagsInUseLine: false,
			PreRunE:               w.InitArgs,
			Run:                   w.CmdGoogleUpdate,
			Args:                  cobra.MinimumNArgs(1),
		}
		w.SelfCmd.AddCommand(CmdGoogleUpdate)
		CmdGoogleUpdate.Example = cmdHelp.PrintExamples(CmdGoogleUpdate, "all", "presence", "user")
		CmdGoogleUpdate.Annotations = map[string]string{"group": "Google"}
	}

	return w.SelfCmd
}

func (w *CmdGoogle) AttachFlags(cmd *cobra.Command, viper *viper.Viper) {
	for range Only.Once {
		cmd.PersistentFlags().StringVarP(&w.GoogleSheet, flagGoogleSheet, "", "", fmt.Sprintf("Google: Sheet URL for updates."))
		viper.SetDefault(flagGoogleSheet, "")
		cmd.PersistentFlags().BoolVarP(&w.GoogleSheetUpdate, flagGoogleSheetUpdate, "", false, fmt.Sprintf("Update Google sheets."))
		viper.SetDefault(flagGoogleSheetUpdate, false)
		_ = cmd.PersistentFlags().MarkHidden(flagGoogleSheetUpdate)
	}
}

func (w *CmdGoogle) InitArgs(_ *cobra.Command, _ []string) error {
	for range Only.Once {
		if w.GoogleSheetUpdate {
			// w.OutputType = string(rune(cmdCloudFlare.TypeGit))
		}
	}

	return w.Error
}

func (w *CmdGoogle) CmdGoogleUpdate(cmd *cobra.Command, args []string) {
	for range Only.Once {
		switch {
			case len(args) == 0:
				w.Error = cmd.Help()

			case args[0] == "all":
				w.Error = w.GoogleUpdate(args...)

			default:
				fmt.Println("Unknown sub-command.")
				_ = cmd.Help()
		}
	}
}

func (w *CmdGoogle) GoogleUpdate(entities ...string) error {

	for range Only.Once {
		// cmds.CloudFlare.OutputType = TypeGoogle

		if len(entities) == 0 {
			entities = cal.TimeFormats
		}
		fmt.Printf("Saving %d entities from the PBX to Google Docs...\n", len(entities))

		// for _, entity := range entities {
		//	w.Error = cmds.Api.ReadDomains(domain)
		//	if w.Error != nil {
		//		break
		//	}
		//
		//	sheet := cmdGoogle.Sheet{
		//		Id:          "",
		//		Credentials: nil,
		//		SheetName:   entity,
		//		Range:       "",
		//		Data:        cmds.CloudFlare.OutputArray,
		//	}
		//	sheet.Set(sheet)
		//	w.Error = sheet.WriteSheet()
		//	if w.Error != nil {
		//		break
		//	}
		// }
	}

	return w.Error
}
