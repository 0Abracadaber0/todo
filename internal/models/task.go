package models

import "time"

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Overdue     bool      `json:"overdue"`
	Completed   bool      `json:"completed"`
}
