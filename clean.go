package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

// CleanCommand removes some unnecessary files within .clip/
type CleanCommand struct {
	ui cli.Ui
}

func (c *CleanCommand) Synopsis() string {
	return "Remove not linked illustrations from .clip/"
}

func (c *CleanCommand) Help() string {
	return "Usage: clip clean"
}

func (c *CleanCommand) Run(args []string) int {
	if err := os.RemoveAll(".clip"); err != nil {
		c.ui.Error(fmt.Sprint(err))
		return 1
	}

	c.ui.Info("Deleted .clip/")

	return 0
}
