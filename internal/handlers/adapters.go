package handlers

import (
	"time"
	"todo/internal/models"
)

type TaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
}

type TaskResponse struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Overdue     bool   `json:"overdue"`
	Completed   bool   `json:"completed"`
}

func FromTaskRequest(req TaskRequest) (models.Task, error) {
	var overdue models.CustomDate
	if req.DueDate == "" {
		overdue.Time = time.Time{}
	}
	var err error
	overdue.Time, err = time.Parse(models.DateFormat, req.DueDate)
	if err != nil {
		return models.Task{}, err
	}
	task := models.Task{
		Title:       req.Title,
		Description: req.Description,
		DueDate:     overdue,
		Overdue:     false,
		Completed:   false,
	}

	return task, nil
}

func ToTaskResponse(task models.Task) TaskResponse {
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
