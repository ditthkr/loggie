package middleware

import (
	"github.com/ditthkr/loggie"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

func FiberZapMiddleware(logger *zap.Logger) fiber.Handler {
	return func(c fiber.Ctx) error {
		ctx, traceId := loggie.WithTraceId(c.Context())

		ctx = loggie.WithLogger(ctx, &loggie.ZapLogger{
			L: logger.With(zap.String("trace_id", traceId)),
		})
		c.SetContext(ctx)

		c.Set("X-Trace-Id", traceId)
		return c.Next()
	}
}
