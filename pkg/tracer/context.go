package tracer

import "context"

type ctxkey struct{}

var spanIDKey ctxkey

// pass the parent span in the context
func Inject_Context(ctx context.Context, span *SDKSpan) context.Context {
	return context.WithValue(ctx, spanIDKey, span)
}

// pull the context in child spans
func Pull_Context(ctx context.Context) *SDKSpan {
	span, _ := ctx.Value(spanIDKey).(*SDKSpan)
	return span
}
