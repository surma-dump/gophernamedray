package object

import (
	"math"

	"github.com/surma-dump/gophernamedray/gnr"
)

type AxisAlignedBox struct {
	Min, Max gnr.Vector3f
}

func (aab AxisAlignedBox) RayInteraction(r gnr.Ray) []gnr.InteractionResult {
	planes := []gnr.Object{
		Plane{
			Normal:   gnr.Vector3f{-1, 0, 0},
			Distance: -gnr.VectorProduct(gnr.Vector3f{-1, 0, 0}, aab.Min),
		},
		Plane{
			Normal:   gnr.Vector3f{1, 0, 0},
			Distance: -gnr.VectorProduct(gnr.Vector3f{1, 0, 0}, aab.Max),
		},
		Plane{
			Normal:   gnr.Vector3f{0, -1, 0},
			Distance: -gnr.VectorProduct(gnr.Vector3f{0, -1, 0}, aab.Min),
		},
		Plane{
			Normal:   gnr.Vector3f{0, 1, 0},
			Distance: -gnr.VectorProduct(gnr.Vector3f{0, 1, 0}, aab.Max),
		},
		Plane{
			Normal:   gnr.Vector3f{0, 0, -1},
			Distance: -gnr.VectorProduct(gnr.Vector3f{0, 0, -1}, aab.Min),
		},
		Plane{
			Normal:   gnr.Vector3f{0, 0, 1},
			Distance: -gnr.VectorProduct(gnr.Vector3f{0, 0, 1}, aab.Max),
		},
	}
	irs := gnr.ObjectSlice(planes).AggregateSliceInteractionResult(func(irs []gnr.InteractionResult, o gnr.Object) []gnr.InteractionResult {
		newIrs := o.RayInteraction(r)
		p := o.(Plane)
		newIrs = gnr.InteractionResultSlice(newIrs).SelectInteractionResult(func(ir gnr.InteractionResult) gnr.InteractionResult {
			// Correct for IEEE754 errors by aligning PoIs with sides of box
			if math.Abs(p.Normal.X) == 1 {
				ir.PointOfImpact.X = -p.Distance * p.Normal.X
			}
			if math.Abs(p.Normal.Y) == 1 {
				ir.PointOfImpact.Y = -p.Distance * p.Normal.Y
			}
			if math.Abs(p.Normal.Z) == 1 {
				ir.PointOfImpact.Z = -p.Distance * p.Normal.Z
			}
			return ir
		})
		return append(irs, newIrs...)
	})
	irs = gnr.InteractionResultSlice(irs).Where(func(ir gnr.InteractionResult) bool {
		return aab.Contains(ir.PointOfImpact)
	})
	return irs
}

func (aab AxisAlignedBox) Contains(p gnr.Vector3f) bool {
	return p.X >= aab.Min.X && p.X <= aab.Max.X &&
		p.Y >= aab.Min.Y && p.Y <= aab.Max.Y &&
		p.Z >= aab.Min.Z && p.Z <= aab.Max.Z
}

func (aab AxisAlignedBox) Normalize() AxisAlignedBox {
	if aab.Min.X > aab.Max.X {
		aab.Min.X, aab.Max.X = aab.Max.X, aab.Min.X
	}
	if aab.Min.Y > aab.Max.Y {
		aab.Min.Y, aab.Max.Y = aab.Max.Y, aab.Min.Y
	}
	if aab.Min.Z > aab.Max.Z {
		aab.Min.Z, aab.Max.Z = aab.Max.Z, aab.Min.Z
	}
	return aab
}
