package main

import (
	"fmt"

	"github.com/urfave/cli"
)

// Clean remove some unnecessary files in .clip/
func Clean(c *cli.Context) error {
	fmt.Println("Clean")
	return nil
}
