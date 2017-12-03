package main

import "image/color"

type bitwiseXor struct {
	r uint32
	g uint32
	b uint32
	a uint32
}

func (op *bitwiseXor) Reset() {
	op.r = 0
	op.g = 0
	op.b = 0
	op.a = 0
}

func (op *bitwiseXor) Add(c color.Color) {
	r, g, b, a := c.RGBA()
	op.r ^= r
	op.g ^= g
	op.b ^= b
	if a > op.a {
		op.a = a
	}
}

func (op *bitwiseXor) Merge() color.Color {
	var c color.NRGBA64
	c.R = uint16(op.r)
	c.G = uint16(op.g)
	c.B = uint16(op.b)
	c.A = uint16(op.a)
	return c
}
