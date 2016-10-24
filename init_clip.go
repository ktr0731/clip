package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// InitClip create .clip/ and update post-commit in .git/hooks/
func InitClip(c *cli.Context) error {
	if c.NArg() != 1 {
		fmt.Println("Usage: clip init TARGET_FILE")
		os.Exit(1)
	}
	if IsExists(".clip/") {
		fmt.Println("Already initialized")
		os.Exit(1)
	}

	if !IsExists(".git/hooks/") {
		fmt.Println(".git/hooks/ Not Found")
		os.Exit(1)
	}

	os.Mkdir(".clip/", 0755)
	fmt.Println("Created .clip")

	f, err := os.OpenFile(".git/hooks/post-commit", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Cannot open post-commit")
		os.Exit(1)
	}

	defer f.Close()

	data := `# Clip https://github.com/lycoris0731/clip
NAME=$(git log -1 HEAD | head -1 | sed -e 's/commit //g')
clip export %s $NAME`

	f.WriteString(fmt.Sprintf(string(data), c.Args()[0]))

	fmt.Println("Updated .git/hooks/post-commit")

	os.Chmod(".git/hooks/post-commit", 0755)

	return nil
}
