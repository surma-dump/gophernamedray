package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"

	"github.com/surma-dump/gophernamedray/gnr"
	"github.com/surma-dump/gophernamedray/gnr/object"
)

const (
	Width  = 600
	Height = 600
)

func main() {
	scene := gnr.Scene{
		Camera: gnr.Camera{
			Position:      gnr.Vector3f{0, 0.5, -3},
			PixelWidth:    Width,
			PixelHeight:   Height,
			VirtualWidth:  1,
			VirtualHeight: 1,
			Angle:         120.0,
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
		Max: image.Point{Width * 2, Height * 2},
	})
	draw.Draw(img, img.Bounds(), &image.Uniform{color.Black}, image.Point{0, 0}, draw.Over)

	hitImg := gnr.SubImage(img, image.Rect(0, 0, Width, Height))
	distImg := gnr.SubImage(img, image.Rect(Width, 0, 2*Width, Height))
	normImg := gnr.SubImage(img, image.Rect(0, Height, Width, 2*Height))
	// colImg := SubImage(img, image.Rect(Width, Height, 2*Width, 2*Height))

	fFog := gnr.LerpCap(0, 10, 255, 0)
	fColor := gnr.LerpCap(0, 1, 0, 255)
	_, _ = fFog, fColor
	for x := uint64(0); x < Width; x++ {
		for y := uint64(0); y < Height; y++ {
			hit, matColor, distance, normal := scene.TracePixel(x, y)

			_, _, _, _ = hit, matColor, distance, normal

			// Hit image
			if hit {
				hitImg.Set(int(x), int(y), color.White)
			}

			// Distance image
			col := color.RGBA{
				R: uint8(fFog(distance)),
				G: uint8(fFog(distance)),
				B: uint8(fFog(distance)),
				A: 255,
			}
			if hit {
				distImg.Set(int(x), int(y), col)
			}

			// Normal image
			col = color.RGBA{
				R: uint8(fColor(normal.X)),
				G: uint8(fColor(normal.Y)),
				B: uint8(fColor(normal.Z)),
				A: 255,
			}
			if hit {
				normImg.Set(int(x), int(y), col)
			}
		}
	}

	f, e := os.Create("image.png")
	if e != nil {
		log.Fatalf("Could not open file for writing: %s\n", e)
	}
	defer f.Close()
	png.Encode(f, img)
}
