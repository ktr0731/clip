package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestSeekSQLiteHeader(t *testing.T) {
	expect := 3494297

	data, err := ioutil.ReadFile("tests/assets/sample.clip")
	if err != nil {
		t.Error(err)
	}

	actual, err := seekSQLiteHeader(data)
	if err != nil {
		t.Error(err)
	}

	if expect != actual {
		t.Errorf("Expect: %d, Actual: %d", expect, actual)
	}
}

func TestExtractSQLiteDB(t *testing.T) {
	if err := extractSQLiteDB("tests/assets/sample.clip"); err != nil {
		t.Error(err)
	}

	if err := os.Remove("db"); err != nil {
		t.Error(err)
	}
}

func TestExtractIllustration(t *testing.T) {
	if err := extractIllustration("tests/assets/db"); err != nil {
		t.Error(err)
	}

	if err := os.Remove(".clip/image.png"); err != nil {
		t.Error(err)
	}
}
