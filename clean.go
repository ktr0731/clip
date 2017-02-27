package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// clean removes some unnecessary files within .clip/
func clean(c *cli.Context) error {
	if err := os.RemoveAll(".clip"); err != nil {
		return err
	}

	fmt.Println("Delete .clip/")

	return nil
}
