package services

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"todo/internal/db"
	database "todo/internal/db/gen"
	"todo/internal/models"
	"todo/internal/utils"
)

//func getTask(id int64) (models.Task, error) {
//	query := database.New(db.DB)
//	taskDb, err := query.GetTask(context.Background(), id)
//	task := models.Task{
//		ID:          taskDb.ID,
//		Title:       utils.ToNormalType(taskDb.Title).(string),
//		Description: utils.ToNullType(taskDb.Description).(string),
//		DueDate:
//	}
//}

func CreateTask(task models.Task) (models.Task, error) {

	tx, err := db.DB.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return models.Task{}, fmt.Errorf("failed to create transaction: %w", err)
	}

	query := database.New(tx)
	if err := query.CreateTask(context.Background(), database.CreateTaskParams{
		Title:       task.Title,
		Description: utils.ToNullType(task.Description).(sql.NullString),
		DueDate:     utils.ToNullType(task.DueDate).(sql.NullString),
		Overdue:     utils.ToNullType(task.Overdue).(sql.NullInt64),
		Completed:   utils.ToNullType(task.Completed).(sql.NullInt64),
	}); err != nil {
		_ = tx.Rollback()
		return models.Task{}, fmt.Errorf("failed to create task: %w", err)
	}

	lastID, err := query.GetLastID(context.Background())
	if err != nil {
		_ = tx.Rollback()
		return models.Task{}, fmt.Errorf("failed to get last id: %w", err)
	}

	fmt.Println("last id", lastID)

	if err := tx.Commit(); err != nil {
		return models.Task{}, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return models.Task{}, nil
}

func OverdueChecker(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		// TODO: проверка задач на просрочку
	}
}
