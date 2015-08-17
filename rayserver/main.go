package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/surma-dump/gophernamedray/gnr"
	"github.com/surma-dump/gophernamedray/gnr/object"
)

const (
	Width  = 1024
	Height = 1024
)

func main() {
	scene := gnr.Scene{
		Camera: gnr.Camera{
			Position:      gnr.Vector3f{0, 0.5, -3},
			PixelWidth:    Width,
			PixelHeight:   Height,
			VirtualWidth:  1,
			VirtualHeight: 1,
			Angle:         90.0,
		},
		Objects: []gnr.Object{
			object.Plane{
				Normal:   gnr.Vector3f{0, 1, 0},
				Distance: 0,
			},
			object.Triangle{
				Points: [3]gnr.Vector3f{
					gnr.Vector3f{-0.5, 3, 1},
					gnr.Vector3f{0.5, 3, 1},
					gnr.Vector3f{0, 4, 1},
				},
			},
			object.Triangle{
				Points: [3]gnr.Vector3f{
					gnr.Vector3f{-1.5, 3, 1},
					gnr.Vector3f{-0.5, 3, 1},
					gnr.Vector3f{-1, 4, 1},
				},
			},
			object.Triangle{
				Points: [3]gnr.Vector3f{
					gnr.Vector3f{0.5, 3, 1},
					gnr.Vector3f{1.5, 3, 1},
					gnr.Vector3f{1, 4, 1},
				},
			},
			object.Triangle{
				Points: [3]gnr.Vector3f{
					gnr.Vector3f{-0.5, 4, 1},
					gnr.Vector3f{0.5, 4, 1},
					gnr.Vector3f{0, 5, 1},
				},
			},
			object.Triangle{
				Points: [3]gnr.Vector3f{
					gnr.Vector3f{-0.5, 0, 1},
					gnr.Vector3f{0.5, 0, 1},
					gnr.Vector3f{0, 1, 1},
				},
			},
		},
	}

	img := image.NewRGBA(image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{Width, Height},
	})

	fog := gnr.LerpCap(0, 10, 255, 0)
	for x := uint64(0); x < Width; x++ {
		for y := uint64(0); y < Height; y++ {
			_, d := scene.TracePixel(x, y)
			grey := uint8(fog(math.Floor(d)))
			col := color.RGBA{
				R: grey,
				G: grey,
				B: grey,
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
