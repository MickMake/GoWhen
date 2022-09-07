package cmd

import (
	"GoWhen/cmd/cal"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"github.com/MickMake/GoUnify/cmdExec"
	"github.com/MickMake/GoUnify/cmdHelp"
	"github.com/spf13/cobra"
	"path/filepath"
	"time"
)


//goland:noinspection GoNameStartsWithPackageName
type CmdConvert CmdDefault

func NewCmdConvert() *CmdConvert {
	var ret *CmdConvert

	for range Only.Once {
		ret = &CmdConvert{
			Error:   nil,
			cmd:     nil,
			SelfCmd: nil,
		}
	}

	return ret
}

func (w *CmdConvert) AttachCommand(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		w.cmd = cmd

		// ******************************************************************************** //
		w.SelfCmd = &cobra.Command{
			Use:                   "convert",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Convert"},
			Short:                 "Date/time format conversion tables.",
			Long:                  "Date/time format conversion tables.",
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdConvert,
			Args:                  cobra.MinimumNArgs(0),
		}
		cmd.AddCommand(w.SelfCmd)
		w.SelfCmd.Example = cmdHelp.PrintExamples(w.SelfCmd,
			"test . '1967-07-01 09:00:00'",
			"test '%F %T' '1967-07-01 09:00:00'",
			"test 'yyyy-MM-dd HH:mm:ss' '2022-12-31 09:00:00'\"",
			"test",
			"list",
		)

		// ******************************************************************************** //
		var CmdConvertList = &cobra.Command{
			Use:                   "list",
			Aliases:               []string{"ls"},
			Annotations:           map[string]string{"group": "Convert"},
			Short:                 "Show format conversion tables.",
			Long:                  "Show format conversion tables.",
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdConvertList,
			Args:                  cobra.MinimumNArgs(0),
		}
		w.SelfCmd.AddCommand(CmdConvertList)
		CmdConvertList.Example = cmdHelp.PrintExamples(CmdConvertList,
			"",
			)

		// ******************************************************************************** //
		var CmdConvertLayout = &cobra.Command{
			Use:                   "layouts",
			Aliases:               []string{"layout"},
			Annotations:           map[string]string{"group": "Convert"},
			Short:                 "Show predefined GoLang layouts.",
			Long:                  "Show predefined GoLang layouts.",
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdConvertLayouts,
			Args:                  cobra.MinimumNArgs(0),
		}
		w.SelfCmd.AddCommand(CmdConvertLayout)
		CmdConvertLayout.Example = cmdHelp.PrintExamples(CmdConvertLayout,
			"",
		)

		// ******************************************************************************** //
		var CmdConvertFormat = &cobra.Command{
			Use:                   "options",
			Aliases:               []string{"option"},
			Annotations:           map[string]string{"group": "Convert"},
			Short:                 "Show GoLang layout options.",
			Long:                  "Show GoLang layout options.",
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdConvertOptions,
			Args:                  cobra.MinimumNArgs(0),
		}
		w.SelfCmd.AddCommand(CmdConvertFormat)
		CmdConvertFormat.Example = cmdHelp.PrintExamples(CmdConvertFormat,
			"",
		)

		// ******************************************************************************** //
		var CmdConvertTable = &cobra.Command{
			Use:                   "table",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Convert"},
			Short:                 "Show conversion table.",
			Long:                  "Show conversion table.",
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdConvertTable,
			Args:                  cobra.MinimumNArgs(0),
		}
		w.SelfCmd.AddCommand(CmdConvertTable)
		CmdConvertTable.Example = cmdHelp.PrintExamples(CmdConvertTable,
			"",
		)

		// ******************************************************************************** //
		var CmdConvertJson = &cobra.Command{
			Use:                   "json",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Convert"},
			Short:                 "Show conversion table as JSON.",
			Long:                  "Show conversion table as JSON.",
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdConvertJson,
			Args:                  cobra.MinimumNArgs(0),
		}
		w.SelfCmd.AddCommand(CmdConvertJson)
		CmdConvertJson.Example = cmdHelp.PrintExamples(CmdConvertJson,
			"",
		)

		// ******************************************************************************** //
		var CmdConvertTest = &cobra.Command{
			Use:                   "test <format> <date/time>",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Convert"},
			Short:                 "Test format conversion tables.",
			Long:                  "Test format conversion tables.",
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.InitArgs,
			RunE:                  cmds.CmdConvertTest,
			Args:                  cobra.MinimumNArgs(0),
		}
		w.SelfCmd.AddCommand(CmdConvertTest)
		CmdConvertTest.Example = cmdHelp.PrintExamples(CmdConvertTest,
			". '1967-07-01 09:00:00'",
			"",
			"'%F %T' '1967-07-01 09:00:00'",
			"'yyyy-MM-dd HH:mm:ss' '2022-12-31 09:00:00'\"",
			)

	}

	return w.SelfCmd
}

func (cs *Cmds) CmdConvert(_ *cobra.Command, _ []string) error {
	for range Only.Once {
		_ = cs.Convert.SelfCmd.Help()
	}

	return cs.Error
}

