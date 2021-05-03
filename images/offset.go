package images

import (
	"image"
	"image/color"
)

// Offset resets an offset image to zeroed.
type Offset struct {
	image.Image
}

// ColorModel returns the Image's color model.
func (im Offset) ColorModel() color.Model {
	return im.Image.ColorModel()
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (im Offset) Bounds() image.Rectangle {
	return im.Image.Bounds().Sub(image.Point{X: im.Image.Bounds().Min.X, Y: im.Image.Bounds().Min.Y})
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (im Offset) At(x, y int) color.Color {
	return im.Image.At(x+im.Image.Bounds().Min.X, y+im.Image.Bounds().Min.Y)
}
