package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// InitClip create .clip/ and update post-commit in .git/hooks/
func InitClip(c *cli.Context) {
	if c.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "Usage: clip init TARGET_FILE")
		os.Exit(1)
	}
	if IsExists(".clip/") {
		fmt.Fprintln(os.Stderr, "Already initialized")
		os.Exit(1)
	}

	if !IsExists(".git/hooks/") {
		fmt.Fprintln(os.Stderr, ".git/hooks/ Not Found")
		os.Exit(1)
	}

	postCommit, err := os.OpenFile(".git/hooks/post-commit", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot open post-commit")
		os.Exit(1)
	}
	defer postCommit.Close()

	data := `# Clip https://github.com/lycoris0731/clip
NAME=$(git log -1 HEAD | head -1 | sed -e 's/commit //g')
clip export %s $NAME`

	postCommit.WriteString(fmt.Sprintf(string(data), c.Args()[0]))

	fmt.Println("Updated .git/hooks/post-commit")

	clipconfig, err := os.OpenFile(".clipconfig", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot open .clipconfig")
		os.Exit(1)
	}
	defer clipconfig.Close()

	clipconfig.WriteString(c.Args()[0])

	gitignore, err := os.OpenFile(".gitignore", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot open .gitignore")
		os.Exit(1)
	}
	defer gitignore.Close()

	gitignore.WriteString("# Clip\n.clip")
	fmt.Println("Updated .gitignore")

	os.Chmod(".git/hooks/post-commit", 0755)
}
