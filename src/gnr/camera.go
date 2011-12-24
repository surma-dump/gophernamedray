package gnr

import (
	"math"
)

type Camera struct {
	Position      Vector3f
	PixelWidth    uint64
	PixelHeight   uint64
	VirtualWidth  float64
	VirtualHeight float64
	Angle         float64
}

func (c *Camera) GetRayForPixel(x, y uint64) *Ray {
	r := &Ray{
		Origin: c.Position,
	}

	tmpdir := Vector3f{0, 0, 1}
	screenPoint := VectorSum(&c.Position, tmpdir.Multiply(c.GetScreenDistance()))
	screenPoint.X -= c.VirtualWidth / 2.0
	screenPoint.Y -= c.VirtualHeight / 2.0

	screenPoint.X += float64(x) * c.VirtualWidth / float64(c.PixelWidth)
	screenPoint.Y += float64(y) * c.VirtualHeight / float64(c.PixelHeight)

	r.Direction = *VectorDifference(screenPoint, &c.Position)
	r.Direction.Normalize()
	return r
}

func (c *Camera) GetScreenDistance() float64 {
	return c.VirtualWidth / 2 * math.Tan(c.Angle*2*math.Pi/2/360)
}
