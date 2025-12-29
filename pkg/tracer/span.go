package tracer

import "time"

type SDKSpan struct {
	// Id for the whole req
	TraceID uint64

	// Id for the specific service
	SpanID uint64

	// Id for the parent span
	ParentID uint64

	// Label for each span
	Name string

	// Sampled gives sampling decsion trace or not
	Sampled bool

	// StartTime of the span
	StartTime time.Time

	// EndTime of the span
	EndTime time.Time

	// TODO
	// Annotations map[string]string

	// where to send the span when done
	Tracer *Tracer
}

// stops and sends the span to the async recorder
func (s *SDKSpan) Finish() {
	s.EndTime = time.Now()
	s.Tracer.recordSpan(s)
}
