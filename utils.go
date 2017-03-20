package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// isExists checks whether the path is exists
func isExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}

// isDir checks whether the path is directory or not
func isDir(path string) bool {
	stat, err := os.Stat(path)

	return err == nil && stat.IsDir()
}

// mkClipDir makes .clip directory
func mkClipDir() {
	os.Mkdir(".clip", 0755)
	fmt.Println("Created .clip")
}

// pickValidCommits picks all valid commits corresponding to pictures by asc
func pickValidCommits() ([]string, error) {
	result, err := exec.Command("git", "rev-list", "--all").Output()
	if err != nil {
		return nil, err
	}

	tmp := strings.Split(string(result), "\n")

	hashes := make([]string, len(tmp))
	for _, hash := range tmp[:len(tmp)] {
		if strings.TrimSpace(hash) == "" {
			continue
		}

		fmt.Println("hash: ", hash)
		if isExists(filepath.Join(".clip", hash)) {
			hashes = append(hashes, hash)
		}
	}

	return reverse(hashes), nil
}
