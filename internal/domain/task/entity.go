package task

const (
	EntityType = "task"
)

// Task ...
type Task struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// New func is a constructor for the Task
func New() *Task {
	return &Task{}
}
