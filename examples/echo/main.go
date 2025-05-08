package main

import (
	"context"
	"github.com/ditthkr/loggie"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"net/http"
)

func main() {

	// Zap

	rawLogger, _ := zap.NewProduction(zap.AddCallerSkip(1))
	defer rawLogger.Sync()

	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx, traceId := loggie.Injection(c.Request().Context(), &loggie.ZapLogger{L: rawLogger})
			req := c.Request().WithContext(ctx)
			c.SetRequest(req)
			c.Response().Header().Set("X-Trace-Id", traceId)
			return next(c)
		}
	})

	e.GET("/ping", func(c echo.Context) error {
		ctx := c.Request().Context()
		ctx = loggie.WithCustomField(ctx, "user_id", 4321)

		log := loggie.FromContext(ctx)
		log.Info("received /ping request")

		if err := processPing(ctx); err != nil {
			log.Error("ping processing failed", "error", err)

			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed"})
		}

		log.Info("successfully responded to /ping")
		return c.JSON(http.StatusOK, map[string]string{
			"message": "pong",
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
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
