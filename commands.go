package main

import "github.com/mitchellh/cli"

var commands = map[string]cli.CommandFactory{
	"init": func() (cli.Command, error) {
		return &InitCommand{}, nil
	},
	"export": func() (cli.Command, error) {
		return &ExportCommand{}, nil
	},
	"clean": func() (cli.Command, error) {
		return &CleanCommand{}, nil
	},
	"show": func() (cli.Command, error) {
		return &ShowCommand{}, nil
	},
	"gif": func() (cli.Command, error) {
		return &GifCommand{}, nil
	},
}
