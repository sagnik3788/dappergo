package tracer

import "math/rand"

func newTraceID() uint64 {
	return rand.Uint64()
}

func newSpanID() uint64 {
	return rand.Uint64()
}
