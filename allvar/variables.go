package allvar

import "time"

type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Priority    string     `json:"priority"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

type TaskManager struct {
	Tasks []Task `json:"tasks"`
	NextID int `json:"next_id"`
	FilePath string `json:"-"`
}

var (
	TaskFile string
	Priority string
	ShowAll bool
	TaskManagerVar *TaskManager
)

func Truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}