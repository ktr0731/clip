package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/urfave/cli"
)

func initClip(c *cli.Context) error {
	if IsExists(".clip/") {
		fmt.Println("Already initialized")

		return nil
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

	f.WriteString("\nclip -h")
	fmt.Println("Updated .git/hooks/post-commit")

	os.Chmod(".git/hooks/post-commit", 0755)

	return nil
}

func seekSQLiteHeader(data []byte) (int, error) {
	header := []byte{
		0x53, 0x51, 0x4c, 0x69, 0x74, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x20, 0x33,
	}

	n, at := 0, 0

	for i := range data {
		if data[i] == header[n] {
			if n == 0 {
				at = i
			}

			n++
		} else if n > 0 {
			n = 0
		}

		if n == len(header) {
			return at, nil
		}
	}

	return -1, fmt.Errorf("SQLite header not found")
}

func extractSQLiteDB() error {
	fileName, dbName := "sample.clip", "db"

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("Cannot open sample.clip")
	}

	at, err := seekSQLiteHeader(data)

	f, err := os.OpenFile(dbName, os.O_RDONLY, 0644)
	if err != nil {
		return fmt.Errorf("Cannot open db")
	}

	f.Write(data[at:])

	return nil
}

func export(c *cli.Context) error {
	fmt.Println("Export")
	fmt.Println("Extract db")

	extractSQLiteDB()
	return nil
}

func clean(c *cli.Context) error {
	fmt.Println("Clean")
	return nil
}

func diff(c *cli.Context) error {
	fmt.Println("Diff")
	return nil
}

// Commands Sub-commands for cli
var Commands = []cli.Command{
	{
		Name:    "init",
		Aliases: []string{"i"},
		Usage:   "Create .clip/ and update post-commit hook",
		Action:  initClip,
	},
	{
		Name:    "export",
		Aliases: []string{"e"},
		Usage:   "Export an illustration from latest .clip file",
		Action:  export,
	},
	{
		Name:    "clean",
		Aliases: []string{"c"},
		Usage:   "Remove not linked illustrations from .clip/",
		Action:  clean,
	},
	{
		Name:    "diff",
		Aliases: []string{"d"},
		Usage:   "Show changes between two commits",
		Action:  diff,
	},
}
