package main

import (
	"context"
	"github.com/ditthkr/loggie"
	"github.com/ditthkr/loggie/middleware"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	rawLogger, _ := zap.NewProduction()
	defer rawLogger.Sync()

	app := fiber.New()
	app.Use(middleware.FiberZapMiddleware(rawLogger))

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
