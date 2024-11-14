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
	if s == "" {
		d.Time = time.Time{}
		return nil
	}
	t, err := time.Parse(DateFormat, s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

func (d *CustomDate) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return json.Marshal("")
	}
	return json.Marshal(d.Time.Format(DateFormat))
}

type Task struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DueDate     CustomDate `json:"due_date"`
	Overdue     bool       `json:"overdue"`
	Completed   bool       `json:"completed"`
}
