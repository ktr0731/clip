package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/png"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// GifCommand generates the gif of the production process from all pictures
type GifCommand struct{}

func (c *GifCommand) Synopsis() string {
	return "Generate Gif of the production process"
}

func (c *GifCommand) Help() string {
	return "Usage: clip gif"
}

func (c *GifCommand) Run(args []string) int {
	var name string
	var delay int
	var all bool
	flags := flag.NewFlagSet("gif", flag.ContinueOnError)
	flags.StringVar(&name, "output", "process.gif", "Output file name")
	flags.IntVar(&delay, "delay", 1000, "Delay time (ms)")
	flags.BoolVar(&all, "all", false, "Create pictures if there is no picture corresponding to commits")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	if delay < 0 {
		fmt.Fprintln(os.Stderr, "Invalid delay time")
	}

	output := &gif.GIF{}
	path := ".clip/%s"

	if all {
		_target, err := ioutil.ReadFile(".clipconfig")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}

		target := strings.TrimSpace(string(_target))

		result, err := exec.Command("git", "rev-list", "--all").Output()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}

		export := &ExportCommand{}
		for _, hash := range strings.Split(string(result), "\n") {
			if !isExists(fmt.Sprintf(path, hash)) {
				if status := export.Run([]string{target, hash}); status != 0 {
					fmt.Fprintln(os.Stderr, err)
					return 1
				}
			}
		}
	}

	hashes, err := pickValidCommits()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	for i, hash := range hashes {
		fmt.Printf("Generating... %d %%\r", int(float32(i)/float32(len(hashes))*100))

		f, err := os.OpenFile(fmt.Sprintf(path, hash), os.O_RDONLY, 0600)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}

		buf := bytes.Buffer{}
		tmp, err := png.Decode(f)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}

		gif.Encode(&buf, tmp, nil)

		input, err := gif.Decode(&buf)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}

		// TODO: 関数に切り出してdefer
		f.Close()

		output.Image = append(output.Image, input.(*image.Paletted))
		output.Delay = append(output.Delay, delay/10)
	}

	fmt.Println("Generating... done!")

	f, _ := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0600)
	defer f.Close()

	gif.EncodeAll(f, output)

	return 0
}

func reverse(s []string) []string {
	result := make([]string, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		result[len(s)-i-1] = s[i]
	}

	return result
}
