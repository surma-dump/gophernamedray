package gnr

import (
	"math/rand"
)

// VectorSource represents a source of Vector3f, each component
// being uniformly-distributed pseudo-random in range [0, 1).
type VectorSource interface {
	Vector3f() *Vector3f
	Seed(seed int64)
}

func NewVectorSource(seed int64) VectorSource {
	return &randVectorSource{
		rnd: rand.New(rand.NewSource(seed)),
	}
}

type randVectorSource struct {
	rnd *rand.Rand
}

func (rvs *randVectorSource) Vector3f() *Vector3f {
	return &Vector3f{
		X: rvs.rnd.Float64(),
		Y: rvs.rnd.Float64(),
		Z: rvs.rnd.Float64(),
	}
}

func (rvs *randVectorSource) Seed(seed int64) {
	rvs.rnd.Seed(seed)
}
