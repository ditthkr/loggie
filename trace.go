package loggie

import (
	"context"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
)

type traceKeyType struct{}

var traceKey = traceKeyType{}

type customFieldKey struct{}

// WithTraceId returns a new context with a generated or existing trace ID,
// and the trace ID itself as a string.
//
// It first checks if a trace ID exists in the context, otherwise it generates a new UUID.
// This function is compatible with both OpenTelemetry-based contexts and local fallback mode.
func WithTraceId(ctx context.Context) (context.Context, string) {
	traceId := TraceId(ctx)
	return context.WithValue(ctx, traceKey, traceId), traceId
}

// TraceId returns the current trace ID from the context.
//
// It checks the following, in order:
//  1. If an OpenTelemetry span exists in the context and has a valid TraceID, it is returned.
//  2. If a local trace ID was previously stored in context via WithTraceID(), it is returned.
//  3. Otherwise, a new UUID is generated and returned.
func TraceId(ctx context.Context) string {
	// Check OpenTelemetry span context first
	if span := trace.SpanFromContext(ctx); span != nil {
		sc := span.SpanContext()
		if sc.IsValid() {
			return sc.TraceID().String()
		}
	}

	// Fallback: return from context if available
	if v := ctx.Value(traceKey); v != nil {
		if tid, ok := v.(string); ok {
			return tid
		}
	}

	// Final fallback: generate new UUID
	return uuid.New().String()
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
