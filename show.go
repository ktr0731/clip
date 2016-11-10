package main

import (
	"fmt"
	"os"

	"github.com/skratchdot/open-golang/open"
	"github.com/urfave/cli"
)

// Show illustrations from commit hashes
func Show(c *cli.Context) {
	if c.NArg() > 0 {
		for _, hash := range c.Args() {
			if IsExists(fmt.Sprintf(".clip/%s", hash)) {
				open.Run(fmt.Sprintf(".clip/%s", hash))
			} else {
				fmt.Fprintln(os.Stderr, "Invalid hash")
				os.Exit(1)
			}
		}
	} else {
		fmt.Println("Usage: clip show [commit-hash ...]")
	}
}
