package echolog

import (
	"github.com/ditthkr/loggie"
	"github.com/labstack/echo/v4"
)

// Middleware returns a generic Echo middleware that injects a logger implementing loggie.Logger.
// If no logger is provided (i.e. nil), a default no-op logger will be used.
func Middleware(logger loggie.Logger) echo.MiddlewareFunc {
	if logger == nil {
		logger = loggie.DefaultLogger()
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx, traceId := loggie.WithTraceId(c.Request().Context())
			ctx = loggie.WithLogger(ctx, logger)
			req := c.Request().WithContext(ctx)
			c.SetRequest(req)
			c.Response().Header().Set("X-Trace-Id", traceId)
			return next(c)
		}
	}
}
