package main

import "os"

// IsExists Check whether the path is exists
func IsExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}
