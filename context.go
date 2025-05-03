package loggie

import "context"

type Logger interface {
	Info(msg string, fields ...any)
	Error(msg string, fields ...any)
	With(fields ...any) Logger
}

type ctxKey struct{}

func WithLogger(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, ctxKey{}, logger)
}

func FromContext(ctx context.Context) Logger {
	logger, ok := ctx.Value(ctxKey{}).(Logger)
	if !ok {
		logger = defaultLogger
	}

	fields, _ := ctx.Value(customFieldKey{}).(map[string]any)
	if len(fields) == 0 {
		return logger
	}

	var kv []any
	for k, v := range fields {
		kv = append(kv, k, v)
	}
	return logger.With(kv...)
}
