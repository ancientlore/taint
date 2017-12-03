package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"log"
	"math"
	"os"
)

var (
	useColor = flag.Bool("color", false, "Output in color.")
	useOp    = flag.String("op", "avg", "Operation: avg, mul, min, max, and, or, xor")
)

// Merger implements operations on pixels
type Merger interface {
	Reset()
	Add(c color.Color)
	Merge() color.Color
}

func main() {
	flag.Parse()

	var (
		images []image.Image
		bounds = image.Rectangle{
			Min: image.Point{
				X: math.MinInt32,
				Y: math.MinInt32,
			},
			Max: image.Point{
				X: math.MaxInt32,
				Y: math.MaxInt32,
			},
		}
	)

	// compute final image size
	for i := 0; i < flag.NArg(); i++ {
		img, err := loadImage(flag.Arg(i))
		if err != nil {
			log.Fatal(err)
		}
		images = append(images, img)
		b := img.Bounds()
		fmt.Println(flag.Arg(i), " ", b)
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
	// log.Print(bounds)

	// create operation
	var op Merger
	switch *useOp {
	case "avg":
		op = &average{}
	case "mul":
		op = &multiply{}
	case "min":
		op = &minimum{}
	case "max":
		op = &maximum{}
	case "and":
		op = &bitwiseAnd{}
	case "or":
		op = &bitwiseOr{}
	case "xor":
		op = &bitwiseXor{}
	default:
		log.Fatal("Invalid operation: ", *useOp)
	}

	// create final image
	var finalImg draw.Image
	if *useColor {
		finalImg = image.NewNRGBA(bounds)
	} else {
		finalImg = image.NewGray16(bounds)
	}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			var (
				c color.Color
			)
			op.Reset()
			for _, img := range images {
				op.Add(img.At(x, y))
			}
			c = op.Merge()
			finalImg.Set(x, y, c)
		}
	}

	out, err := os.Create("output.png")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	err = png.Encode(out, finalImg)
	if err != nil {
		log.Fatal(err)
	}
}

func loadImage(fn string) (image.Image, error) {
	reader, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	m, _, err := image.Decode(reader)
	return m, err
}
