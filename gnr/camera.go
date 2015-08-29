package gnr

import (
	"math"
)

type Camera interface {
	// GetRayForPixel creates a ray in 3D space that corresponds to the pixel on
	// the 2D canvas of the image. This should be the only function where these
	// 2 spaces meet.
	GetRayForPixel(x, y uint64) *Ray
	Normalize()
}

type SphericalCamera struct {
	Position      *Vector3f
	ViewDirection *Vector3f
	UpDirection   *Vector3f
	PixelWidth    uint64
	PixelHeight   uint64
	Angle         float64
}

func (c *SphericalCamera) Normalize() {
	c.ViewDirection.Normalize()
	side := VectorCross(c.UpDirection, c.ViewDirection)
	c.UpDirection = VectorCross(c.ViewDirection, side).Normalize()
}

func (c *SphericalCamera) GetRayForPixel(x, y uint64) *Ray {
	r := &Ray{
		Origin: CopyVector3f(c.Position),
	}

	side := VectorCross(c.UpDirection, c.ViewDirection).Normalize()

	// Note: The 3D coordinate system has inverted y axis compared
	// to the 2D coordinate system of the image canvas.
	angle := Deg2Rad(c.Angle / 2)
	xAngle := Lerp(0, float64(c.PixelWidth), -angle, angle)(float64(x))
	yAngle := Lerp(0, float64(c.PixelHeight), -angle, angle)(float64(y))
	xRotation := RotationMatrix(c.UpDirection, xAngle)
	yRotation := RotationMatrix(side, yAngle)

	r.Direction = yRotation.VectorProduct(xRotation.VectorProduct(c.ViewDirection))
	return r
}

type PlanarCamera struct {
	Position      *Vector3f
	ViewDirection *Vector3f
	UpDirection   *Vector3f
	PixelWidth    uint64
	PixelHeight   uint64
	VirtualWidth  float64
	VirtualHeight float64
	Angle         float64
}

func (c *PlanarCamera) Normalize() {
	c.ViewDirection.Normalize()
	side := VectorCross(c.UpDirection, c.ViewDirection)
	c.UpDirection = VectorCross(c.ViewDirection, side).Normalize()
}

func (c *PlanarCamera) GetRayForPixel(x, y uint64) *Ray {
	r := &Ray{
		Origin: CopyVector3f(c.Position),
	}

	side := VectorCross(c.UpDirection, c.ViewDirection).Normalize()

	// Note: The 3D coordinate system has inverted y axis compared
	// to the 2D coordinate system of the image canvas.
	projectionPlane := CopyVector3f(c.ViewDirection).ScalarMultiply(c.GetScreenDistance()).Add(c.Position)
	projectionPlane = projectionPlane.Add(side.ScalarMultiply(-c.VirtualWidth/2 + float64(x)*c.VirtualWidth/float64(c.PixelWidth)))
	projectionPlane = projectionPlane.Add(CopyVector3f(c.UpDirection).ScalarMultiply(c.VirtualHeight/2 - float64(y)*c.VirtualHeight/float64(c.PixelHeight)))

	r.Direction = projectionPlane.Subtract(c.Position).Normalize()
	return r
}

func (c *PlanarCamera) GetScreenDistance() float64 {
	return c.VirtualWidth / (2 * math.Tan(Deg2Rad(c.Angle/2)))
}

func Deg2Rad(d float64) float64 {
	return d / 360 * 2 * math.Pi
}
