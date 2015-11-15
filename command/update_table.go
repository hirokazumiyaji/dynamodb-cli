package command

import (
	"flag"
	"strings"

	"github.com/mitchellh/cli"
)

type UpdateTableCommand struct {
	Ui cli.Ui
}

func (c *UpdateTableCommand) Help() string {
	helpText := `
`

	return strings.TrimSpace(helpText)
}

func (c *UpdateTableCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("update-table", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}
	return 0
}

func (c *UpdateTableCommand) Synopsis() string {
	return ""
}
