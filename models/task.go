package models

type Task struct {
	ID          string
	Name        string
	IsCompleted bool
}

func (t Task) Hello() string {
	return "Hello"
}
