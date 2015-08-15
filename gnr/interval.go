package gnr

import (
	"fmt"
	"math"
)

type Interval struct {
	Min, Max float64
}

var (
	IntervalAll = Interval{
		Min: math.Inf(-1),
		Max: math.Inf(1),
	}

	IntervalPositive = Interval{
		Min: 0,
		Max: math.Inf(1),
	}

	IntervalEmpty = Interval{
		Min: 0,
		Max: -1,
	}
)

func IntervalIntersection(a, b Interval) Interval {
	if a.Empty() || b.Empty() {
		return IntervalEmpty
	}
	return Interval{
		Min: max(a.Min, b.Min),
		Max: min(a.Max, b.Max),
	}
}

func (i Interval) Intersect(other Interval) Interval {
	return IntervalIntersection(i, other)
}

func (i Interval) Normalize() Interval {
	if i.Min > i.Max {
		i.Min, i.Max = i.Max, i.Min
	}
	return i
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

func (i Interval) Empty() bool {
	return i.Max < i.Min
}

func (i Interval) String() string {
	if i.Empty() {
		return "ø"
	}
	return fmt.Sprintf("[%f;%f]", i.Min, i.Max)
}
