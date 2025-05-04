package loggie

import "github.com/sirupsen/logrus"

type LogrusLogger struct {
	L *logrus.Entry
}

func (l *LogrusLogger) Info(msg string, fields ...any) {
	l.L.WithFields(toLogrusFields(fields...)).Info(msg)
}

func (l *LogrusLogger) Error(msg string, fields ...any) {
	l.L.WithFields(toLogrusFields(fields...)).Error(msg)
}

func (l *LogrusLogger) With(fields ...any) Logger {
	return &LogrusLogger{
		L: l.L.WithFields(toLogrusFields(fields...)),
	}
}

func toLogrusFields(fields ...any) logrus.Fields {
	out := logrus.Fields{}
	for i := 0; i < len(fields)-1; i += 2 {
		if key, ok := fields[i].(string); ok {
			out[key] = fields[i+1]
		}
	}
	return out
}