func (cs *Cmds) CmdConvertList(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		// var arg []string
		// arg, args = cmdExec.PopArgs(1, args)
		// ######################################## //


		fmt.Println("\n############################################################")
		cs.Error = cs.CmdConvertTable(cmd, args)
		if cs.Error != nil {
			break
		}
		fmt.Println("\n############################################################")
		cs.Error = cs.CmdConvertLayouts(cmd, args)
		if cs.Error != nil {
			break
		}
		fmt.Println("\n############################################################")
		cs.Error = cs.CmdConvertOptions(cmd, args)
		if cs.Error != nil {
			break
		}


		// ######################################## //
		// cs.last, cs.Error = cmdExec.ReparseArgs(cmd, args)
		// cs.LastPrint()
	}

	return cs.Error
}

func (cs *Cmds) CmdConvertLayouts(_ *cobra.Command, _ []string) error {
	for range Only.Once {
		// var arg []string
		// arg, args = cmdExec.PopArgs(1, args)
		// ######################################## //


		fmt.Println(cal.PrintLayouts())


		// ######################################## //
		// cs.last, cs.Error = cmdExec.ReparseArgs(cmd, args)
		// cs.LastPrint()
	}

	return cs.Error
}

func (cs *Cmds) CmdConvertOptions(_ *cobra.Command, _ []string) error {
	for range Only.Once {
		// var arg []string
		// arg, args = cmdExec.PopArgs(1, args)
		// ######################################## //


		fmt.Println(cal.PrintLayoutOptions(nil))


		// ######################################## //
		// cs.last, cs.Error = cmdExec.ReparseArgs(cmd, args)
		// cs.LastPrint()
	}

	return cs.Error
}

func (cs *Cmds) CmdConvertTable(_ *cobra.Command, args []string) error {
	for range Only.Once {
		var arg []string
		arg, args = cmdExec.PopArgs(2, args)
		// ######################################## //


		fn := filepath.Join(cs.Unify.Commands.CmdConfig.Dir, "convert.json")
		if arg[0] != "" {
			fn = arg[0]
		}

		var cnv *cal.Convert
		cnv, cs.Error = cal.ReadConvert(fn)
		if cs.Error != nil {
			break
		}
		fmt.Println(cnv.String())

		fmt.Println("To select a different date/time layout. Do one of the following:")
		fmt.Printf("%s config write --format=go\n", cs.Convert.cmd.Name())
		fmt.Printf("%s config write --format=cpp\n", cs.Convert.cmd.Name())
		fmt.Printf("%s config write --format=java\n", cs.Convert.cmd.Name())


		// ######################################## //
		// cs.last, cs.Error = cmdExec.ReparseArgs(cmd, args)
		// cs.LastPrint()
	}

	return cs.Error
}

func (cs *Cmds) CmdConvertJson(_ *cobra.Command, args []string) error {
	for range Only.Once {
		var arg []string
		arg, args = cmdExec.PopArgs(2, args)
		// ######################################## //


		fn := filepath.Join(cs.Unify.Commands.CmdConfig.Dir, "convert.json")
		if arg[0] != "" {
			fn = arg[0]
		}

		var cnv *cal.Convert
		cnv, cs.Error = cal.ReadConvert(fn)
		if cs.Error != nil {
			break
		}
		fmt.Println(cnv.Json())


		// ######################################## //
		// cs.last, cs.Error = cmdExec.ReparseArgs(cmd, args)
		// cs.LastPrint()
	}

	return cs.Error
}

func (cs *Cmds) CmdConvertTest(_ *cobra.Command, args []string) error {
	for range Only.Once {
		var arg []string
		arg, args = cmdExec.PopArgs(2, args)
		// ######################################## //


		now := time.Now()
		if arg[0] == "" {
			fmt.Printf("Java input: \"EEE d MMM yyyy HH:mm:ss z\"\n")
			fmt.Printf("Go output: \"%s\"\n", cs.Data.Convert.FromJava("EEE d MMM yyyy HH:mm:ss z"))
			fmt.Println("")
			fmt.Printf("CPP input: \"%%a %%e %%b %%Y %%T %%Z\"\n")
			fmt.Printf("Go output: \"%s\"\n", cs.Data.Convert.FromCpp("%a %e %b %Y %T %Z"))
		} else {
			cs.Data.ConvertFormat(arg[0])
			fmt.Printf("%s input:\t\"%s\"\n", cs.Data.FormatType, arg[0])
			fmt.Printf("Go output:\t\"%s\"\n", cs.Data.Format)
		}

		if arg[1] != "" {
			cs.Error = cs.Data.DateParse(arg[0], arg[1])
			if cs.Error != nil {
				break
			}
			now = *cs.Data.FromDate.Time
		}

		fmt.Println("")
		fmt.Println(cal.PrintLayoutOptions(&now))


		// ######################################## //
		// cs.last, cs.Error = cmdExec.ReparseArgs(cmd, args)
		// cs.LastPrint()
	}

	return cs.Error
}
