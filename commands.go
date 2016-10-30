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
		Name:    "show",
		Aliases: []string{"s"},
		Usage:   "Show illustrations from commit hashes",
		Action:  Show,
	},
	{
		Name:    "gif",
		Aliases: []string{"g"},
		Usage:   "Generate Gif of the production process",
		Action:  Gif,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "output, o",
				Value: "process.gif",
				Usage: "Output file `name`",
			},
			cli.IntFlag{
				Name:  "delay, d",
				Value: 1000,
				Usage: "Delay `time` (ms)",
			},
			cli.BoolFlag{
				Name:  "all",
				Usage: "Create pictures if there is no picture corresponding to commits (Not yet implemented)",
			},
		},
	},
}
