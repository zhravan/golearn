package image_processing_utility

import (
	"image"
	"image/color"
	"image/draw"
	"testing"
)

func createTestImage() image.Image {
	rect := image.Rect(0, 0, 1, 1)
	img := image.NewRGBA(rect)
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{255, 0, 0, 255}}, image.Point{}, draw.Src)
	return img
}

func TestGrayscale(t *testing.T) {
	img := createTestImage()
	grayImg := Grayscale(img)

	c := grayImg.At(0, 0)
	r, g, b, a := c.RGBA()
	if r != g || g != b || a != 0xFFFF {
		t.Errorf("Expected grayscale color, got %v", c)
	}
}

func TestInvert(t *testing.T) {
	img := createTestImage()
	invertedImg := Invert(img)

	c := invertedImg.At(0, 0)
	r, g, b, a := c.RGBA()
	// Original color was red (255,0,0,255) in RGBA terms (0xFFFF, 0x0000, 0x0000, 0xFFFF)
	// Inverted should be (0, 255, 255, 255) in RGBA terms (0x0000, 0xFFFF, 0xFFFF, 0xFFFF)

	// Since the image/color.RGBA.RGBA() returns values scaled to 0-0xFFFF, we check against that.
	if r != 0 || g != 0xFFFF || b != 0xFFFF || a != 0xFFFF {
		t.Errorf("Expected inverted color (0, 0xFFFF, 0xFFFF, 0xFFFF), got (%d, %d, %d, %d)", r, g, b, a)
	}
}

