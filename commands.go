package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func initClip(c *cli.Context) error {
	if IsExists(".clip/") {
		fmt.Println("Already initialized")

		return nil
	}

	if !IsExists(".git/hooks/") {
		return fmt.Errorf(".git/hooks/ Not Found")
	}

	return nil
}

func export(c *cli.Context) error {
	fmt.Println("Export")
	return nil
}

func clean(c *cli.Context) error {
	fmt.Println("Clean")
	return nil
}

func diff(c *cli.Context) error {
	fmt.Println("Diff")
	return nil
}

// Commands Sub-commands for cli
var Commands = []cli.Command{
	{
		Name:    "init",
		Aliases: []string{"i"},
		Usage:   "Create .clip/ and update post-commit hook",
		Action:  initClip,
	},
	{
		Name:    "export",
		Aliases: []string{"e"},
		Usage:   "Export an illustration from latest .clip file",
		Action:  export,
	},
	{
		Name:    "clean",
		Aliases: []string{"c"},
		Usage:   "Remove not linked illustrations from .clip/",
		Action:  clean,
	},
	{
		Name:    "diff",
		Aliases: []string{"d"},
		Usage:   "Show changes between two commits",
		Action:  diff,
	},
}
