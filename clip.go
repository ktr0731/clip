package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("Clip", "1.2.0")
	c.Args = os.Args[1:]
	c.Commands = commands

	exitCode, err := c.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	os.Exit(exitCode)
}
