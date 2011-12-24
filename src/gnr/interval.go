package gnr

import (
	"fmt"
	"math"
)

type Interval struct {
	Min, Max float64
}

var (
	INTERVAL_ALL = &Interval{
		Min: math.Inf(-1),
		Max: math.Inf(1),
	}

	INTERVAL_POSITIVE = &Interval{
		Min: 0,
		Max: math.Inf(1),
	}
)

func IntervalIntersection(a, b *Interval) *Interval {
	if a.Empty() || b.Empty() {
		return &Interval{
			Min: 0,
			Max: -1,
		}
	}
	return &Interval{
		Min: max(a.Min, b.Min),
		Max: min(a.Max, b.Max),
	}
}

func (i *Interval) Intersect(other *Interval) *Interval {
	return IntervalIntersection(i, other)
}

func (i *Interval) Normalize() *Interval {
	r := *i
	if i.Min > i.Max {
		r.Min, r.Max = r.Max, r.Min
	}
	return &r
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func (i *Interval) Empty() bool {
	return i.Max < i.Min
}

func (i *Interval) String() string {
	if i.Empty() {
		return "ø"
	}
	return fmt.Sprintf("[%f;%f]", i.Min, i.Max)
}
