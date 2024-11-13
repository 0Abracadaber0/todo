package handlers

import (
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"todo/internal/models"
	"todo/internal/services"
)

func getTaskService(c *fiber.Ctx) *services.TaskService {
	return c.Locals("taskService").(*services.TaskService)
}

func getLogger(c *fiber.Ctx) *slog.Logger {
	return c.Locals("logger").(*slog.Logger)
}

func CreateTaskHandler(c *fiber.Ctx) error {
	taskService := getTaskService(c)
	log := getLogger(c)

	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		log.Error("failed to parse request body", "error", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	createdTask, err := taskService.CreateTask(log, task)
	if err != nil {
		log.Error("failed to create task", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create task",
		})
	}
	log.Info("created task", "id", createdTask.ID)
	return c.Status(fiber.StatusCreated).JSON(createdTask)
}

func ListTaskHandler(c *fiber.Ctx) error {
	return nil
}

func UpdateTaskHandler(c *fiber.Ctx) error {
	return nil
}

func DeleteTaskHandler(c *fiber.Ctx) error {
	return nil
}

func CompleteTaskHandler(c *fiber.Ctx) error {
	return nil
}
