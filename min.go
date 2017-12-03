package main

import (
	"image/color"
	"math"
)

type minimum struct {
	r uint32
	g uint32
	b uint32
	a uint32
}

func (op *minimum) Reset() {
	op.r = math.MaxUint32
	op.g = math.MaxUint32
	op.b = math.MaxUint32
	op.a = 0
}

func (op *minimum) Add(c color.Color) {
	r, g, b, a := c.RGBA()
	if r < op.r {
		op.r = r
	}
	if g < op.g {
		op.g = g
	}
	if b < op.b {
		op.b = b
	}
	if a > op.a {
		op.a = a
	}
}

func (op *minimum) Merge() color.Color {
	var c color.RGBA64
	c.R = uint16(op.r)
	c.G = uint16(op.g)
	c.B = uint16(op.b)
	c.A = uint16(op.a)
	return c
}
