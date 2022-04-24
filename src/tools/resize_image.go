package tools

import (
	"errors"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

type LocalResize struct{}

func (lr LocalResize) Resize(from, to string, height, width uint) error {
	img, err := lr.open(from)
	if err != nil {
		return err
	}
	resizeImage, _, _ := lr.doResize(img, width, height)
	return lr.save(resizeImage, to)
}

func (lr LocalResize) save(img image.Image, path string) error {
	if f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm); err == nil {
		defer f.Close()
		return jpeg.Encode(f, img, &(jpeg.Options{75}))
	} else {
		return err
	}
}

func (lr LocalResize) open(path string) (image.Image, error) {
	if f, err := os.Open(path); err == nil {
		defer f.Close()
		var img image.Image
		var err2 error
		ext := strings.ToLower(filepath.Ext(path))
		switch {
		case strings.EqualFold(ext, ".jpg") || strings.EqualFold(ext, ".jpeg"):
			img, err2 = jpeg.Decode(f)
			break
		case strings.EqualFold(ext, ".png"):
			img, err2 = png.Decode(f)
			break
		default:
			err2 = errors.New("unknown format")
		}
		if err2 == nil {
			return img, nil
		} else {
			return nil, err2
		}
	} else {
		return nil, err
	}
}

func (lr LocalResize) doResize(img image.Image, width, height uint) (image.Image, uint, uint) {
	x, y := float32(img.Bounds().Size().X), float32(img.Bounds().Size().Y)
	if float32(height) > y || float32(width) > x {
		return img, uint(x), uint(y)
	}
	switch {
	case width == 0 && height == 0:
		return img, uint(x), uint(y)
	case width == 0:
		width = uint((float32(height) / y) * x)
	case height == 0:
		height = uint((float32(width) / x) * y)
	}
	return resize.Resize(width, height, img, resize.Bicubic), width, height
}
