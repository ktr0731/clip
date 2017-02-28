package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const dbName = "db"

// ExportCommand creates image file from CLIP STUDIO file
type ExportCommand struct{}

func (c *ExportCommand) Synopsis() string {
	return "Export an illustration from latest .clip file"
}

func (c *ExportCommand) Help() string {
	return "Usage: clip export CLIP_STUDIO_FILE IMG_NAME"
}

func (c *ExportCommand) Run(args []string) int {
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, c.Help())
		return 1
	}

	clipFileName := args[0]
	outputFileName := args[1]

	if !isExists(".clip") {
		mkClipDir()
	}

	if !isExists(clipFileName) {
		fmt.Fprintf(os.Stderr, "%s: no such file\n", clipFileName)
		return 1
	}

	f, err := os.Open(clipFileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	i, err := seekSQLiteHeader(buf)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	stat, err := f.Stat()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := extractSQLiteDB(f, int64(i), stat.Size()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	if err := extractIllustration(outputFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	if err := os.Remove(dbName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	return 0
}

func seekSQLiteHeader(buf io.ByteReader) (int, error) {
	header := []byte{
		0x53, 0x51, 0x4c, 0x69, 0x74, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x20, 0x33,
	}

	n, at, i := 0, 0, 0
	for {
		b, err := buf.ReadByte()
		if err == io.EOF {
			return -1, fmt.Errorf("SQLite header not found")
		}

		if err != nil {
			return -1, err
		}

		if b == header[n] {
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

		i++
	}
}

func extractSQLiteDB(r io.ReaderAt, at, size int64) error {
	data := make([]byte, size)
	_, err := r.ReadAt(data, at)
	if err != nil && err != io.EOF {
		return err
	}

	fmt.Println(len(data))

	// TODO: ioutil.TempFile 使う
	f, err := os.OpenFile(dbName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	f.Write(data)

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
