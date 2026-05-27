package cmd

import (
	"fmt"
	"strings"

	"GoWhen/defaults"
	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/MickMake/GoUnify/Only"
	"github.com/spf13/cobra"
)

const filtersSectionHeading = "## Filters and piped stdin"

func AttachFilterHelpCommand(cmd *cobra.Command) {
	for range Only.Once {
		if cmd == nil {
			break
		}

		filters := filterHelpMarkdown()
		cmdFilters := &cobra.Command{
			Use:                   "filters",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Help"},
			Short:                 "Filter and piped stdin help",
			Long:                  filters,
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
	filters := filterHelpMarkdown()
	w := getDocWidth(filters)
	result := markdown.Render(filters, w, 6)
	fmt.Printf("%s", result)
	return nil
}

func filterHelpMarkdown() string {
	start := strings.Index(defaults.Readme, filtersSectionHeading)
	if start == -1 {
		return defaults.Readme
	}

	section := defaults.Readme[start:]
	if next := strings.Index(section[len(filtersSectionHeading):], "\n## "); next != -1 {
		section = section[:len(filtersSectionHeading)+next]
	}

	return strings.TrimSpace(section) + "\n"
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