package main

import (
	// "log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Clip"
	app.Usage = "The content track helper for CLIP STUDIO files"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang, l",
			Usage: "lang for ..",
		},
	}

	app.Action = func(c *cli.Context) error {
		return nil
	}

	app.Commands = Commands

	app.Run(os.Args)
}
