package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/surma-dump/gophernamedray/gnr"
	"github.com/surma-dump/gophernamedray/gnr/object"
)

const (
	Width  = 320
	Height = 240
)

func main() {
	scene := gnr.Scene{
		Camera: gnr.Camera{
			Position: gnr.Vector3f{
				X: 2,
				Y: 1,
				Z: -3,
			},
			PixelWidth:    Width,
			PixelHeight:   Height,
			VirtualWidth:  4,
			VirtualHeight: 3,
			Angle:         60.0,
		},
		Objects: []gnr.Object{
			object.Plane{
				Normal: gnr.Vector3f{
					X: 0,
					Y: 1,
					Z: 0,
				},
				Distance: 0,
			},
		},
	}

	img := image.NewRGBA(image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{Width, Height},
	})

	for x := uint64(0); x < Width; x++ {
		for y := uint64(0); y < Height; y++ {
			c, _ := scene.TracePixel(x, y)
			col := color.RGBA{
				R: uint8(255 * c.R),
				G: uint8(255 * c.G),
				B: uint8(255 * c.B),
				A: 255,
			}
			img.Set(int(x), int(y), col)
		}
	}

	f, e := os.Create("image.png")
	if e != nil {
		log.Fatalf("Could not open file for writing: %s\n", e)
	}
	defer f.Close()
	png.Encode(f, img)
}
