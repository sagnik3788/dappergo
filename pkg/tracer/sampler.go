package tracer

import "math/rand"

type Sampler interface {
	// main Sample method
	Sample() bool
}

type ProbabilisticSampler struct {
	// controls the prob of sampling (0.0–1.0)
	Rate float64
}

func (s ProbabilisticSampler) Sample() bool {
	return rand.Float64() < s.Rate
}
