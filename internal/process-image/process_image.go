package processimage

import (
	"image"
	"image/color"
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

func Resize(img image.Image, width, height int) image.Image {
	return imaging.Resize(img, width, height, imaging.Lanczos)
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

func Grayscale(img image.Image) image.Image {
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := img.At(x, y)
			grayPixel := color.GrayModel.Convert(pixel)
			grayImg.Set(x, y, grayPixel)
		}
	}

	return grayImg
}
