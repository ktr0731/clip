package main

import (
	"fmt"
	"os"
)

// CleanCommand removes some unnecessary files within .clip/
type CleanCommand struct{}

func (c *CleanCommand) Synopsis() string {
	return "Remove not linked illustrations from .clip/"
}

func (c *CleanCommand) Help() string {
	return "Usage: clip clean"
}

func (c *CleanCommand) Run(args []string) int {
	if err := os.RemoveAll(".clip"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Println("Deleted .clip/")

	return 0
}
