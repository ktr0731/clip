package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	pipeline "github.com/mattn/go-pipeline"
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
		return showWithFuzzySearch()
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
			if err := open.Run(path); err != nil {
				c.ui.Error(fmt.Sprintf("cannot open image: %s", err))
				return 1
			}
		} else {
			c.ui.Error(fmt.Sprintf("Invalid hash: %s\n", hash))
			return 1
		}
	}
	return 0
}

func showWithFuzzySearch() int {
	var finder string
	for _, cmd := range []string{"fzf", "fzy", "peco"} {
		b, err := exec.Command("which", cmd).Output()
		if err != nil {
			return 1
		}
		if len(b) != 0 {
			finder = cmd
			break
		}
	}
	if len(finder) == 0 {
		return 1
	}

	_, err := pipeline.Output(
		[]string{"git", "log", "--pretty=format:'%s %h'"},
		[]string{finder},
	)
	if err != nil {
		return 1
	}

	return 0
}
