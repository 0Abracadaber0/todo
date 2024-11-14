package models

import (
	"time"
)

const DateFormat = "2006-01-02"

type CustomDate struct {
	time.Time
}

type Task struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DueDate     CustomDate `json:"due_date"`
	Overdue     bool       `json:"overdue"`
	Completed   bool       `json:"completed"`
}
