package command

import (
	"flag"
	"strings"

	"github.com/mitchellh/cli"
)

type GetItemCommand struct {
	Ui cli.Ui
}

func (c *GetItemCommand) Help() string {
	helpText := `
`

	return strings.TrimSpace(helpText)
}

func (c *GetItemCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("get-item", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}
	return 0
}

func (c *GetItemCommand) Synopsis() string {
	return ""
}
