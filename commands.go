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
}
