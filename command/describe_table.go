package command

import (
	"flag"
	"strings"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/mitchellh/cli"
)

type DescribeTableCommand struct {
	Ui cli.Ui
}

func (c *DescribeTableCommand) Help() string {
	helpText := `
Usage: dynamodb-cli describe-table -table-name=<value>
`

	return strings.TrimSpace(helpText)
}

func (c *DescribeTableCommand) Run(args []string) int {
	var tableName string
	cmdFlags := flag.NewFlagSet("describe-table", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	cmdFlags.StringVar(&tableName, "table-name", "", "table name")
	if err := cmdFlags.Parse(args); err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	svc := dynamodb.New(session.New())
	params := &dynamodb.DescribeTableInput{
		TableName: tableName,
	}

	res, err := svc.DescribeTable(params)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}
	return 0
}

func (c *DescribeTableCommand) Synopsis() string {
	return ""
}
