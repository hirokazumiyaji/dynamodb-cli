package command

import (
	"flag"
	"strings"

	"github.com/mitchellh/cli"
)

type BatchWriteItemCommand struct {
	Ui cli.Ui
}

func (c *BatchWriteItemCommand) Help() string {
	helpText := `
`

	return strings.TrimSpace(helpText)
}

func (c *BatchWriteItemCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("batch-write-item", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}
	return 0
}

func (c *BatchWriteItemCommand) Synopsis() string {
	return ""
}
