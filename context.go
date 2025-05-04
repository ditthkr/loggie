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

var defaultLogger Logger = &noopLogger{}

type noopLogger struct{}

func (n *noopLogger) Info(msg string, fields ...any)  {}
func (n *noopLogger) Error(msg string, fields ...any) {}
func (n *noopLogger) With(fields ...any) Logger       { return n }

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

	fields := []any{"trace_id", TraceId(ctx)}

	// Add custom fields if any
	if custom, ok := ctx.Value(customFieldKey{}).(map[string]any); ok {
		for k, v := range custom {
			fields = append(fields, k, v)
		}
	}

	return logger.With(fields...)
}

// DefaultLogger returns the internal fallback logger used when no logger is found.
// This logger does nothing and is safe to call anywhere.
func DefaultLogger() Logger {
	return defaultLogger
}
