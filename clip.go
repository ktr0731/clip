package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

func main() {
	// app := cli.NewApp()
	// app.Name = "Clip"
	// app.Usage = "The content track helper for CLIP STUDIO files"
	// app.Commands = Commands
	// app.Version = "1.1.0"
	// app.Run(os.Args)
	c := cli.NewCLI("Clip", "1.1.0")
	c.Args = os.Args[1:]
	c.Commands = commands

	exitCode, err := c.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	os.Exit(exitCode)
}
