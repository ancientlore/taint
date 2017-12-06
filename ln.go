package main

import (
	"image/color"
	"math"
)

type ln struct {
	r     uint64
	g     uint64
	b     uint64
	a     uint64
	count uint64
}

func (op *ln) Reset() {
	op.r = 0
	op.g = 0
	op.b = 0
	op.a = 0
	op.count = 0
}

func (op *ln) Add(c color.Color) {
	r, g, b, a := c.RGBA()
	op.r += uint64(r)
	op.g += uint64(g)
	op.b += uint64(b)
	if uint64(a) > op.a {
		op.a = uint64(a)
	}
	op.count++
}

func (op *ln) Merge() color.Color {
	var c color.NRGBA64
	factor := math.Log(float64(op.count)*float64(math.MaxUint16) + 1.0)
	c.R = uint16(math.Log(float64(op.r)+1) * math.MaxUint16 / factor)
	c.G = uint16(math.Log(float64(op.g)+1) * math.MaxUint16 / factor)
	c.B = uint16(math.Log(float64(op.b)+1) * math.MaxUint16 / factor)
	c.A = uint16(op.a)
	return c
}
