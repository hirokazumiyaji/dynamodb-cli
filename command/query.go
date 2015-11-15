package command

import (
	"flag"
	"strings"

	"github.com/mitchellh/cli"
)

type QueryCommand struct {
	Ui cli.Ui
}

func (c *QueryCommand) Help() string {
	helpText := `
`

	return strings.TrimSpace(helpText)
}

func (c *QueryCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("query", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}
	return 0
}

func (c *QueryCommand) Synopsis() string {
	return ""
}
