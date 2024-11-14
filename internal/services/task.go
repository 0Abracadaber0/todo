package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
	"todo/internal/db"
	database "todo/internal/db/gen"
	"todo/internal/models"
	"todo/internal/utils"
)

func getTask(tx *sql.Tx, id int64) (models.Task, error) {
	query := database.New(tx)
	taskDb, err := query.GetTask(context.Background(), id)
	if err != nil {
		return models.Task{}, err
	}

	task := models.Task{
		ID:          taskDb.ID,
		Title:       taskDb.Title,
		Description: utils.ToNormalType(taskDb.Description).(string),
		DueDate:     utils.ToNormalType(taskDb.DueDate).(models.CustomDate),
		Overdue:     utils.ToNormalType(taskDb.Overdue).(bool),
		Completed:   utils.ToNormalType(taskDb.Completed).(bool),
	}
	return task, nil
}

func CreateTask(task models.Task) (models.Task, error) {

	tx, err := db.DB.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return models.Task{}, fmt.Errorf("failed to create transaction: %w", err)
	}

	query := database.New(tx)
	id, err := query.CreateTask(context.Background(), database.CreateTaskParams{
		Title:       task.Title,
		Description: utils.ToNullType(task.Description).(sql.NullString),
		DueDate:     utils.ToNullType(task.DueDate).(sql.NullString),
		Overdue:     utils.ToNullType(task.Overdue).(sql.NullInt64),
		Completed:   utils.ToNullType(task.Completed).(sql.NullInt64),
	})
	if err != nil {
		_ = tx.Rollback()
		return models.Task{}, fmt.Errorf("failed to create task: %w", err)
	}

	task, err = getTask(tx, id)
	if err != nil {
		_ = tx.Rollback()
		return models.Task{}, fmt.Errorf("failed to get task: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return task, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return task, nil
}

func GetTasks() ([]models.Task, error) {
	query := database.New(db.DB)
	tasksDb, err := query.GetTasks(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks: %w", err)
	}
	tasks := make([]models.Task, len(tasksDb))

	for i, taskDb := range tasksDb {
		tasks[i] = models.Task{
			ID:          taskDb.ID,
			Title:       taskDb.Title,
			Description: utils.ToNormalType(taskDb.Description).(string),
			DueDate:     utils.ToNormalType(taskDb.DueDate).(models.CustomDate),
			Overdue:     utils.ToNormalType(taskDb.Overdue).(bool),
			Completed:   utils.ToNormalType(taskDb.Completed).(bool),
		}
	}

	return tasks, nil
}

func DeleteTask(id int64) error {
	tx, err := db.DB.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return fmt.Errorf("failed to create transaction: %w", err)
	}

	query := database.New(tx)
	_, err = query.GetTask(context.Background(), id)
	if errors.Is(err, sql.ErrNoRows) {
		_ = tx.Rollback()
		return err
	} else if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("failed to get task: %w", err)
	}

	err = query.DeleteTask(context.Background(), id)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("failed to delete task: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func OverdueChecker(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		// TODO: проверка задач на просрочку
	}
}
