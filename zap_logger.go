package loggie

import "go.uber.org/zap"

type ZapLogger struct {
	L *zap.Logger
}

func (r *ZapLogger) Info(msg string, fields ...any) {
	r.L.Info(msg, toZapFields(fields...)...)
}

func (r *ZapLogger) Error(msg string, fields ...any) {
	r.L.Error(msg, toZapFields(fields...)...)
}

func (r *ZapLogger) With(fields ...any) Logger {
	return &ZapLogger{L: r.L.With(toZapFields(fields...)...)}
}

func toZapFields(fields ...any) []zap.Field {
	var out []zap.Field
	for i := 0; i < len(fields); i += 2 {
		if i+1 < len(fields) {
			key, _ := fields[i].(string)
			out = append(out, zap.Any(key, fields[i+1]))
		}
	}
	return out
}
