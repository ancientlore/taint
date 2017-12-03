package main

import (
	"image/color"
	"math"
)

type multiply struct {
	r     uint64
	g     uint64
	b     uint64
	a     uint64
	count uint64
}

func (op *multiply) Reset() {
	op.r = 1
	op.g = 1
	op.b = 1
	op.a = 0
	op.count = 0
}

func (op *multiply) Add(c color.Color) {
	r, g, b, a := c.RGBA()
	op.r *= uint64(r)
	op.g *= uint64(g)
	op.b *= uint64(b)
	if uint64(a) > op.a {
		op.a = uint64(a)
	}
	op.count++
}

func (op *multiply) Merge() color.Color {
	var c color.NRGBA64
	c.R = uint16(math.Pow(float64(op.r), 1/float64(op.count)))
	c.G = uint16(math.Pow(float64(op.g), 1/float64(op.count)))
	c.B = uint16(math.Pow(float64(op.b), 1/float64(op.count)))
	c.A = uint16(op.a)
	return c
}
