package main

import (
	"gnr"
	"gnr/object"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

const (
	WIDTH  = 320
	HEIGHT = 240
)

func main() {
	scene := gnr.Scene{
		Camera: gnr.Camera{
			Position: gnr.Vector3f{
				X: 2,
				Y: 0,
				Z: -3,
			},
			PixelWidth:    WIDTH,
			PixelHeight:   HEIGHT,
			VirtualWidth:  4,
			VirtualHeight: 3,
			Angle:         60.0,
		},
		Objects: []gnr.Object{
			&object.Cube{
				CornerMin: gnr.Vector3f{
					X: -1,
					Y: -1,
					Z: -1,
				},
				CornerMax: gnr.Vector3f{
					X: 1,
					Y: 1,
					Z: 1,
				},
			},
		},
	}

	img := image.NewRGBA(image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{WIDTH, HEIGHT},
	})

	for x := uint64(0); x < WIDTH; x++ {
		for y := uint64(0); y < HEIGHT; y++ {
			c := scene.TracePixel(x, y)
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
