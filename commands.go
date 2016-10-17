package main

import "github.com/urfave/cli"

// Commands Sub-commands for cli
var Commands = []cli.Command{
	{
		Name:    "init",
		Aliases: []string{"i"},
		Usage:   "Create .clip/ and update post-commit hook",
		Action:  InitClip,
	},
	{
		Name:    "export",
		Aliases: []string{"e"},
		Usage:   "Export an illustration from latest .clip file",
		Action:  Export,
	},
	{
		Name:    "clean",
		Aliases: []string{"c"},
		Usage:   "Remove not linked illustrations from .clip/",
		Action:  Clean,
	},
	{
		Name:    "diff",
		Aliases: []string{"d"},
		Usage:   "Show changes between two commits",
		Action:  Diff,
	},
}
