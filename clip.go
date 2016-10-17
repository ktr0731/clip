package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Clip"
	app.Usage = "The content track helper for CLIP STUDIO files"
	app.Commands = Commands
	app.Run(os.Args)
}
