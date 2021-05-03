package images

import (
	"image"
	"image/color"
)

// Or is a list of images combined with the OR operator.
type Or []image.Image

// ColorModel returns the Image's color model.
func (im Or) ColorModel() color.Model {
	return color.NRGBA64Model
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (im Or) Bounds() image.Rectangle {
	return images(im).Bounds()
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (im Or) At(x, y int) color.Color {
	var (
		r uint32
		g uint32
		b uint32
		a uint32
	)

	for _, img := range im {
		c := img.At(x, y)
		ir, ig, ib, ia := c.RGBA()

		r |= ir
		g |= ig
		b |= ib
		if ia > a {
			a = ia
		}
	}

	var c color.NRGBA64
	c.R = uint16(r)
	c.G = uint16(g)
	c.B = uint16(b)
	c.A = uint16(a)

	return c
}
