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

// GetRayForPixel creates a ray in 3D space that corresponds to the pixel on
// the 2D canvas of the image. This should be the only function where these
// 2 spaces meet.
func (c Camera) GetRayForPixel(x, y uint64) Ray {
	r := Ray{
		Origin: c.Position,
	}

	// Note: The 3D coordinate system has inverted y axis compared
	// to the 2D coordinate system of the image canvas.
	tmpdir := Vector3f{0, 0, 1}
	screenPoint := VectorSum(c.Position, tmpdir.Multiply(c.GetScreenDistance()))
	screenPoint.X -= c.VirtualWidth / 2.0
	screenPoint.Y += c.VirtualHeight / 2.0

	screenPoint.X += float64(x) * c.VirtualWidth / float64(c.PixelWidth)
	screenPoint.Y -= float64(y) * c.VirtualHeight / float64(c.PixelHeight)

	r.Direction = VectorDifference(screenPoint, c.Position).Normalize()
	return r
}

func (c Camera) GetScreenDistance() float64 {
	return c.VirtualWidth / 2 * math.Tan(c.Angle*2*math.Pi/2/360)
}
