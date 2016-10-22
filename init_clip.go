package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/urfave/cli"
)

// InitClip create .clip/ and update post-commit in .git/hooks/
func InitClip(c *cli.Context) error {
	if IsExists(".clip/") {
		return fmt.Errorf("Already initialized")
	}

	if !IsExists(".git/hooks/") {
		return fmt.Errorf(".git/hooks/ Not Found")
	}

	os.Mkdir(".clip/", 0755)
	fmt.Println("Created .clip")

	f, err := os.OpenFile(".git/hooks/post-commit", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("Cannot open post-commit")
	}

	defer f.Close()

	data, err := ioutil.ReadFile("post-commit")
	if err != nil {
		return err
	}

	f.Write(data)

	fmt.Println("Updated .git/hooks/post-commit")

	os.Chmod(".git/hooks/post-commit", 0755)

	return nil
}
