package tracer

import (
	"context"
	"fmt"
	"time"

	tracerpb "github.com/sagnik3788/dappergo/pkg/proto/tracerpb"
)

type Recorder struct {
	// buffered channel (dont block the main thread)
	spanChan chan *SDKSpan

	// grpc client
	client tracerpb.CollectorClient
}

func NewRecorder(buffersize int, client tracerpb.CollectorClient) *Recorder {
	return &Recorder{
		spanChan: make(chan *SDKSpan, buffersize),
		client:   client,
	}
}

// RecordSpan accepts a span and pass it to the channel
func (r *Recorder) RecordSpan(s *SDKSpan) {
	select {
	case r.spanChan <- s:
	default:
		fmt.Println("Trace buffer full, dropping span")
	}
}

// ProcessSpans runs in a background goroutine and exports completed spans
// to the collector after serialization.
func (r *Recorder) ProcessSpans() {

	for span := range r.spanChan {
		req := &tracerpb.ExportRequest{
			Spans: []*tracerpb.Span{
				toProtoSpan(span),
			},
		}

		ctx, cancel := context.WithTimeout(
			context.Background(),
			100*time.Millisecond,
		)

		_, err := r.client.Collect(ctx, req)
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
