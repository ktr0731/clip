package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// Clean remove some unnecessary files in .clip/
func Clean(c *cli.Context) {
	if err := os.RemoveAll(".clip"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	} else {
		fmt.Println("Delete .clip/")
	}
}
