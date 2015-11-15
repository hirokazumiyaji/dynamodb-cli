package command

import (
	"flag"
	"strings"

	"github.com/mitchellh/cli"
)

type ListTablesCommand struct {
	Ui cli.Ui
}

func (c *ListTablesCommand) Help() string {
	helpText := `
	Usage: dynamodb-cli list-tables [options]

Options:

  -cli-input-json=<value>
  -start-token=<value>
  -page-size=<value>
  -max-items=<value>
  -generate-cli-skeleton=true/false
`

	return strings.TrimSpace(helpText)
}

func (c *ListTablesCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("list-tables", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}
	return 0
}

func (c *ListTablesCommand) Synopsis() string {
	return ""
}
