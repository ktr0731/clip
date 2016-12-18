package main

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/png"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli"
)

func reverse(s []string) []string {
	result := make([]string, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		result[len(s)-i-1] = s[i]
	}

	return result
}

func generate(name string, delay int, all bool) error {
	if delay < 0 {
		return fmt.Errorf("Invalid delay time")
	}

	output := &gif.GIF{}
	path := ".clip/%s"

	if all {
		result, err := exec.Command("git", "rev-list", "--all").Output()
		if err != nil {
			return err
		}

		for _, hash := range strings.split(string(result), "\n") {
			if !IsExists(fmt.Sprintf("./clip/%s", hash)) {
				ExportPicture("", hash) // TODO: ファイル名の保存
			}
		}
		// TODO: Implement a function to generate a picture from a old commit
		// ExportPicture()
	} // else {
	// 	hashes, err := PickValidCommits()
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	hashes, err := PickValidCommits()
	if err != nil {
		return err
	}

	for i, hash := range hashes {
		fmt.Printf("Generating... %d %%\r", int(float32(i)/float32(len(hashes))*100))

		f, err := os.OpenFile(fmt.Sprintf(path, hash), os.O_RDONLY, 0600)
		if err != nil {
			return err
		}

		buf := bytes.Buffer{}
		tmp, err := png.Decode(f)
		if err != nil {
			return err
		}

		gif.Encode(&buf, tmp, nil)

		input, err := gif.Decode(&buf)
		if err != nil {
			return err
		}

		f.Close()

		output.Image = append(output.Image, input.(*image.Paletted))
		output.Delay = append(output.Delay, delay/10)
	}

	fmt.Println("Generating... done!")

	f, _ := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0600)
	defer f.Close()

	gif.EncodeAll(f, output)

	return nil
}

// Gif generate gif from all pictures
func Gif(c *cli.Context) {
	err := generate(c.String("output"), c.Int("delay"), c.Bool("all"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
