package main

import "image/color"

type bitwiseAnd struct {
	r uint32
	g uint32
	b uint32
	a uint32
}

func (op *bitwiseAnd) Reset() {
	op.r = 0xFFFF
	op.g = 0xFFFF
	op.b = 0xFFFF
	op.a = 0
}

func (op *bitwiseAnd) Add(c color.Color) {
	r, g, b, a := c.RGBA()
	op.r &= r
	op.g &= g
	op.b &= b
	if a > op.a {
		op.a = a
	}
}

func (op *bitwiseAnd) Merge() color.Color {
	var c color.NRGBA64
	c.R = uint16(op.r)
	c.G = uint16(op.g)
	c.B = uint16(op.b)
	c.A = uint16(op.a)
	return c
}
