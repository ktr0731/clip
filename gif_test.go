package main

import "testing"

func TestReverse(t *testing.T) {
	actual := []string{"Makise", "Kurisu", "ha", "kawaii"}

	expect := reverse(actual)
	if actual[0] != expect[len(actual)-1] {
		t.Errorf("Expect: %s, Actual: %s", expect[len(actual)-1], actual[0])
	}
}

func TestGenerate(t *testing.T) {
	if err := generate(); err != nil {
		t.Error(err)
	}
}
