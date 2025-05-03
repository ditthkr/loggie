package middleware

import (
	"github.com/ditthkr/loggie"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func EchoZapMiddleware(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx, traceId := loggie.WithTraceId(c.Request().Context())

			ctx = loggie.WithLogger(ctx, &loggie.ZapLogger{
				L: logger.With(zap.String("trace_id", traceId)),
			})

			req := c.Request().WithContext(ctx)
			c.SetRequest(req)

			c.Response().Header().Set("X-Trace-Id", traceId)

			return next(c)
		}
	}
}
