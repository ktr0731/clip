package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("Clip", "1.2.0")
	c.Args = os.Args[1:]

	ui := &cli.ColoredUi{
		InfoColor:  cli.UiColorBlue,
		ErrorColor: cli.UiColorRed,
		Ui: &cli.BasicUi{
			Reader:      os.Stdin,
			Writer:      os.Stdout,
			ErrorWriter: os.Stderr,
		},
	}

	c.Commands = commands(ui)

	exitCode, err := c.Run()
	if err != nil {
		ui.Error(fmt.Sprintf("Clip: %s\n", err))
	}
	os.Exit(exitCode)
}
