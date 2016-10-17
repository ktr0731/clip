package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/urfave/cli"
)

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

func extractSQLiteDB() error {
	fileName, dbName := "sample.clip", "db"

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("Cannot open sample.clip")
	}

	at, err := seekSQLiteHeader(data)

	f, err := os.OpenFile(dbName, os.O_RDONLY, 0644)
	if err != nil {
		return fmt.Errorf("Cannot open db")
	}

	f.Write(data[at:])

	return nil
}

// Export create image file from CLIP STUDIO file
func Export(c *cli.Context) error {
	fmt.Println("Export")
	fmt.Println("Extract db")

	extractSQLiteDB()
	return nil
}
