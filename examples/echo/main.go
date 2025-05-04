package main

import (
	"context"
	"github.com/ditthkr/loggie"
	"github.com/ditthkr/loggie/middleware/echolog"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {

	// Logrus

	logger := logrus.New()
	rawLogger := logrus.NewEntry(logger)

	e := echo.New()
	e.Use(echolog.Middleware(&loggie.LogrusLogger{L: rawLogger}))

	// Zap

	//rawLogger, _ := zap.NewProduction(zap.AddCallerSkip(1))
	//defer rawLogger.Sync()
	//
	//e := echo.New()
	//e.Use(echolog.Middleware(&loggie.ZapLogger{L: rawLogger}))

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
