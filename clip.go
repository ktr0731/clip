package main

import (
  "fmt"
  "os"
)

func isExists(path string) bool {
  _, err := os.Stat(path)

  return err == nil
}

func clipInit() {
  if isExists(".clip/") {
    fmt.Println("Already initialized.")

    return
  }

  // if isExists(".git/hooks/") {
  //
  // }
}
