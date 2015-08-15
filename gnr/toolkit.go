package gnr

import (
	"math"
)

// Lerp returns a function that performs a lerp on the given parameters. Lerp
// does not check boundaries on the input value.
func Lerp(inMin, inMax, outMin, outMax float64) func(d float64) float64 {
	return func(d float64) float64 {
		return (d-inMin)/(inMax-inMin)*(outMax-outMin) + outMin
	}
}

// LerpCap is the same as Lerp but caps the input to be in between
// inMin and inMax.
func LerpCap(inMin, inMax, outMin, outMax float64) func(d float64) float64 {
	f := Lerp(inMin, inMax, outMin, outMax)
	return func(d float64) float64 {
		return f(math.Max(math.Min(d, inMax), inMin))
	}
}
