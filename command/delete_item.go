package command

import (
	"flag"
	"strings"

	"github.com/mitchellh/cli"
)

type DeleteItemCommand struct {
	Ui cli.Ui
}

func (c *DeleteItemCommand) Help() string {
	helpText := `
`

	return strings.TrimSpace(helpText)
}

func (c *DeleteItemCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("delete-item", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}
	return 0
}

func (c *DeleteItemCommand) Synopsis() string {
	return ""
}
