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

// VLerp is the same as Lerp but for Vector3f.
func VLerp(inMin, inMax float64, outMin, outMax *Vector3f) func(d float64) *Vector3f {
	xLerp := Lerp(inMin, inMax, outMin.X, outMax.X)
	yLerp := Lerp(inMin, inMax, outMin.Y, outMax.Y)
	zLerp := Lerp(inMin, inMax, outMin.Z, outMax.Z)
	return func(d float64) *Vector3f {
		return &Vector3f{xLerp(d), yLerp(d), zLerp(d)}
	}
}

// VLerpCap is the same as VLerp but caps the input to be in between
// inMin and inMax.
func VLerpCap(inMin, inMax float64, outMin, outMax *Vector3f) func(d float64) *Vector3f {
	xLerp := LerpCap(inMin, inMax, outMin.X, outMax.X)
	yLerp := LerpCap(inMin, inMax, outMin.Y, outMax.Y)
	zLerp := LerpCap(inMin, inMax, outMin.Z, outMax.Z)
	return func(d float64) *Vector3f {
		return &Vector3f{xLerp(d), yLerp(d), zLerp(d)}
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
	NewColor *Vector3f
}

func (cc *ColorChanger) RayInteraction(r *Ray) []*InteractionResult {
	irs := cc.Object.RayInteraction(r)
	for i := range irs {
		irs[i].Color = cc.NewColor
	}
	return irs
}

type XZChecker struct {
	Object
	ColorA, ColorB *Vector3f
}

func (cc *XZChecker) RayInteraction(r *Ray) []*InteractionResult {
	irs := cc.Object.RayInteraction(r)
	for i := range irs {
		x := int(math.Floor(irs[i].PointOfImpact.X))
		z := int(math.Floor(irs[i].PointOfImpact.Z))
		if (x+z)%2 == 0 {
			irs[i].Color = cc.ColorA
		} else {
			irs[i].Color = cc.ColorB
		}
	}
	return irs
}

type Disable struct {
	Object
}

func (d *Disable) RayInteraction(r *Ray) []*InteractionResult {
	return []*InteractionResult{}
}

func (d *Disable) Contains(p *Vector3f) bool {
	return false
}

type Layers struct {
	Objects []Object
}

func NewLayers(o ...Object) *Layers {
	return &Layers{o}
}

func (l *Layers) RayInteraction(r *Ray) []*InteractionResult {
	for _, o := range l.Objects {
		irs := o.RayInteraction(r)
		if len(irs) > 0 {
			return irs
		}
	}
	return []*InteractionResult{}
}

func (l *Layers) Contains(p *Vector3f) bool {
	return ObjectSlice(l.Objects).Any(func(o Object) bool {
		return o.Contains(p)
	})
}

type Logger struct {
	Object
	Prefix string
}

func NewLogger(prefix string, o Object) *Logger {
	return &Logger{o, prefix}
}

func (l *Logger) RayInteraction(r *Ray) []*InteractionResult {
	irs := l.Object.RayInteraction(r)
	log.Printf("%s.RayInteraction(%s) = [%d]InteractionResult", l.Prefix, r, len(irs))
	return irs
}
