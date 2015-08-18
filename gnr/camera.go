package gnr

import (
	"math"
)

type Camera struct {
	Position      Vector3f
	ViewDirection Vector3f
	UpDirection   Vector3f
	PixelWidth    uint64
	PixelHeight   uint64
	VirtualWidth  float64
	VirtualHeight float64
	Angle         float64
}

func (c Camera) Normalize() Camera {
	c.ViewDirection = c.ViewDirection.Normalize()
	side := VectorCross(c.ViewDirection, c.UpDirection)
	c.UpDirection = VectorCross(side, c.ViewDirection).Normalize()
	return c
}

// GetRayForPixel creates a ray in 3D space that corresponds to the pixel on
// the 2D canvas of the image. This should be the only function where these
// 2 spaces meet. Camera must be normalized.
func (c Camera) GetRayForPixel(x, y uint64) Ray {
	r := Ray{
		Origin: c.Position,
	}

	side := VectorCross(c.UpDirection, c.ViewDirection).Normalize()

	// Note: The 3D coordinate system has inverted y axis compared
	// to the 2D coordinate system of the image canvas.
	projectionPlane := VectorSum(c.Position, c.ViewDirection.Multiply(c.GetScreenDistance()))
	projectionPlane = VectorSum(projectionPlane, side.Multiply(-c.VirtualWidth/2))
	projectionPlane = VectorSum(projectionPlane, c.UpDirection.Multiply(c.VirtualHeight/2))

	projectionPlane = VectorSum(projectionPlane, side.Multiply(float64(x)*c.VirtualWidth/float64(c.PixelWidth)))
	projectionPlane = VectorSum(projectionPlane, c.UpDirection.Multiply(-float64(y)*c.VirtualHeight/float64(c.PixelHeight)))

	r.Direction = VectorDifference(projectionPlane, c.Position).Normalize()
	return r
}

func (c Camera) GetScreenDistance() float64 {
	return c.VirtualWidth / (2 * math.Tan(c.Angle/2/360*2*math.Pi))
}
