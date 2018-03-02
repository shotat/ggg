package task

type Task struct {
	ID          int
	Name        string
	Description string
}

func NewTask(ID int, Name string, Description string) *Task {
	return &Task{ID, Name, Description}
}
