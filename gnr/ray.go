package gnr

import (
	"fmt"
)

type Ray struct {
	Origin    *Vector3f
	Direction *Vector3f
	Intensity float64
}

func (r *Ray) String() string {
	return fmt.Sprintf("%s+t*%s (%0.3f)", r.Origin.String(), r.Direction.String(), r.Intensity)
}
