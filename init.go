package main

import (
	"fmt"
	"os"
)

func isExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}

func clipInit() error {
	if isExists(".clip/") {
		fmt.Println("Already initialized")

		return nil
	}

	if isExists(".git/hooks/") {
		return fmt.Errorf(".git/hooks/ Not Found")
	}

	return nil
}
