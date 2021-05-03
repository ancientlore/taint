package images

import (
	"image"
	"image/color"
)

// Max is a list of images combined by taking the maximum value.
type Max []image.Image

// ColorModel returns the Image's color model.
func (im Max) ColorModel() color.Model {
	return color.RGBA64Model
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (im Max) Bounds() image.Rectangle {
	return images(im).Bounds()
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (im Max) At(x, y int) color.Color {
	var (
		r uint32
		g uint32
		b uint32
		a uint32
	)

	for _, img := range im {
		c := img.At(x, y)
		ir, ig, ib, ia := c.RGBA()
		if ir > r {
			r = ir
		}
		if ig > g {
			g = ig
		}
		if ib > b {
			b = ib
		}
		if ia > a {
			a = ia
		}
	}

	var c color.RGBA64
	c.R = uint16(r)
	c.G = uint16(g)
	c.B = uint16(b)
	c.A = uint16(a)

	return c
}
