package main

import (
	"os"
	"testing"
)

func TestReverse(t *testing.T) {
	actual := []string{"Makise", "Kurisu", "ha", "kawaii"}

	expect := reverse(actual)
	if actual[0] != expect[len(actual)-1] {
		t.Errorf("Expect: %s, Actual: %s", expect[len(actual)-1], actual[0])
	}
}

func TestGenerate(t *testing.T) {
	generate := &GifCommand{}
	if status := generate.Run([]string{}); status != 0 {
		t.Error("exit with 1")
	}

	if err := os.Remove("process.gif"); err != nil {
		t.Error(err)
	}
}
