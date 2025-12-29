package tracer

import (
	"context"
	"fmt"
	"time"

	tracerpb "github.com/sagnik3788/dappergo/pkg/proto/tracerpb"
)

type Recorder struct {
	// Buffered channel (dont block the main thread)
	SpanChan chan *SDKSpan

	// gRPC client
	Client tracerpb.CollectorClient
}

func NewRecorder(buffersize int, client tracerpb.CollectorClient) *Recorder {
	return &Recorder{
		SpanChan: make(chan *SDKSpan, buffersize),
		Client:   client,
	}
}

// RecordSpan accepts a span and pass it to the channel
func (r *Recorder) RecordSpan(s *SDKSpan) {
	select {
	case r.SpanChan <- s:
	default:
		fmt.Println("Trace buffer full, dropping span")
	}
}

// ProcessSpans runs in a background goroutine and exports completed spans
// to the collector after serialization.
func (r *Recorder) ProcessSpans() {

	for span := range r.SpanChan {
		req := &tracerpb.ExportRequest{
			Spans: []*tracerpb.Span{
				toProtoSpan(span),
			},
		}

		ctx, cancel := context.WithTimeout(
			context.Background(),
			100*time.Millisecond,
		)

		_, err := r.Client.Collect(ctx, req)
		cancel()

		if err != nil {
			continue
		}
	}
}

// convert the normal span to pb format
func toProtoSpan(s *SDKSpan) *tracerpb.Span {
	return &tracerpb.Span{
		TraceId:           s.TraceID,
		SpanId:            s.SpanID,
		ParentId:          s.ParentID,
		Name:              s.Name,
		StartTimeUnixNano: s.StartTime.UnixNano(),
		EndTimeUnixNano:   s.EndTime.UnixNano(),
	}
}
