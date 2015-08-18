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
			Position:      gnr.Vector3f{0, 0.5, -1},
			PixelWidth:    Width,
			PixelHeight:   Height,
			VirtualWidth:  1,
			VirtualHeight: 1,
			Angle:         90.0,
		},
		Objects: []gnr.Object{
			gnr.XZChecker{
				Object: object.Plane{
					Normal:   gnr.Vector3f{0, 1, 0},
					Distance: 0,
				},
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
					gnr.Vector3f{-0.5, 1, 1},
					gnr.Vector3f{0.5, 1, 1},
					gnr.Vector3f{0, 2, 1},
				},
			},
			gnr.ColorChanger{
				Object: object.Sphere{
					Center: gnr.Vector3f{1, 1, 2},
					Radius: 1,
				},
				NewColor: gnr.ColorRed,
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
	colImg := gnr.SubImage(img, image.Rect(Width, Height, 2*Width, 2*Height))

	fFog := gnr.LerpCap(0, 15, 255, 0)
	fNormal := gnr.LerpCap(-1, 1, 0, 255)
	fColor := gnr.LerpCap(0, 1, 0, 255)
	for x := uint64(0); x < Width; x++ {
		for y := uint64(0); y < Height; y++ {
			ir, hit := scene.TracePixel(x, y)

			// Hit image
			if hit {
				hitImg.Set(int(x), int(y), color.White)
			}

			// Distance image
			col := color.RGBA{
				R: uint8(fFog(ir.Distance)),
				G: uint8(fFog(ir.Distance)),
				B: uint8(fFog(ir.Distance)),
				A: 255,
			}
			if hit {
				distImg.Set(int(x), int(y), col)
			}

			// Normal image
			col = color.RGBA{
				R: uint8(fNormal(ir.Normal.X)),
				G: uint8(fNormal(ir.Normal.Y)),
				B: uint8(fNormal(ir.Normal.Z)),
				A: 255,
			}
			if hit {
				normImg.Set(int(x), int(y), col)
			}

			// Color image
			col = color.RGBA{
				R: uint8(fColor(ir.Color.R)),
				G: uint8(fColor(ir.Color.G)),
				B: uint8(fColor(ir.Color.B)),
				A: 255,
			}
			if hit {
				colImg.Set(int(x), int(y), col)
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
