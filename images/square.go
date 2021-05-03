package images

import (
	"image"
	"image/color"
	"math"
)

// Square is a list of images that are combined using a sum of squares.
type Square []image.Image

// ColorModel returns the Image's color model.
func (im Square) ColorModel() color.Model {
	return color.NRGBA64Model
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (im Square) Bounds() image.Rectangle {
	return images(im).Bounds()
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (im Square) At(x, y int) color.Color {
	var (
		r uint64
		g uint64
		b uint64
		a uint64
	)

	for _, img := range im {
		c := img.At(x, y)
		ir, ig, ib, ia := c.RGBA()
		r += uint64(ir) * uint64(ir)
		g += uint64(ig) * uint64(ig)
		b += uint64(ib) * uint64(ib)
		if uint64(ia) > a {
			a = uint64(ia)
		}
	}

	var c color.NRGBA64
	count := uint64(len(im))
	factor := math.Sqrt(float64(math.MaxUint16) * float64(math.MaxUint16) * float64(count))
	c.R = uint16(float64(math.MaxUint16) * math.Sqrt(float64(r)) / factor)
	c.G = uint16(float64(math.MaxUint16) * math.Sqrt(float64(g)) / factor)
	c.B = uint16(float64(math.MaxUint16) * math.Sqrt(float64(b)) / factor)
	c.A = uint16(a)

	return c
}
