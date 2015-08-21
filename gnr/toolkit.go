package gnr

import (
	"image"
	"image/color"
	"image/draw"
	"log"
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

func (cc ColorChanger) RayInteraction(r Ray) ([]InteractionResult, bool) {
	irs, ok := cc.Object.RayInteraction(r)
	for i := range irs {
		irs[i].Color = cc.NewColor
	}
	return irs, ok
}

type XZChecker struct {
	Object
	ColorA, ColorB Color
}

func (cc XZChecker) RayInteraction(r Ray) ([]InteractionResult, bool) {
	irs, ok := cc.Object.RayInteraction(r)
	for i := range irs {
		x := int(math.Floor(irs[i].PointOfImpact.X))
		z := int(math.Floor(irs[i].PointOfImpact.Z))
		if (x+z)%2 == 0 {
			irs[i].Color = cc.ColorA
		} else {
			irs[i].Color = cc.ColorB
		}
	}
	return irs, ok
}

type Disable struct {
	Object
}

func (d Disable) RayInteraction(r Ray) ([]InteractionResult, bool) {
	return []InteractionResult{}, false
}

func (d Disable) Contains(p Vector3f) bool {
	return false
}

type Layers struct {
	Objects []Object
}

func NewLayers(o ...Object) Layers {
	return Layers{o}
}

func (l Layers) RayInteraction(r Ray) ([]InteractionResult, bool) {
	for _, o := range l.Objects {
		irs, ok := o.RayInteraction(r)
		if ok {
			return irs, ok
		}
	}
	return []InteractionResult{}, false
}

func (l Layers) Contains(p Vector3f) bool {
	return ObjectSlice(l.Objects).Any(func(o Object) bool {
		return o.Contains(p)
	})
}

type Logger struct {
	Object
	Prefix string
}

func NewLogger(prefix string, o Object) Logger {
	return Logger{o, prefix}
}

func (l Logger) RayInteraction(r Ray) ([]InteractionResult, bool) {
	irs, ok := l.Object.RayInteraction(r)
	log.Printf("%s.RayInteraction(%s) = ([%d]InteractionResult, %#v)", l.Prefix, r, len(irs), ok)
	return irs, ok
}
