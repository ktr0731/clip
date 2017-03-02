package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// InitCommand creates .clip/ and update post-commit in .git/hooks/
type InitCommand struct{}

func (c *InitCommand) Synopsis() string {
	return "Create .clip/ and update post-commit hook"
}

func (c *InitCommand) Help() string {
	return "Usage: clip init TARGET_FILE"
}

func (c *InitCommand) Run(args []string) int {
	const clipDir = ".clip"

	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, c.Help())
		return 1
	}

	if isExists(clipDir) && isDir(clipDir) {
		fmt.Fprintln(os.Stderr, "Already initialized")
		return 1
	}

	hooksPath := filepath.Join(".git", "hooks")
	if !isExists(hooksPath) {
		fmt.Fprintln(os.Stderr, hooksPath+" Not Found")
		return 1
	}

	postCommitPath := filepath.Join(hooksPath, "post-commit")
	postCommit, err := os.OpenFile(postCommitPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot open post-commit: %s\n", err)
		return 1
	}
	defer postCommit.Close()

	data := `# Clip https://github.com/lycoris0731/clip
NAME=$(git log -1 HEAD | head -1 | sed -e 's/commit //g')
clip export %s $NAME`

	postCommit.WriteString(fmt.Sprintf(string(data), args[0]))

	fmt.Println("Updated post-commit")

	clipconfig, err := os.OpenFile(".clipconfig", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot open .clipconfig: %s\n", err)
		return 1
	}
	defer clipconfig.Close()

	clipconfig.WriteString(args[0])

	gitignore, err := os.OpenFile(".gitignore", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot open .gitignore: %s\n", err)
		return 1
	}
	defer gitignore.Close()

	gitignore.WriteString("# Clip\n.clip")
	fmt.Println("Updated .gitignore")

	os.Chmod(postCommitPath, 0755)

	return 0
}
