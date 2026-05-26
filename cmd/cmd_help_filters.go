package cmd

import (
	"fmt"
	"strings"

	"GoWhen/defaults"
	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/MickMake/GoUnify/Only"
	"github.com/spf13/cobra"
)

func AttachFilterHelpCommand(cmd *cobra.Command) {
	for range Only.Once {
		if cmd == nil {
			break
		}

		cmdFilters := &cobra.Command{
			Use:                   "filters",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Help"},
			Short:                 "Filter and piped stdin help",
			Long:                  defaults.Filters,
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			RunE:                  CmdHelpFilters,
			Args:                  cobra.RangeArgs(0, 0),
			Hidden:                true,
		}
		cmd.AddCommand(cmdFilters)
	}
}

func CmdHelpFilters(_ *cobra.Command, _ []string) error {
	w := getDocWidth(defaults.Filters)
	result := markdown.Render(defaults.Filters, w, 6)
	fmt.Printf("%s", result)
	return nil
}

func getDocWidth(text string) int {
	width := 80
	for _, line := range strings.Split(text, "\n") {
		if len(line) > width {
			width = len(line)
		}
	}
	return width
}
