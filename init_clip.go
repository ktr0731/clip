package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/cli"
)

// InitCommand creates .clip/ and update post-commit in .git/hooks/
type InitCommand struct {
	ui cli.Ui
}

func (c *InitCommand) Synopsis() string {
	return "Create .clip/ and update post-commit hook"
}

func (c *InitCommand) Help() string {
	return "Usage: clip init <target CLIP STUDIO PAINT file>"
}

func (c *InitCommand) Run(args []string) int {
	const clipDir = ".clip"

	if len(args) != 1 {
		c.ui.Error(c.Help())
		return 1
	}

	if isExists(clipDir) && isDir(clipDir) {
		c.ui.Error("Already initialized")
		return 1
	}

	if !isExists(".git") {
		c.ui.Error("Not a git repository")
		return 1
	}

	if !isExists(args[0]) {
		c.ui.Error(fmt.Sprintf("Target CLIP STUDIO PAINT file: %s not found", args[0]))
		return 1
	}

	hooksPath := filepath.Join(".git", "hooks")
	if !isExists(hooksPath) {
		c.ui.Error(hooksPath + " Not Found")
		return 1
	}

	postCommitPath := filepath.Join(hooksPath, "post-commit")
	postCommit, err := os.OpenFile(postCommitPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		c.ui.Error(fmt.Sprintf("Cannot open post-commit: %s\n", err))
		return 1
	}
	defer postCommit.Close()

	data := `# Clip https://github.com/lycoris0731/clip
NAME=$(git log -1 HEAD | head -1 | sed -e 's/commit //g')
clip export %s $NAME`

	postCommit.WriteString(fmt.Sprintf(string(data), args[0]))

	c.ui.Info("Updated post-commit")

	clipconfig, err := os.OpenFile(".clipconfig", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		c.ui.Error(fmt.Sprintf("Cannot open .clipconfig: %s\n", err))
		return 1
	}
	defer clipconfig.Close()

	clipconfig.WriteString(args[0])

	gitignore, err := os.OpenFile(".gitignore", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		c.ui.Error(fmt.Sprintf("Cannot open .gitignore: %s\n", err))
		return 1
	}
	defer gitignore.Close()

	gitignore.WriteString("# Clip\n.clip")
	c.ui.Info("Updated .gitignore")

	os.Chmod(postCommitPath, 0755)

	return 0
}
