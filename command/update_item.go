package command

import (
	"flag"
	"strings"

	"github.com/mitchellh/cli"
)

type UpdateItemCommand struct {
	Ui cli.Ui
}

func (c *UpdateItemCommand) Help() string {
	helpText := `
`

	return strings.TrimSpace(helpText)
}

func (c *UpdateItemCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("update-item", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}
	return 0
}

func (c *UpdateItemCommand) Synopsis() string {
	return ""
}
