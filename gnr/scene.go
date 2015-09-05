package gnr

type Scene struct {
	Camera Camera
	Object Object
}

func (s *Scene) TracePixel(x, y uint64) []*InteractionResult {
	r := s.Camera.GetRayForPixel(x, y)
	r.Intensity = 1.0
	irs := s.Object.RayInteraction(r)
	return InteractionResultSlice(irs).SortBy(InteractionResultDistance)
}
