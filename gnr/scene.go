package gnr

type Scene struct {
	GlobalLighting bool
	Camera         Camera
	Object         Object
}

func (s Scene) TracePixel(x, y uint64) (InteractionResult, bool) {
	r := s.Camera.GetRayForPixel(x, y)
	r.Intensity = 1.0
	return s.Object.RayInteraction(r)
}
