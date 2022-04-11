package tools

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func openImage(path string) (image.Image, error) {
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

func resizeImage(img image.Image, width, height uint) (image.Image, uint, uint) {
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
	//return resizer.Resize(width, height, img, resizer.Bicubic), width, height
	// TODO
	return nil, 0, 0
}
