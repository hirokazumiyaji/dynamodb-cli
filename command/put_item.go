package command

import (
	"flag"
	"strings"

	"github.com/mitchellh/cli"
)

type PutItemCommand struct {
	Ui cli.Ui
}

func (c *PutItemCommand) Help() string {
	helpText := `
`

	return strings.TrimSpace(helpText)
}

func (c *PutItemCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("put-item", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}
	return 0
}

func (c *PutItemCommand) Synopsis() string {
	return ""
}
