package command

import (
	"flag"
	"strings"

	"github.com/mitchellh/cli"
)

type BatchGetItemCommand struct {
	Ui cli.Ui
}

func (c *BatchGetItemCommand) Help() string {
	helpText := `
Usage: dynamodb-cli batch-get-item -table_name=<value> -key=<value> [options]

Options:

  -attributes-to-get=<value>
  -consistent-read=true/false
  -no-consistent-read=true/false
  -return-consumed-capacity=<value>
  -projection-expression=<value>
  -expression-attribute-names=<value>
  -cli-input-json=value
  -generate-cli-skeleton

`
	return strings.TrimSpace(helpText)
}

func (c *BatchGetItemCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("batch-get-item", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	return 0
}

func (c *BatchGetItemCommand) Synopsis() string {
	return ""
}
