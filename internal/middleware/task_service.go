package middleware

import (
	"github.com/gofiber/fiber/v2"
	"todo/internal/services"
)

func TaskServiceMiddleware(taskService *services.TaskService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("taskService", taskService)
		return c.Next()
	}
}
