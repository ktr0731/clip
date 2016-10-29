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
	result, err := exec.Command("git", "rev-list", "--all").Output()
	if err != nil {
		return err
	}

	if delay < 0 {
		return fmt.Errorf("Invalid delay time")
	}

	output := &gif.GIF{}
	path := ".clip/%s"

	tmp := strings.Split(string(result), "\n")
	tmp = tmp[:len(tmp)-1]

	hashes := reverse(tmp)

	for _, hash := range hashes {
		f, err := os.OpenFile(fmt.Sprintf(path, hash), os.O_RDONLY, 0600)
		if err != nil {
			if all {
				// return err
			} else {
				continue
			}
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

	f, _ := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0600)
	defer f.Close()

	gif.EncodeAll(f, output)

	return nil
}

// Gif generate gif from all pictures
func Gif(c *cli.Context) {
	err := generate(c.String("output"), c.Int("delay"), c.Bool("all"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
