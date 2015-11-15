package command

import (
	"flag"
	"strings"

	"github.com/mitchellh/cli"
)

type ScanCommand struct {
	Ui cli.Ui
}

func (c *ScanCommand) Help() string {
	helpText := `
`

	return strings.TrimSpace(helpText)
}

func (c *ScanCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("scan", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}
	return 0
}

func (c *ScanCommand) Synopsis() string {
	return ""
}
