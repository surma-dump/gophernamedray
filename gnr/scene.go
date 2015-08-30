package gnr

type Scene struct {
	Camera Camera
	Object Object
}

func (s *Scene) TracePixel(x, y uint64) (*InteractionResult, bool) {
	r := s.Camera.GetRayForPixel(x, y)
	r.Intensity = 1.0
	irs := s.Object.RayInteraction(r)
	if len(irs) <= 0 {
		return &InteractionResult{
			Color:         ColorBlack,
			PointOfImpact: &Vector3f{0, 0, 0},
			Normal:        &Vector3f{0, 0, 0},
		}, false
	}
	return InteractionResultSlice(irs).SortBy(InteractionResultDistance)[0], true
}
