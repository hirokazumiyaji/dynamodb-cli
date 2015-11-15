package command

import (
	"flag"
	"strings"

	"github.com/mitchellh/cli"
)

type CreateTableCommand struct {
	Ui cli.Ui
}

func (c *CreateTableCommand) Help() string {
	helpText := `
`

	return strings.TrimSpace(helpText)
}

func (c *CreateTableCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("create-table", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}
	return 0
}

func (c *CreateTableCommand) Synopsis() string {
	return ""
}
