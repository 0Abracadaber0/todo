package handlers

import (
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"todo/internal/models"
	"todo/internal/services"
)

type TaskResponse struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Overdue     bool   `json:"overdue"`
	Completed   bool   `json:"completed"`
}

func getLogger(c *fiber.Ctx) *slog.Logger {
	return c.Locals("logger").(*slog.Logger)
}

func toTaskResponse(task models.Task) TaskResponse {
	var dueDate string
	if task.DueDate.Time.IsZero() {
		dueDate = ""
	} else {
		dueDate = task.DueDate.Time.Format(models.DateFormat)
	}

	return TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		DueDate:     dueDate,
		Overdue:     task.Overdue,
		Completed:   task.Completed,
	}
}

func CreateTaskHandler(c *fiber.Ctx) error {
	log := getLogger(c)

	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		log.Error("failed to parse request body", "error", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	createdTask, err := services.CreateTask(task)
	if err != nil {
		log.Error("failed to create task", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create task",
		})
	}
	log.Info("created task with", "id", createdTask.ID)
	return c.Status(fiber.StatusCreated).JSON(toTaskResponse(createdTask))
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
