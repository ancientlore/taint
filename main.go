package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"log"
	"math"
	"os"

	"github.com/ancientlore/taint/images"
)

var (
	center   = flag.Bool("center", false, "Center images.")
	useColor = flag.Bool("color", false, "Output in color.")
	useOp    = flag.String("op", "avg", "Operation: avg, mul, min, max, and, or, xor, sq, ln")
)

func main() {
	flag.Parse()

	var (
		imgs   []image.Image
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
		imgs = append(imgs, img)
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
	log.Print(bounds)

	// center images
	if *center {
		for i := range imgs {
			b := imgs[i].Bounds()
			dx := (b.Dx() - bounds.Dx()) / 2
			dy := (b.Dy() - bounds.Dy()) / 2
			if dx > 0 || dy > 0 {
				r := bounds.Add(image.Pt(dx, dy))
				fmt.Printf("Image %d is bigger, adjust rect to %s\n", i, r)
				switch img := imgs[i].(type) {
				case *image.Alpha:
					imgs[i] = images.Offset{Image: img.SubImage(r)}
				case *image.CMYK:
					imgs[i] = images.Offset{Image: img.SubImage(r)}
				case *image.Alpha16:
					imgs[i] = images.Offset{Image: img.SubImage(r)}
				case *image.Gray:
					imgs[i] = images.Offset{Image: img.SubImage(r)}
				case *image.Gray16:
					imgs[i] = images.Offset{Image: img.SubImage(r)}
				case *image.NRGBA:
					imgs[i] = images.Offset{Image: img.SubImage(r)}
				case *image.NRGBA64:
					imgs[i] = images.Offset{Image: img.SubImage(r)}
				case *image.NYCbCrA:
					imgs[i] = images.Offset{Image: img.SubImage(r)}
				case *image.Paletted:
					imgs[i] = images.Offset{Image: img.SubImage(r)}
				case *image.RGBA:
					imgs[i] = images.Offset{Image: img.SubImage(r)}
				case *image.RGBA64:
					imgs[i] = images.Offset{Image: img.SubImage(r)}
				case *image.YCbCr:
					imgs[i] = images.Offset{Image: img.SubImage(r)}
				}
				// fmt.Printf("Image bounds are now %s\n", imgs[i].Bounds())
			}
		}
	}

	// create operation
	var src image.Image
	switch *useOp {
	case "avg":
		src = images.Average(imgs)
	case "mul":
		src = images.Multiply(imgs)
	case "min":
		src = images.Min(imgs)
	case "max":
		src = images.Max(imgs)
	case "and":
		src = images.And(imgs)
	case "or":
		src = images.Or(imgs)
	case "xor":
		src = images.Xor(imgs)
	case "sq":
		src = images.Square(imgs)
	case "ln":
		src = images.Ln(imgs)
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

	draw.Draw(finalImg, finalImg.Bounds(), src, src.Bounds().Min, draw.Src)

	out, err := os.Create("output.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	err = jpeg.Encode(out, finalImg, &jpeg.Options{Quality: 90})
	//err = png.Encode(out, finalImg)
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
