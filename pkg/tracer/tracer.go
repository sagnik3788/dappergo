package tracer

import (
	"context"
	"time"
)

type Tracer struct {
	// Recorder for passing to tracer
	Recorder *Recorder

	// Sampler for passing to tracer
	Sampler Sampler
}

// eg. recorder := NewRecorder(1024, grpcClient)
// tracer := NewTracer(recorder)
func NewTracer(r *Recorder, s Sampler) *Tracer {
	return &Tracer{
		Recorder: r,
		Sampler:  s,
	}
}

// StartSpan create a parent span for first service
// eg. StartSpan('payment-service') it will return the whole span in duration and ids
func (t *Tracer) StartSpan(ctx context.Context, name string) (context.Context, *SDKSpan) {
	sampled := t.Sampler.Sample()
	span := &SDKSpan{
		TraceID:   newTraceID(),
		SpanID:    newSpanID(),
		ParentID:  0,
		Name:      name,
		Sampled:   sampled,
		StartTime: time.Now(),
		Tracer:    t,
	}

	return Inject_Context(ctx, span), span
}

// StartSpan create a child span for next each service
func (t *Tracer) StartChildSpan(ctx context.Context, name string) *SDKSpan {
	parent := Pull_Context(ctx)

	return &SDKSpan{
		TraceID:   parent.TraceID,
		SpanID:    newSpanID(),
		ParentID:  parent.SpanID,
		Name:      name,
		Sampled:   parent.Sampled,
		StartTime: time.Now(),
		Tracer:    t,
	}
}

func (t *Tracer) recordSpan(s *SDKSpan) {
	if !s.Sampled {
		return
	}
	t.Recorder.RecordSpan(s)
}
