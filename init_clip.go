package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// InitClip create .clip/ and update post-commit in .git/hooks/
func InitClip(c *cli.Context) error {
	if c.NArg() != 1 {
		return fmt.Errorf("Usage: clip init TARGET_FILE")
	}
	if IsExists(".clip/") {
		return fmt.Errorf("Already initialized")
	}

	if !IsExists(".git/hooks/") {
		return fmt.Errorf(".git/hooks/ Not Found")
	}

	postCommit, err := os.OpenFile(".git/hooks/post-commit", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("Cannot open post-commit: %s", err)
	}
	defer postCommit.Close()

	data := `# Clip https://github.com/lycoris0731/clip
NAME=$(git log -1 HEAD | head -1 | sed -e 's/commit //g')
clip export %s $NAME`

	postCommit.WriteString(fmt.Sprintf(string(data), c.Args()[0]))

	fmt.Println("Updated .git/hooks/post-commit")

	clipconfig, err := os.OpenFile(".clipconfig", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("Cannot open .clipconfig: %s", err)
	}
	defer clipconfig.Close()

	clipconfig.WriteString(c.Args()[0])

	gitignore, err := os.OpenFile(".gitignore", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("Cannot open .gitignore: %s", err)
	}
	defer gitignore.Close()

	gitignore.WriteString("# Clip\n.clip")
	fmt.Println("Updated .gitignore")

	os.Chmod(".git/hooks/post-commit", 0755)

	return nil
}
