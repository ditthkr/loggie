package fiberlog

import (
	"github.com/ditthkr/loggie"
	"github.com/gofiber/fiber/v3"
)

// Middleware returns a generic Fiber middleware that injects a logger implementing loggie.Logger.
// If no logger is provided (i.e. nil), a default no-op logger will be used.
func Middleware(logger loggie.Logger) fiber.Handler {
	if logger == nil {
		logger = loggie.DefaultLogger()
	}
	return func(c fiber.Ctx) error {
		ctx, traceId := loggie.WithTraceId(c.Context())
		ctx = loggie.WithLogger(ctx, logger.With("trace_id", traceId))
		c.SetContext(ctx)
		c.Set("X-Trace-Id", traceId)
		return c.Next()
	}
}
