package command

import (
	"flag"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/mitchellh/cli"
)

type ListTablesCommand struct {
	Ui cli.Ui
}

func (c *ListTablesCommand) Help() string {
	helpText := `
Usage: dynamodb-cli list-tables [options]

Options:

  -exclusive-start-table-name=<value>
  -limit=<value>
`

	return strings.TrimSpace(helpText)
}

func (c *ListTablesCommand) Run(args []string) int {
	var exclusiveStartTableName string
	var limit int

	cmdFlags := flag.NewFlagSet("list-tables", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	cmdFlags.StringVar(&exclusiveStartTableName, "exclusive-start-table-name", "", "exclusive start table name")
	cmdFlags.IntVar(&limit, "limit", 100, "limit")
	if err := cmdFlags.Parse(args); err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	svc := dynamodb.New(session.New())
	params := &dynamodb.ListTablesInput{
		ExclusiveStartTableName: aws.String(exclusiveStartTableName),
		Limit: limit,
	}

	res, err := svc.ListTables(params)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}
	c.Ui.Output(res)
	return 0
}

func (c *ListTablesCommand) Synopsis() string {
	return ""
}
