package loggie

import (
	"context"
	"github.com/google/uuid"
)

type traceKey struct{}
type customFieldKey struct{}

// WithTraceId adds a randomly generated trace_id to the context.
// It returns the new context and the trace ID.
func WithTraceId(ctx context.Context) (context.Context, string) {
	traceID := uuid.NewString()
	ctx = context.WithValue(ctx, traceKey{}, traceID)
	return ctx, traceID
}

// TraceId extracts the trace_id string from the context.
// If none is found, returns "no-trace".
func TraceId(ctx context.Context) string {
	if id, ok := ctx.Value(traceKey{}).(string); ok {
		return id
	}
	return "no-trace"
}

// WithCustomField adds a custom key-value pair to the context.
// These fields will be automatically injected into log output.
func WithCustomField(ctx context.Context, key string, value any) context.Context {
	fields, _ := ctx.Value(customFieldKey{}).(map[string]any)
	if fields == nil {
		fields = make(map[string]any)
	}
	fields[key] = value
	return context.WithValue(ctx, customFieldKey{}, fields)
}
