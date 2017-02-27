package main

import (
	"fmt"

	"github.com/skratchdot/open-golang/open"
	"github.com/urfave/cli"
)

// show illustrations by commit hashes
func show(c *cli.Context) error {
	if c.NArg() > 0 {
		for _, hash := range c.Args() {
			if isExists(fmt.Sprintf(".clip/%s", hash)) {
				open.Run(fmt.Sprintf(".clip/%s", hash))
			} else {
				return fmt.Errorf("Invalid hash")
			}
		}
	} else {
		return fmt.Errorf("Usage: clip show [commit-hash ...]")
	}

	return nil
}
