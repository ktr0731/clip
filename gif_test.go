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
	if err := generate("process.gif", 100, false); err != nil {
		t.Error(err)
	}

	if err := os.Remove("process.gif"); err != nil {
		t.Error(err)
	}
}
