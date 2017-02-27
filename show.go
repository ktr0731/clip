package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/skratchdot/open-golang/open"
)

// ShowCommand shows illustrations by commit hashes
type ShowCommand struct{}

func (c *ShowCommand) Synopsis() string {
	return "Show illustrations from commit hashes"
}

func (c *ShowCommand) Help() string {
	return "Usage: clip show [commit-hash ...]"
}

func (c *ShowCommand) Run(args []string) int {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, c.Help())
		return 1
	}

	for _, hash := range args {
		if strings.Contains(hash, "HEAD") {
			bytes, err := exec.Command("git", "rev-parse", hash).Output()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return 1
			}
			hash = strings.TrimSpace(string(bytes))
		}

		if isExists(fmt.Sprintf(".clip/%s", hash)) {
			open.Run(fmt.Sprintf(".clip/%s", hash))
		} else {
			fmt.Fprintln(os.Stderr, "Invalid hash")
			return 1
		}
	}
	return 0
}
