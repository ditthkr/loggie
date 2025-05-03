package ginlog

import (
	"github.com/ditthkr/loggie"
	"github.com/gin-gonic/gin"
)

// Middleware returns a generic Gin middleware that injects a logger implementing loggie.Logger.
// If no logger is provided (i.e. nil), a default no-op logger will be used.
func Middleware(logger loggie.Logger) gin.HandlerFunc {
	if logger == nil {
		logger = loggie.DefaultLogger()
	}
	return func(c *gin.Context) {
		ctx, traceId := loggie.WithTraceId(c.Request.Context())
		ctx = loggie.WithLogger(ctx, logger.With("trace_id", traceId))
		c.Request = c.Request.WithContext(ctx)

		c.Set("X-Trace-Id", traceId)
		c.Next()
	}
}
