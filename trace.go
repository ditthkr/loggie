package loggie

import (
	"context"
	"github.com/google/uuid"
)

type traceKey struct{}
type customFieldKey struct{}

func WithTraceID(ctx context.Context) (context.Context, string) {
	traceID := uuid.NewString()
	ctx = context.WithValue(ctx, traceKey{}, traceID)
	return ctx, traceID
}

func WithCustomField(ctx context.Context, key string, value any) context.Context {
	fields, _ := ctx.Value(customFieldKey{}).(map[string]any)
	if fields == nil {
		fields = make(map[string]any)
	}
	fields[key] = value
	return context.WithValue(ctx, customFieldKey{}, fields)
}
