package main

import (
	"os"

	"github.com/hirokazumiyaji/dynamodb-cli/command"
	"github.com/mitchellh/cli"
)

var Commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{Writer: os.Stdout}

	Commands = map[string]cli.CommandFactory{
		"batch-get-item": func() (cli.Command, error) {
			return &command.BatchGetItemCommand{
				Ui: ui,
			}, nil
		},
		"batch-write-item": func() (cli.Command, error) {
			return &command.BatchWriteItemCommand{
				Ui: ui,
			}, nil
		},
		"create-table": func() (cli.Command, error) {
			return &command.CreateTableCommand{
				Ui: ui,
			}, nil
		},
		"delete-item": func() (cli.Command, error) {
			return &command.DeleteItemCommand{
				Ui: ui,
			}, nil
		},
		"delete-table": func() (cli.Command, error) {
			return &command.DeleteTableCommand{
				Ui: ui,
			}, nil
		},
		"describe-table": func() (cli.Command, error) {
			return &command.DescribeTableCommand{
				Ui: ui,
			}, nil
		},
		"get-item": func() (cli.Command, error) {
			return &command.GetItemCommand{
				Ui: ui,
			}, nil
		},
		"list-tables": func() (cli.Command, error) {
			return &command.ListTablesCommand{
				Ui: ui,
			}, nil
		},
		"put-item": func() (cli.Command, error) {
			return &command.PutItemCommand{
				Ui: ui,
			}, nil
		},
		"query": func() (cli.Command, error) {
			return &command.QueryCommand{
				Ui: ui,
			}, nil
		},
		"scan": func() (cli.Command, error) {
			return &command.ScanCommand{
				Ui: ui,
			}, nil
		},
		"update-item": func() (cli.Command, error) {
			return &command.UpdateItemCommand{
				Ui: ui,
			}, nil
		},
		"update-table": func() (cli.Command, error) {
			return &command.UpdateTableCommand{
				Ui: ui,
			}, nil
		},
		"wait": func() (cli.Command, error) {
			return &command.WaitCommand{
				Ui: ui,
			}, nil
		},
		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Revision:          GitCommit,
				Version:           Version,
				VersionPrerelease: VersionPrerelease,
				Ui:                ui,
			}, nil
		},
	}
}
