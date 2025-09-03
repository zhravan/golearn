package image_processing_utility

import (
    "image"
)

// TODO:
// - Implement basic image transforms:
//   - Grayscale: convert each pixel to grayscale using color.GrayModel.
//   - Invert: invert RGBA channels for each pixel.
// - Keep function names and return types; tests inspect pixel values.

// Grayscale converts an image to grayscale.
func Grayscale(img image.Image) *image.Gray {
    // TODO: convert img to grayscale and return *image.Gray
    return nil
}

// Invert inverts the colors of an image.
func Invert(img image.Image) *image.RGBA {
    // TODO: invert colors and return *image.RGBA
    return nil
}
