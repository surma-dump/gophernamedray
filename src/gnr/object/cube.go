package object

import (
	"gnr"
)

type Cube struct {
	CornerMin, CornerMax gnr.Vector3f
}

func (c *Cube) Normalize() {
	if c.CornerMin.X > c.CornerMax.X {
		c.CornerMin.X, c.CornerMax.X = c.CornerMax.X, c.CornerMin.X
	}
	if c.CornerMin.Y > c.CornerMax.Y {
		c.CornerMin.Y, c.CornerMax.Y = c.CornerMax.Y, c.CornerMin.Y
	}
	if c.CornerMin.Z > c.CornerMax.Z {
		c.CornerMin.Z, c.CornerMax.Z = c.CornerMax.Z, c.CornerMin.Z
	}
}

func (c *Cube) RayCollision(r *gnr.Ray) bool {
	x := (&gnr.Interval{
		Min: (c.CornerMin.X - r.Origin.X) / (r.Direction.X),
		Max: (c.CornerMax.X - r.Origin.X) / (r.Direction.X),
	}).Normalize().Intersect(gnr.INTERVAL_POSITIVE)

	y := (&gnr.Interval{
		Min: (c.CornerMin.Y - r.Origin.Y) / (r.Direction.Y),
		Max: (c.CornerMax.Y - r.Origin.Y) / (r.Direction.Y),
	}).Normalize().Intersect(gnr.INTERVAL_POSITIVE)

	z := (&gnr.Interval{
		Min: (c.CornerMin.Z - r.Origin.Z) / (r.Direction.Z),
		Max: (c.CornerMax.Z - r.Origin.Z) / (r.Direction.Z),
	}).Normalize().Intersect(gnr.INTERVAL_POSITIVE)

	return !x.Intersect(y.Intersect(z)).Empty()
}

func (c *Cube) RayManipulation(r *gnr.Ray) (*gnr.Color, []*gnr.Ray) {
	return gnr.COLOR_WHITE, []*gnr.Ray{}
}
