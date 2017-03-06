package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/skratchdot/open-golang/open"
)

// ShowCommand shows illustrations by commit hashes
type ShowCommand struct {
	ui cli.Ui
}

func (c *ShowCommand) Synopsis() string {
	return "Show illustrations from commit hashes"
}

func (c *ShowCommand) Help() string {
	return "Usage: clip show [commit hash ...]"
}

func (c *ShowCommand) Run(args []string) int {
	if len(args) == 0 {
		c.ui.Error(c.Help())
		return 1
	}

	for _, hash := range args {
		if strings.Contains(hash, "HEAD") {
			bytes, err := exec.Command("git", "rev-parse", hash).Output()
			if err != nil {
				c.ui.Error(fmt.Sprint(err))
				return 1
			}
			hash = strings.TrimSpace(string(bytes))
		}

		path := filepath.Join(".clip", hash)
		if isExists(path) {
			open.Run(path)
		} else {
			c.ui.Error(fmt.Sprintf("Invalid hash: %s\n", hash))
			return 1
		}
	}
	return 0
}
