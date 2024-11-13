package services

import (
	"database/sql"
	"log/slog"
	"time"
	"todo/internal/models"
)

type TaskService struct {
	db *sql.DB
}

func NewTaskService(db *sql.DB) *TaskService {
	return &TaskService{db: db}
}

func (s *TaskService) OverdueChecker(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		// TODO: проверка задач на просрочку
	}
}

func (s *TaskService) CreateTask(log *slog.Logger, task models.Task) (models.Task, error) {
	log.Debug("CreateTask")
	return models.Task{}, nil
}
