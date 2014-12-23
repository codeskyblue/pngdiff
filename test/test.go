package main

import (
	"log"

	"github.com/codeskyblue/pngdiff"
)

func main() {
	im1, _ := pngdiff.ReadFile("im1.png")
	im2, _ := pngdiff.ReadFile("im2.png")
	patch, err := pngdiff.Diff(im1, im2)
	if err != nil {
		log.Fatal(err)
	}
	pngdiff.WriteFile("patch.png", patch)
	out, err := pngdiff.Patch(im1, patch)
	if err != nil {
		log.Fatal(err)
	}
	pngdiff.WriteFile("out.png", out)
}
