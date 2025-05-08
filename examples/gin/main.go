package main

import (
	"context"
	"github.com/ditthkr/loggie"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"net/http"
)

func main() {

	// Zap

	rawLogger, _ := zap.NewProduction(zap.AddCallerSkip(1))
	defer rawLogger.Sync()

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		ctx, traceId := loggie.Injection(c.Request.Context(), &loggie.ZapLogger{L: rawLogger})
		c.Request = c.Request.WithContext(ctx)
		c.Writer.Header().Set("X-Trace-Id", traceId)
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		ctx := c.Request.Context()

		ctx = loggie.WithCustomField(ctx, "user_id", 1234)
		ctx = loggie.WithCustomField(ctx, "order_id", "ORD-5678")

		log := loggie.FromContext(ctx)
		log.Info("received /ping request")

		if err := processPing(ctx); err != nil {
			log.Error("ping processing failed", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed"})
			return
		}

		log.Info("successfully responded to /ping")
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	err := r.Run()
	if err != nil {
		return
	}
}

func processPing(ctx context.Context) error {
	log := loggie.FromContext(ctx)
	log.Info("start processPing")

	if err := queryDatabase(ctx); err != nil {
		log.Error("query failed", "error", err)
		return err
	}

	log.Info("finish processPing")
	return nil
}

func queryDatabase(ctx context.Context) error {
	log := loggie.FromContext(ctx)

	log.Info("querying database", "sql", "_SELECT * FROM pings LIMIT 1")

	return nil
}
