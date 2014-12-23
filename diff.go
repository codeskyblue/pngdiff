// Package main provides ...
package pngdiff

import (
	"errors"
	"fmt"
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
	n := 0
	tot := 0

	for i := 0; i < len(imold.Pix)/4; i++ {
		if pixold[i*4] != pixnew[i*4] ||
			pixold[i*4+1] != pixnew[i*4+1] ||
			pixold[i*4+2] != pixnew[i*4+2] ||
			pixold[i*4+3] != pixnew[i*4+3] {
			copy(patch.Pix[i*4:i*4+4], pixnew[i*4:i*4+4])
			n++
		}
		tot++
	}
	fmt.Println(n, tot)
	return
}

func Patch(im *image.RGBA, patch *image.RGBA) (out *image.RGBA, err error) {
	if im.Rect != patch.Rect {
		return nil, ErrSizeNotMatch
	}
	out = image.NewRGBA(im.Rect)
	n := 0
	tot := 0
	for i := 0; i < len(im.Pix)/4; i++ {
		tot++
		if patch.Pix[i*4+3] == 0 {
			copy(out.Pix[i*4:i*4+4], im.Pix[i*4:i*4+4])
		} else {
			n++
			copy(out.Pix[i*4:i*4+4], patch.Pix[i*4:i*4+4])
		}
	}
	fmt.Println(n, tot)
	return
}

/*
func rawRead(filename string) (im *image.RGBA, err error) {
	var width, height, format int32
	bf, err := os.Open(filename)
	if err != nil {
		return
	}
	binary.Read(bf, binary.LittleEndian, &width)
	binary.Read(bf, binary.LittleEndian, &height)
	binary.Read(bf, binary.LittleEndian, &format)
	fmt.Println(width, height, format)
	im = image.NewRGBA(image.Rectangle{image.ZP, image.Point{int(width), int(height)}})
	_, err = bf.Read(im.Pix)
	return
}
*/

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

func WriteFile(filename string, im *image.RGBA) error {
	fd, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fd.Close()
	return png.Encode(fd, im)
}
