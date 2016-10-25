package main

import (
	"fmt"
	"os"
)

// IsExists Check whether the path is exists
func IsExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}

// MkClipDir make .clip directory
func MkClipDir() {
	os.Mkdir(".clip/", 0755)
	fmt.Println("Created .clip")
}
