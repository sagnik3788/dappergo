package collector

import (
	"context"
	"fmt"
	"time"

	"github.com/sagnik3788/dappergo/pkg/proto/tracerpb"
)

type CollectorService struct {
	tracerpb.UnimplementedCollectorServer
}

func NewCollectorService() *CollectorService {
	return &CollectorService{}
}

func (c *CollectorService) Collect(ctx context.Context, req *tracerpb.ExportRequest) (*tracerpb.ExportResponse, error) {
	for _, sp := range req.Spans {
		start := time.Unix(0, sp.StartTimeUnixNano)
		end := time.Unix(0, sp.EndTimeUnixNano)
		duration := end.Sub(start)

		fmt.Printf("TraceID=%d SpanID=%d ParentID=%d Name=%s Duration=%v\n",
			sp.TraceId, sp.SpanId, sp.ParentId, sp.Name, duration)
	}

	return &tracerpb.ExportResponse{Ok: true}, nil
}
