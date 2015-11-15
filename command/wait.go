package command

import (
	"flag"
	"strings"

	"github.com/mitchellh/cli"
)

type WaitCommand struct {
	Ui cli.Ui
}

func (c *WaitCommand) Help() string {
	helpText := `
`

	return strings.TrimSpace(helpText)
}

func (c *WaitCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("wait", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}
	return 0
}

func (c *WaitCommand) Synopsis() string {
	return ""
}
