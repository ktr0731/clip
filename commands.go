package main

import "github.com/mitchellh/cli"

func commands(ui cli.Ui) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"init": func() (cli.Command, error) {
			return &InitCommand{ui}, nil
		},
		"export": func() (cli.Command, error) {
			return &ExportCommand{ui}, nil
		},
		"clean": func() (cli.Command, error) {
			return &CleanCommand{ui}, nil
		},
		"show": func() (cli.Command, error) {
			return &ShowCommand{ui}, nil
		},
		"gif": func() (cli.Command, error) {
			return &GifCommand{ui}, nil
		},
	}
}
