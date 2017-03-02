package main

import (
	"image/gif"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestReverse(t *testing.T) {
	actual := []string{"Makise", "Kurisu", "ha", "kawaii"}

	expect := reverse(actual)
	if actual[0] != expect[len(actual)-1] {
		t.Errorf("Expect: %s, Actual: %s", expect[len(actual)-1], actual[0])
	}
}

func BenchmarkGenerate(b *testing.B) {
	var generated *gif.GIF
	var err error

	hashes := []string{
		"98c3b2b8caeadff1b02bf872a52362f6e6f028ce",
		"d4e8dcf0288b76745e3d334901dc985522d91c0e",
		"48dec9f7d265b4b510abe10d3f421d7dd90d6850",
		"9bcf34cc261553fd1cbb897e23c3f78b3a4c9845",
		"7681979469c8a907f2dd3396fc9f2581c33ed461",
	}

	out, err := ioutil.TempFile("", "clip")
	defer func() {
		out.Close()
		if err := os.Remove(out.Name()); err != nil {
			b.Error(err)
		}
	}()

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		generated, err = generate(filepath.Join("tests", "assets", "gif"), hashes, 1000)
		if err != nil {
			b.Error(err)
		}
		gif.EncodeAll(out, generated)
	}
}
