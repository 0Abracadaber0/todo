package handlers

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
	"todo/internal/models"
)

func TestFromTaskRequest(t *testing.T) {
	tests := []struct {
		name     string
		req      TaskRequest
		expected models.Task
		err      bool
	}{
		{
			name: "Valid TaskRequest with valid due date",
			req: TaskRequest{
				Title:       "Test Task",
				Description: "This is a test task",
				DueDate:     "2024-11-14",
			},
			expected: models.Task{
				Title:       "Test Task",
				Description: "This is a test task",
				DueDate:     models.CustomDate{Time: time.Date(2024, 11, 14, 0, 0, 0, 0, time.UTC)},
				Overdue:     false,
				Completed:   false,
			},
			err: false,
		},
		{
			name: "Invalid TaskRequest with invalid date format",
			req: TaskRequest{
				Title:       "Test Task",
				Description: "This is a test task",
				DueDate:     "invalid-date",
			},
			expected: models.Task{},
			err:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task, err := FromTaskRequest(tt.req)
			if tt.err {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, task)
			}
		})
	}
}

func TestToTaskResponse(t *testing.T) {
	tests := []struct {
		name     string
		task     models.Task
		expected TaskResponse
	}{
		{
			name: "Task with empty DueDate",
			task: models.Task{
				ID:          1,
				Title:       "Test Task",
				Description: "This is a test task",
				Overdue:     false,
				Completed:   false,
			},
			expected: TaskResponse{
				ID:          1,
				Title:       "Test Task",
				Description: "This is a test task",
				DueDate:     "",
				Overdue:     false,
				Completed:   false,
			},
		},
		{
			name: "Task with filled DueDate",
			task: models.Task{
				ID:          1,
				Title:       "Test Task",
				Description: "This is a test task",
				DueDate:     models.CustomDate{Time: time.Date(2024, 11, 14, 0, 0, 0, 0, time.UTC)},
				Overdue:     false,
				Completed:   false,
			},
			expected: TaskResponse{
				ID:          1,
				Title:       "Test Task",
				Description: "This is a test task",
				DueDate:     "2024-11-14",
				Overdue:     false,
				Completed:   false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := ToTaskResponse(tt.task)
			assert.Equal(t, tt.expected, resp)
		})
	}
}
