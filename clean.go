package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// Clean remove some unnecessary files in .clip/
func Clean(c *cli.Context) error {
	if err := os.RemoveAll(".clip"); err != nil {
		return err
	}

	fmt.Println("Delete .clip/")

	return nil
}
