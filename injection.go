package loggie

import (
	"context"
)

// Injection injects logger and trace_id into a context.
// Use it inside any web framework middleware.
func Injection(ctx context.Context, logger Logger) (context.Context, string) {
	if logger == nil {
		logger = DefaultLogger()
	}
	ctx, traceId := WithTraceId(ctx)
	ctx = WithLogger(ctx, logger)
	return ctx, traceId
}
