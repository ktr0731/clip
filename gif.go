package main

import (
	"flag"
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/png"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
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
	dir := ".clip"

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
		return 1
	}

	// Generate images which is not generated yet, but committed
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
			if isExists(filepath.Join(dir, hash)) {
				continue
			}

			if status := export.Run([]string{target, hash}); status != 0 {
				fmt.Fprintf(os.Stderr, "cannot export %s@%s\n", target, hash)
				return 1
			}
		}
	}

	hashes, err := pickValidCommits()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	out, _ := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0600)
	defer out.Close()

	generated, err := generate(dir, hashes, delay)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	gif.EncodeAll(out, generated)

	return 0
}

func generate(dir string, hashes []string, delay int) (*gif.GIF, error) {
	output := &gif.GIF{}
	output.Image = make([]*image.Paletted, len(hashes))
	output.Delay = make([]int, len(hashes))
	for i, hash := range hashes {
		fmt.Printf("Generating... %d %%\r", int(float32(i)/float32(len(hashes))*100))

		f, err := os.Open(filepath.Join(dir, hash))
		if err != nil {
			return nil, err
		}

		tmp, err := png.Decode(f)
		if err != nil {
			return nil, err
		}

		paletted := image.NewPaletted(tmp.Bounds(), palette.WebSafe)
		draw.FloydSteinberg.Draw(paletted, tmp.Bounds(), tmp, image.ZP)

		f.Close()

		output.Image[i] = paletted
		output.Delay[i] = delay / 10
	}

	fmt.Println("Generating... done!")

	return output, nil
}

func reverse(s []string) []string {
	result := make([]string, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		result[len(s)-i-1] = s[i]
	}

	return result
}
