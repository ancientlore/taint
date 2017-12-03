package main

import "image/color"

type average struct {
	r     uint64
	g     uint64
	b     uint64
	a     uint64
	count uint64
}

func (op *average) Reset() {
	op.r = 0
	op.g = 0
	op.b = 0
	op.a = 0
	op.count = 0
}

func (op *average) Add(c color.Color) {
	r, g, b, a := c.RGBA()
	op.r += uint64(r)
	op.g += uint64(g)
	op.b += uint64(b)
	if uint64(a) > op.a {
		op.a = uint64(a)
	}
	op.count++
}

func (op *average) Merge() color.Color {
	var c color.NRGBA64
	c.R = uint16(op.r / op.count)
	c.G = uint16(op.g / op.count)
	c.B = uint16(op.b / op.count)
	c.A = uint16(op.a)
	return c
}
