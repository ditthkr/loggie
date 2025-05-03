package middleware

import (
	"github.com/ditthkr/loggie"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GinZapMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, traceId := loggie.WithTraceID(c.Request.Context())

		ctx = loggie.WithLogger(ctx, &loggie.ZapLogger{
			L: logger.With(zap.String("trace_id", traceId)),
		})
		c.Request = c.Request.WithContext(ctx)

		c.Set("X-Trace-Id", traceId)
		c.Next()
	}
}
