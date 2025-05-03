package main

import (
	"context"
	"github.com/ditthkr/loggie"
	"github.com/ditthkr/loggie/middleware/ginlog"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	rawLogger, _ := zap.NewProduction()
	defer rawLogger.Sync()

	r := gin.Default()
	r.Use(ginlog.Middleware(&loggie.ZapLogger{L: rawLogger}))

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

	r.Run()
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

	log.Info("querying database", "sql", "SELECT * FROM pings LIMIT 1")

	return nil
}
