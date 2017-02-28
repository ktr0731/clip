package main

import (
	"bufio"
	"os"
	"testing"
)

func TestSeekSQLiteHeader(t *testing.T) {
	expect := 3494297

	f, err := os.Open("tests/assets/sample.clip")
	if err != nil {
		t.Error(err)
	}

	data := bufio.NewReader(f)

	actual, err := seekSQLiteHeader(data)
	if err != nil {
		t.Error(err)
	}

	if expect != actual {
		t.Errorf("Expect: %d, Actual: %d", expect, actual)
	}
}

func TestExtractSQLiteDB(t *testing.T) {
	at := 3494297

	tempFile, cleanup, err := makeTempFile()
	if err != nil {
		t.Error(err)
	}
	defer cleanup()

	f, err := os.Open("tests/assets/sample.clip")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		t.Error(err)
	}

	if err := extractSQLiteDB(tempFile, f, int64(at), stat.Size()); err != nil {
		t.Error(err)
	}
}

func TestExtractIllustration(t *testing.T) {
	const illustName string = "test"
	at := 3494297

	tempFile, cleanup, err := makeTempFile()
	if err != nil {
		t.Error(err)
	}
	defer cleanup()

	if !isExists(".clip") {
		if err := os.Mkdir(".clip", 0755); err != nil {
			t.Error(err)
		}
	}

	f, err := os.Open("tests/assets/sample.clip")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		t.Error(err)
	}

	if err := extractSQLiteDB(tempFile, f, int64(at), stat.Size()); err != nil {
		t.Error(err)
	}

	if err := extractIllustration(tempFile.Name(), illustName); err != nil {
		t.Error(err)
	}

	if err := os.RemoveAll(".clip"); err != nil {
		t.Error(err)
	}
}
