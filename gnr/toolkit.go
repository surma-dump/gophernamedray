package gnr

import (
	"image"
	"image/color"
	"image/draw"
	"math"
)

// Lerp returns a function that performs a lerp on the given parameters. Lerp
// does not check boundaries on the input value.
func Lerp(inMin, inMax, outMin, outMax float64) func(d float64) float64 {
	return func(d float64) float64 {
		return (d-inMin)/(inMax-inMin)*(outMax-outMin) + outMin
	}
}

// LerpCap is the same as Lerp but caps the input to be in between
// inMin and inMax.
func LerpCap(inMin, inMax, outMin, outMax float64) func(d float64) float64 {
	f := Lerp(inMin, inMax, outMin, outMax)
	return func(d float64) float64 {
		return f(math.Max(math.Min(d, inMax), inMin))
	}
}

// SubImage extracts a rectangular subset of draw.Image as a separate draw.Image
// with a translated origin to the rectangles canonicalized Min point.
func SubImage(img draw.Image, r image.Rectangle) draw.Image {
	return &subimage{
		Image:  img,
		bounds: img.Bounds().Intersect(r.Canon()),
	}
}

type subimage struct {
	draw.Image
	bounds image.Rectangle
}

func (si *subimage) Bounds() image.Rectangle {
	return si.bounds
}

func (si *subimage) At(x, y int) color.Color {
	p := image.Point{x + si.bounds.Min.X, y + si.bounds.Min.Y}
	if !p.In(si.Bounds()) {
		return color.Black
	}
	return si.Image.At(p.X, p.Y)
}

func (si *subimage) Set(x, y int, c color.Color) {
	p := image.Point{x + si.bounds.Min.X, y + si.bounds.Min.Y}
	if !p.In(si.Bounds()) {
		return
	}
	si.Image.Set(p.X, p.Y, c)
}

type ColorChanger struct {
	Object
	NewColor Color
}

func (cc ColorChanger) RayInteraction(r Ray) (InteractionResult, bool) {
	ir, ok := cc.Object.RayInteraction(r)
	ir.Color = cc.NewColor
	return ir, ok
}

type XZChecker struct {
	Object
}

func (cc XZChecker) RayInteraction(r Ray) (InteractionResult, bool) {
	ir, ok := cc.Object.RayInteraction(r)
	x := int(math.Floor(ir.PointOfImpact.X))
	z := int(math.Floor(ir.PointOfImpact.Z))
	if (x+z)%2 == 0 {
		ir.Color = ColorBlack
	} else {
		ir.Color = ColorWhite
	}
	return ir, ok
}
