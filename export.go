package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli"
)

const dbName = "db"

func seekSQLiteHeader(data []byte) (int, error) {
	header := []byte{
		0x53, 0x51, 0x4c, 0x69, 0x74, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x20, 0x33,
	}

	n, at := 0, 0

	for i := range data {
		if data[i] == header[n] {
			if n == 0 {
				at = i
			}

			n++
		} else if n > 0 {
			n = 0
		}

		if n == len(header) {
			return at, nil
		}
	}

	return -1, fmt.Errorf("SQLite header not found")
}

func extractSQLiteDB(fileName string) error {
	if !IsExists(fileName) {
		return fmt.Errorf("%s: no such file", fileName)
	}

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	at, err := seekSQLiteHeader(data)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(dbName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	f.Write(data[at:])

	return nil
}

func extractIllustration(illustName string) error {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return err
	}
	defer db.Close()

	f, err := os.OpenFile(fmt.Sprintf(".clip/%s", illustName), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	var length int
	err = db.QueryRow("select length(ImageData) from CanvasPreview").Scan(&length)
	if err != nil {
		return err
	}

	image := make([]byte, length)
	err = db.QueryRow("select ImageData from CanvasPreview").Scan(&image)
	if err != nil {
		return err
	}

	f.Write(image)

	return nil
}

// Export create image file from CLIP STUDIO file
func Export(c *cli.Context) error {
	if c.NArg() != 2 {
		fmt.Println("Usage: clip export CLIP_STUDIO_FILE IMG_NAME")
		os.Exit(1)
	}

	if !IsExists(".clip") {
		fmt.Println("Directory .clip not found\nPlease do\n\n\tclip init TARGET_FILE\n")
		os.Exit(1)
	}

	if err := extractSQLiteDB(c.Args()[0]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := extractIllustration(c.Args()[1]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := os.Remove(dbName); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// open.Run(fmt.Sprintf(".clip/%s", c.Args()[1]))

	return nil
}
