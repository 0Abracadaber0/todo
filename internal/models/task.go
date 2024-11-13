package models

import (
	"encoding/json"
	"time"
)

const DateFormat = "2006-01-02"

type CustomDate struct {
	time.Time
}

func (d *CustomDate) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	t, err := time.Parse(DateFormat, s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

func (d *CustomDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format(DateFormat))
}

type Task struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DueDate     CustomDate `json:"due_date"`
	Overdue     bool       `json:"overdue"`
	Completed   bool       `json:"completed"`
}
