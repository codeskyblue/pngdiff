// Package pngdiff provides ...
// image.RGBA diff and patch
package pngdiff

import (
	"errors"
	"image"
	"image/png"
	"os"
)

var ErrSizeNotMatch = errors.New("Image size are different")

// Diff computes the difference bwtween old and new, and return the result.
func Diff(imold, imnew *image.RGBA) (patch *image.RGBA, err error) {
	if imold.Rect != imnew.Rect {
		return nil, ErrSizeNotMatch
	}
	patch = image.NewRGBA(imold.Rect)
	pixold, pixnew := imold.Pix, imnew.Pix
	//n := 0
	//tot := len(imold.Pix)/4

	for i := 0; i < len(imold.Pix)/4; i++ {
		if pixold[i*4] != pixnew[i*4] ||
			pixold[i*4+1] != pixnew[i*4+1] ||
			pixold[i*4+2] != pixnew[i*4+2] ||
			pixold[i*4+3] != pixnew[i*4+3] {
			copy(patch.Pix[i*4:i*4+4], pixnew[i*4:i*4+4])
			//		n++
		}
	}
	//fmt.Println(n, tot)
	return
}

// Patch applies patch to im, and writes the result to out
func Patch(im *image.RGBA, patch *image.RGBA) (out *image.RGBA, err error) {
	if im.Rect != patch.Rect {
		return nil, ErrSizeNotMatch
	}
	out = image.NewRGBA(im.Rect)
	for i := 0; i < len(im.Pix)/4; i++ {
		if patch.Pix[i*4+3] == 0 {
			copy(out.Pix[i*4:i*4+4], im.Pix[i*4:i*4+4])
		} else {
			copy(out.Pix[i*4:i*4+4], patch.Pix[i*4:i*4+4])
		}
	}
	return
}

// Read png file
func ReadFile(filename string) (im *image.RGBA, err error) {
	fd, err := os.Open(filename)
	if err != nil {
		return
	}
	img, err := png.Decode(fd)
	if err != nil {
		return
	}
	return img.(*image.RGBA), nil
}

// Write RGBA struct as png file
func WriteFile(filename string, im *image.RGBA) error {
	fd, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fd.Close()
	return png.Encode(fd, im)
}
