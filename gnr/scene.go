package gnr

type Scene struct {
	GlobalLighting bool
	Camera         Camera
	Objects        []Object
}

func (s Scene) TracePixel(x, y uint64) (InteractionResult, bool) {
	r := s.Camera.GetRayForPixel(x, y)
	r.Intensity = 1.0
	return Union{s.Objects}.RayInteraction(r)
}
