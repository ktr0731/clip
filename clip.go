package main

import (
	"fmt"
	"log"
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

func main() {
	err := clipInit()
	if err != nil {
		log.Fatal(err)
	}
}
