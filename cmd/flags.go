package cmd

import (
	"GoWhen/cmd/cal"
	"errors"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"path/filepath"
)


const (
	flagFormat = "format"
	flagGoFormat = "go"
	flagCppFormat = "cpp"
	flagJavaFormat = "java"
)

func (cs *Cmds) AttachFlags(cmd *cobra.Command, viper *viper.Viper) {
	for range Only.Once {
		cmd.PersistentFlags().StringVarP(&cs.Data.FormatType, flagFormat, "", "go", fmt.Sprintf("Which layout to use for parsing and formatting."))
		viper.SetDefault(flagFormat, "go")
		// cmd.PersistentFlags().StringArrayVarP(&cs.Data.FormatType, flagFormat, "", []string{"go"}, fmt.Sprintf("Which layout to use for parsing and formatting."))
		// viper.SetDefault(flagFormat, "go")
		// cmd.PersistentFlags().BoolVarP(&cs.Data.CppFormat, flagCppFormat, "", false, fmt.Sprintf("Use CPP layout for parsing."))
		// viper.SetDefault(flagCppFormat, false)
		// cmd.PersistentFlags().BoolVarP(&cs.Data.JavaFormat, flagJavaFormat, "", false, fmt.Sprintf("Use Java layout for parsing."))
		// viper.SetDefault(flagJavaFormat, false)
	}
}

func (cs *Cmds) InitArgs(cmd *cobra.Command, _ []string) error {
	for range Only.Once {
		cs.Error = nil
		if cs.reparse {
			// We can re-execute commands inline, (particularly in a shell).
			break
		}
		cs.reparse = true

		cs.Data.SetDateIfNil()
		cs.Data.SetCmd(cmd.Name())
		switch cmd.Flag(flagFormat).Value.String() {
			case flagGoFormat:
				cs.Data.GoFormat = true
				cs.Data.CppFormat = false
				cs.Data.JavaFormat = false
			case flagCppFormat:
				cs.Data.GoFormat = false
				cs.Data.CppFormat = true
				cs.Data.JavaFormat = false
			case flagJavaFormat:
				cs.Data.GoFormat = false
				cs.Data.CppFormat = false
				cs.Data.JavaFormat = true
			default:
				cs.Error = errors.New(fmt.Sprintf("Unknown format layout. Can be only one of \"%s\", \"%s\" or \"%s\"",
					flagGoFormat, flagCppFormat, flagJavaFormat))
				break
		}

		if cs.Data.Convert == nil {
			fn := filepath.Join(cs.Unify.Commands.CmdConfig.Dir, "convert.json")
			cs.Data.Convert, cs.Error = cal.ReadConvert(fn)
			if cs.Error != nil {
				break
			}
		}
	}

	return cs.Error
}
