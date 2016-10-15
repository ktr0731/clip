package main

import "os"

// IsExists Check exists by path name
func IsExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}
