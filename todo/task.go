package todo

import "time"

type Task struct {
	Title       string
	Description string
	Completed   bool

	CreatedAt   time.Time
	CompletedAT *time.Time
}

func NewTask(title string, description string) Task {
	return Task{
		Title:       title,
		Description: description,
		Completed:   false,

		CreatedAt:   time.Now(),
		CompletedAT: nil,
	}
}

func (t *Task) Complete() {
	completTime := time.Now()
	t.Completed = true
	t.CompletedAT = &completTime
}

func (t *Task) Uncomplete() {
	t.Completed = false
	t.CompletedAT = nil
}
