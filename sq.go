package main

import (
	"image/color"
	"math"
)

type sq struct {
	r     uint64
	g     uint64
	b     uint64
	a     uint64
	count uint64
}

func (op *sq) Reset() {
	op.r = 0
	op.g = 0
	op.b = 0
	op.a = 0
	op.count = 0
}

func (op *sq) Add(c color.Color) {
	r, g, b, a := c.RGBA()
	op.r += uint64(r) * uint64(r)
	op.g += uint64(g) * uint64(g)
	op.b += uint64(b) * uint64(b)
	if uint64(a) > op.a {
		op.a = uint64(a)
	}
	op.count++
}

func (op *sq) Merge() color.Color {
	var c color.NRGBA64
	factor := math.Sqrt(float64(math.MaxUint16) * float64(math.MaxUint16) * float64(op.count))
	c.R = uint16(float64(math.MaxUint16) * math.Sqrt(float64(op.r)) / factor)
	c.G = uint16(float64(math.MaxUint16) * math.Sqrt(float64(op.g)) / factor)
	c.B = uint16(float64(math.MaxUint16) * math.Sqrt(float64(op.b)) / factor)
	c.A = uint16(op.a)
	return c
}
