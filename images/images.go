package images

import (
	"image"
	"math"
)

type images []image.Image

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (im images) Bounds() image.Rectangle {
	if len(im) == 0 {
		return image.Rectangle{}
	}
	bounds := image.Rectangle{
		Min: image.Point{
			X: math.MinInt32,
			Y: math.MinInt32,
		},
		Max: image.Point{
			X: math.MaxInt32,
			Y: math.MaxInt32,
		},
	}
	for _, img := range im {
		b := img.Bounds()
		if b.Min.X > bounds.Min.X {
			bounds.Min.X = b.Min.X
		}
		if b.Min.Y > bounds.Min.Y {
			bounds.Min.Y = b.Min.Y
		}
		if b.Max.X < bounds.Max.X {
			bounds.Max.X = b.Max.X
		}
		if b.Max.Y < bounds.Max.Y {
			bounds.Max.Y = b.Max.Y
		}
	}
	return bounds
}
