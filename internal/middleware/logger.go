package middleware

import (
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"time"
)

func LoggerMiddleware(log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		c.Locals("logger", log)
		err := c.Next()

		log.Info("HTTP request",
			slog.String("method", c.Method()),
			slog.String("route", c.Path()),
			slog.Int("status", c.Response().StatusCode()),
			slog.Duration("latency", time.Since(start)),
		)

		return err
	}
}
