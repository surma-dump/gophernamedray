package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"runtime/pprof"

	"github.com/surma-dump/gophernamedray/gnr"
	"github.com/surma-dump/gophernamedray/gnr/object"
)

const (
	Width  = 600
	Height = 600

	AxisHeads = 0.2
)

func main() {
	scene := &gnr.Scene{
		Object: &gnr.FlatShader{
			Object: gnr.NewLayers(
				gnr.NewUnion(
					&gnr.ColorChanger{
						Object: &gnr.Union{
							Objects: []gnr.Object{
								&object.Triangle{
									Points: [3]*gnr.Vector3f{
										{0, -AxisHeads, 0},
										{1, 0, 0},
										{0, AxisHeads, 0},
									},
								},
								&object.Triangle{
									Points: [3]*gnr.Vector3f{
										{0, 0, -AxisHeads},
										{1, 0, 0},
										{0, 0, AxisHeads},
									},
								},
							},
						},
						NewColor: gnr.ColorRed,
					},
					&gnr.ColorChanger{
						Object: gnr.NewUnion(
							&object.Triangle{
								Points: [3]*gnr.Vector3f{
									{-AxisHeads, 0, 0},
									{AxisHeads, 0, 0},
									{0, 1, 0},
								},
							},
							&object.Triangle{
								Points: [3]*gnr.Vector3f{
									{0, 0, -AxisHeads},
									{0, 0, AxisHeads},
									{0, 1, 0},
								},
							},
						),
						NewColor: gnr.ColorGreen,
					},
					&gnr.ColorChanger{
						Object: gnr.NewUnion(
							&object.Triangle{
								Points: [3]*gnr.Vector3f{
									{-AxisHeads, 0, 0},
									{AxisHeads, 0, 0},
									{0, 0, 1},
								},
							},
							&object.Triangle{
								Points: [3]*gnr.Vector3f{
									{0, -AxisHeads, 0},
									{0, 0, 1},
									{0, AxisHeads, 0},
								},
							},
						),
						NewColor: gnr.ColorBlue,
					},
				), // Union
				gnr.NewUnion(
					&gnr.XZChecker{
						Object: &object.Plane{
							Normal:   &gnr.Vector3f{0, 1, 0},
							Distance: 0,
						},
						ColorA: gnr.ColorWhite,
						ColorB: gnr.ColorBlack,
					},
					&gnr.ColorChanger{
						Object: &object.Triangle{
							Points: [3]*gnr.Vector3f{
								{-2, 2, 0},
								{0, 2, 0},
								{-1, 4, 0},
							},
						},
						NewColor: gnr.ColorYellow,
					},
					&gnr.ColorChanger{
						Object: &object.Triangle{
							Points: [3]*gnr.Vector3f{
								{2, 2, 0},
								{0, 2, 0},
								{1, 4, 0},
							},
						},
						NewColor: gnr.ColorCyan,
					},
					&gnr.ColorChanger{
						Object: &object.Sphere{
							Center: &gnr.Vector3f{1, 1, 0},
							Radius: 1,
						},
						NewColor: gnr.ColorBlue,
					},
					gnr.NewIntersection(
						&gnr.ColorChanger{
							Object: &object.AxisAlignedBox{
								Min: &gnr.Vector3f{-1, 0.5, -0.5},
								Max: &gnr.Vector3f{0, 1.5, 0.5},
							},
							NewColor: gnr.ColorRed,
						},
						&gnr.ColorChanger{
							Object: &object.Sphere{
								Center: &gnr.Vector3f{-0.5, 1, 0},
								Radius: 0.7,
							},
							NewColor: gnr.ColorGreen,
						},
					),
					&gnr.Difference{
						Minuend: &gnr.ColorChanger{
							Object: &object.AxisAlignedBox{
								Min: &gnr.Vector3f{-2.5, 0.5, -0.5},
								Max: &gnr.Vector3f{-1.5, 1.5, 0.5},
							},
							NewColor: gnr.ColorMagenta,
						},
						Subtrahend: &gnr.ColorChanger{
							Object: &object.Sphere{
								Center: &gnr.Vector3f{-1.5, 1.5, -0.5},
								Radius: 0.5,
							},
							NewColor: gnr.ColorGreen,
						},
					},
				), // Union
			), // Layers
			FalloffFunc: gnr.NewLinearFalloffFunc(&gnr.Vector3f{-3, -5, 1}),
		}, // FlatShader
	}

	cameras := []*gnr.Camera{
		{
			Position:      &gnr.Vector3f{0, 1, -5},
			ViewDirection: &gnr.Vector3f{0, 0, 1},
			UpDirection:   &gnr.Vector3f{0, 1, 0},
			PixelWidth:    Width,
			PixelHeight:   Height,
			VirtualWidth:  1,
			VirtualHeight: 1,
			Angle:         60.0,
		},
		{
			Position:      &gnr.Vector3f{0, 1, 5},
			ViewDirection: &gnr.Vector3f{0, 0, -1},
			UpDirection:   &gnr.Vector3f{0, 1, 0},
			PixelWidth:    Width,
			PixelHeight:   Height,
			VirtualWidth:  1,
			VirtualHeight: 1,
			Angle:         60.0,
		},
		{
			Position:      &gnr.Vector3f{3, 3, -3},
			ViewDirection: &gnr.Vector3f{-3, -3, 3},
			UpDirection:   &gnr.Vector3f{0, 1, 0},
			PixelWidth:    Width,
			PixelHeight:   Height,
			VirtualWidth:  1,
			VirtualHeight: 1,
			Angle:         60.0,
		},
	}
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	for idx, camera := range cameras {
		scene.Camera = camera
		scene.Camera.Normalize()
		img := renderImage(scene)
		f, e := os.Create(fmt.Sprintf("image_%03d.png", idx))
		if e != nil {
			log.Fatalf("Could not open file for writing: %s\n", e)
		}
		defer f.Close()
		png.Encode(f, img)
	}
}

func renderImage(scene *gnr.Scene) image.Image {
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

			if hit {
				colImg.Set(int(x), int(y), ir.Color.ToColor())
			}
		}
	}
	return img
}
