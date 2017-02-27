package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// isExists checks whether the path is exists
func isExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}

// mkClipDir makes .clip directory
func mkClipDir() {
	os.Mkdir(".clip/", 0755)
	fmt.Println("Created .clip")
}

// pickValidCommits picks all valid commits corresponding to pictures by asc
func pickValidCommits() ([]string, error) {
	result, err := exec.Command("git", "rev-list", "--all").Output()
	if err != nil {
		return nil, err
	}

	tmp := strings.Split(string(result), "\n")

	var hashes []string
	for _, hash := range tmp[:len(tmp)-1] {
		if isExists(fmt.Sprintf(".clip/%s", hash)) {
			hashes = append(hashes, hash)
		}
	}

	return reverse(hashes), nil
}
