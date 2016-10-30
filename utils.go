package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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

// PickValidCommits pick all valid commits corresponding to pictures by asc
func PickValidCommits() ([]string, error) {
	result, err := exec.Command("git", "rev-list", "--all").Output()
	if err != nil {
		return nil, err
	}

	tmp := strings.Split(string(result), "\n")

	var hashes []string
	for _, hash := range tmp[:len(tmp)-1] {
		if IsExists(fmt.Sprintf(".clip/%s", hash)) {
			hashes = append(hashes, hash)
		}
	}

	return reverse(hashes), nil
}
