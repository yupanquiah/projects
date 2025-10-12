package tasks

import "time"

type TaskStatus string

const (
	STATUS_TODO        TaskStatus = "todo"
	STATUS_IN_PROGRESS TaskStatus = "in-progress"
	STATUS_DONE        TaskStatus = "done"
)

type Task struct {
	ID          int64      `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"` // "todo", "in-progress", "done"
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func NewTask(id int64, description string) *Task {
	return &Task{
		ID:          id,
		Description: description,
		Status:      STATUS_TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
