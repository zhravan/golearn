package image_processing_utility

import (
	"image"
	"image/color"
	"image/draw"
)

// Grayscale converts an image to grayscale.
func Grayscale(img image.Image) *image.Gray {
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y

	grayImg := image.NewGray(bounds)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			oldColor := img.At(x, y)
			grayColor := color.GrayModel.Convert(oldColor).(color.Gray)
			grayImg.Set(x, y, grayColor)
		}
	}
	return grayImg
}

// Invert inverts the colors of an image.
func Invert(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	rgbaImg := image.NewRGBA(bounds)

	draw.Draw(rgbaImg, bounds, img, bounds.Min, draw.Src)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := rgbaImg.At(x, y).RGBA()
			r, g, b, a = 0xFFFF-r, 0xFFFF-g, 0xFFFF-b, a
			rgbaImg.SetRGBA(x, y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
		}
	}
	return rgbaImg
}

