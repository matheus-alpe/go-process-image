package processimage

import (
	"image"
	"image/jpeg"
	"os"

	"github.com/disintegration/imaging"
)

func Read(path string) image.Image {
	input, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	img, _, err := image.Decode(input)
	if err != nil {
		panic(err)
	}

	return img
}

func Resize(img image.Image) image.Image {
	return imaging.Resize(img, 250, 250, imaging.Lanczos)
}

func Write(path string, img image.Image) {
	output, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer output.Close()

	if err = jpeg.Encode(output, img, nil); err != nil {
		panic(err)
	}
}
