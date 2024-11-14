package handlers

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"strconv"
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
	log.Info("task successfully created with", "id", createdTask.ID)
	return c.Status(fiber.StatusCreated).JSON(ToTaskResponse(createdTask))
}

func ListTaskHandler(c *fiber.Ctx) error {
	log := getLogger(c)

	tasks, err := services.GetTasks()
	if err != nil {
		log.Error("failed to get tasks", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to get tasks",
		})
	}

	responseTasks := make([]TaskResponse, len(tasks))
	for i, task := range tasks {
		responseTasks[i] = ToTaskResponse(task)
	}

	log.Info("tasks successfully received", "count", len(responseTasks))
	return c.Status(fiber.StatusOK).JSON(responseTasks)
}

func UpdateTaskHandler(c *fiber.Ctx) error {
	log := getLogger(c)

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Error("failed to get id params", "error", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id params",
		})
	}

	var body TaskRequest
	if err := c.BodyParser(&body); err != nil {
		log.Error("failed to parse request body", "error", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	task, err := FromTaskRequest(body)
	task.ID = int64(id)
	if err != nil {
		log.Error("failed to parse request body", "error", err)
	}
	err = services.UpdateTask(task)
	if errors.Is(err, sql.ErrNoRows) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "task not found",
		})
	} else if err != nil {
		log.Error("failed to update task", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to update task",
		})
	}

	log.Info("task successfully updated with", "id", task.ID)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "task successfully updated",
	})
}

func DeleteTaskHandler(c *fiber.Ctx) error {
	log := getLogger(c)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Error("failed to get id params")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id params",
		})
	}

	err = services.DeleteTask(int64(id))
	if errors.Is(err, sql.ErrNoRows) {
		log.Error("task no found", "id", int64(id))
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "task not found",
		})
	} else if err != nil {
		log.Error("failed to delete task", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
	}

	log.Info("task successfully deleted with", "id", int64(id))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "task successfully deleted",
	})
}

func CompleteTaskHandler(c *fiber.Ctx) error {
	log := getLogger(c)

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Error("failed to get id params")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id params",
		})
	}

	if err = services.CompleteTask(int64(id)); errors.Is(err, sql.ErrNoRows) {
		log.Error("task no found", "id", int64(id))
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "task not found",
		})
	} else if err != nil {
		log.Error("failed to complete task", "error", err)
	}

	log.Info("task successfully completed with", "id", int64(id))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "task successfully completed",
	})
}
