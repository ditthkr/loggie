package loggie

import "context"

// Logger is the main interface used for logging.
// It mimics a simplified version of structured loggers like Zap or Logrus.
type Logger interface {
	Info(msg string, fields ...any)
	Error(msg string, fields ...any)
	With(fields ...any) Logger
}

type ctxKey struct{}

// WithLogger stores the given logger inside the context.
// It can be retrieved later using FromContext.
func WithLogger(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, ctxKey{}, logger)
}

// FromContext retrieves the logger from context.
// If no logger is found, it returns a default no-op logger.
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
