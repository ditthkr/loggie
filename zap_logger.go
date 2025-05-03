package loggie

import "go.uber.org/zap"

type ZapLogger struct {
	L *zap.Logger
}

func (z *ZapLogger) Info(msg string, fields ...any) {
	z.L.Info(msg, toZapFields(fields...)...)
}

func (z *ZapLogger) Error(msg string, fields ...any) {
	z.L.Error(msg, toZapFields(fields...)...)
}

func (z *ZapLogger) With(fields ...any) Logger {
	return &ZapLogger{L: z.L.With(toZapFields(fields...)...)}
}

func toZapFields(fields ...any) []zap.Field {
	var zfields []zap.Field
	for i := 0; i < len(fields); i += 2 {
		if i+1 < len(fields) {
			key, _ := fields[i].(string)
			zfields = append(zfields, zap.Any(key, fields[i+1]))
		}
	}
	return zfields
}
