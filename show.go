package main

import (
	"fmt"

	"github.com/skratchdot/open-golang/open"
	"github.com/urfave/cli"
)

// Show illustrations from commit hashes
func Show(c *cli.Context) error {
	if c.NArg() > 0 {
		for _, hash := range c.Args() {
			if IsExists(fmt.Sprintf(".clip/%s", hash)) {
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
