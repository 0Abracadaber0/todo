package handlers

import (
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"todo/internal/services"
)

func getLogger(c *fiber.Ctx) *slog.Logger {
	return c.Locals("logger").(*slog.Logger)
}

func CreateTaskHandler(c *fiber.Ctx) error {
	log := getLogger(c)

	var body TaskRequest
	if err := c.BodyParser(&body); err != nil {
		log.Error("failed to parse request body", "error", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	task, err := FromTaskRequest(body)
	if err != nil {
		log.Error("failed to parse request body", "error", err)
	}
	createdTask, err := services.CreateTask(task)
	if err != nil {
		log.Error("failed to create task", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create task",
		})
	}
	log.Info("created task with", "id", createdTask.ID)
	return c.Status(fiber.StatusCreated).JSON(ToTaskResponse(createdTask))
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
