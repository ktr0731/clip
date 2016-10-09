package main

import (
	"fmt"

	"github.com/urfave/cli"
)

// Commands Sub-commands for cli
var Commands = []cli.Command{
	{
		Name:    "init",
		Aliases: []string{"i"},
		Usage:   "Create .clip/ and update post-commit hook",
		Action: func(c *cli.Context) error {
			fmt.Println("Init!!")
			return nil
		},
	},
	{
		Name:    "export",
		Aliases: []string{"e"},
		Usage:   "Export an illustration from latest .clip file",
		Action: func(c *cli.Context) error {
			fmt.Println("Export")
			return nil
		},
	},
	{
		Name:    "clean",
		Aliases: []string{"c"},
		Usage:   "Remove not linked illustrations from .clip/",
		Action: func(c *cli.Context) error {
			fmt.Println("Clean")
			return nil
		},
	},
	{
		Name:    "diff",
		Aliases: []string{"d"},
		Usage:   "Show changes between two commits",
		Action: func(c *cli.Context) error {
			fmt.Println("Diff")
			return nil
		},
	},
}
