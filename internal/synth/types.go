package synth

//lint:file-ignore U1000 Some providers may be unused during experiments.

import "math/rand"

// constant always returns same value.
func constant[T any](value T) func() T {
	return func() T {
		return value
	}
}

// randomf64 returns float64 within given bounds.
func randomf64(boundMin float64, boundMax float64) func() float64 {
	return func() float64 {
		return boundMin + rand.Float64()*(boundMax-boundMin)
	}
}

// Provider is a shorthand alias for parameter-returning callables.
type Provider[T any] = func() T
