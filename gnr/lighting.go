package gnr

import (
	"math"
	"time"
)

type FalloffFunc func(ir *InteractionResult) *InteractionResult
type FlatShader struct {
	Object
	FalloffFunc
}

func (fs *FlatShader) RayInteraction(r *Ray) []*InteractionResult {
	irs := fs.Object.RayInteraction(r)
	irs = InteractionResultSlice(irs).SelectInteractionResult(fs.FalloffFunc)
	return irs
}

func NewLinearFalloffFunc(lightDir *Vector3f) FalloffFunc {
	return func(ir *InteractionResult) *InteractionResult {
		cosAngle := VectorProduct(lightDir, ir.Normal) / lightDir.Magnitude() / ir.Normal.Magnitude()
		angle := math.Acos(cosAngle)
		ir.Color = VLerpCap(0, math.Pi, ColorBlack, ir.Color)(angle)
		return ir
	}
}

type AmbientOcclusion struct {
	Object
	NumRays     int
	MaxDistance float64
}

func (ac *AmbientOcclusion) RayInteraction(r *Ray) []*InteractionResult {
	irs := InteractionResultSlice(ac.Object.RayInteraction(r)).SortBy(InteractionResultDistance)
	if len(irs) == 0 {
		return irs
	}
	// Only use the closest hit
	ir := irs[0]

	// Shoot NumRays rays from the point of impact and darken
	// the color the more rays hit something before MaxDistance.
	rng := NewVectorSource(time.Now().UnixNano())
	n := CopyVector3f(ir.Normal).Normalize()
	// Generate rays originating from PointOfImpact
	// to a random direction, but in the hemisphere containing the normal.
	rays := RaySlice(make([]*Ray, ac.NumRays)).SelectRay(func(*Ray) *Ray {
		dir := rng.Vector3f().Normalize()
		// Mirror the direction vector at the plane of the normal if the direction
		// points away from the normal
		if VectorProduct(dir, n) < 0 {
			dir.Subtract(CopyVector3f(n).ScalarMultiply(VectorProduct(n, dir) * 2))
		}
		return &Ray{
			// TODO: Magic number :-/
			Origin:    CopyVector3f(n).ScalarMultiply(0.01).Add(ir.PointOfImpact),
			Direction: dir,
		}
	})
	// Cast rays and discard rays that don't hit anything before MaxDistance
	hits := RaySlice(rays).Where(func(r *Ray) bool {
		irs := InteractionResultSlice(ac.Object.RayInteraction(r))
		irs = irs.Where(func(ir *InteractionResult) bool {
			return ir.Distance < ac.MaxDistance
		})
		return len(irs) > 0
	})
	// Darken the color the more things are hit
	colorF := VLerpCap(0, 1, ir.Color, ColorBlack)
	ir.Color = colorF(float64(len(hits)) / float64(ac.NumRays))
	return []*InteractionResult{ir}
}
