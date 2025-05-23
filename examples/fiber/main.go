package main

import (
	"context"
	"github.com/ditthkr/loggie"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"net/http"
)

func main() {

	// Zap

	rawLogger, _ := zap.NewProduction(zap.AddCallerSkip(1))
	defer rawLogger.Sync()

	app := fiber.New()
	//app.Use(fiberlog.Middleware(&loggie.ZapLogger{L: rawLogger}))

	app.Use(func(c fiber.Ctx) error {
		ctx, traceId := loggie.Injection(c.Context(), &loggie.ZapLogger{L: rawLogger})
		c.SetContext(ctx)
		c.Set("X-Trace-Id", traceId)
		return c.Next()
	})

	app.Get("/ping", func(c fiber.Ctx) error {
		ctx := c.Context()
		ctx = loggie.WithCustomField(ctx, "user_id", 4321)

		log := loggie.FromContext(ctx)
		log.Info("received /ping request")

		if err := processPing(ctx); err != nil {
			log.Error("ping processing failed", "error", err)

			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed"})
		}

		log.Info("successfully responded to /ping")

		return c.JSON(fiber.Map{"message": "pong"})
	})

	app.Listen(":8080")
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
