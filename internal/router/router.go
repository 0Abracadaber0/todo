package router

import (
	"github.com/gofiber/fiber/v2"
	"todo/internal/handlers"
)

func SetupRoutes(app *fiber.App) {

	app.Post("/tasks", handlers.CreateTaskHandler)
	app.Get("/tasks", handlers.ListTaskHandler)
	app.Put("/tasks/:id", handlers.UpdateTaskHandler)
	app.Delete("/tasks/:id", handlers.DeleteTaskHandler)
	app.Patch("/tasks/:id/complete", handlers.CompleteTaskHandler)
}
