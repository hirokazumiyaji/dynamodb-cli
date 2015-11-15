package command

import (
	"flag"
	"strings"

	"github.com/mitchellh/cli"
)

type DeleteTableCommand struct {
	Ui cli.Ui
}

func (c *DeleteTableCommand) Help() string {
	helpText := `
`
	return strings.TrimSpace(helpText)
}

func (c *DeleteTableCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("delete-table", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	return 0
}

func (c *DeleteTableCommand) Synopsis() string {
	return ""
}
