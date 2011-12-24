package gnr

import ()

type Scene struct {
	GlobalLighting bool
	Camera         Camera
	Objects        []Object
}

func (s *Scene) TracePixel(x, y uint64) *Color {
	r := s.Camera.GetRayForPixel(x, y)
	r.Intensity = 1.0
	return s.ShootRay(r)
}

func (s *Scene) ShootRay(r *Ray) *Color {
	for _, obj := range s.Objects {
		if obj.RayCollision(r) {
			c, r := obj.RayManipulation(r)
			_ = r
			return c

		}
	}
	return COLOR_BLACK
}
