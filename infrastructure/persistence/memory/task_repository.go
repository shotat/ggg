package memory

import (
	"context"

	"github.com/shotat/ggg/domain/model/task"
	"github.com/shotat/ggg/domain/repository"
)

type TaskRepository struct{}

func NewTaskRepository() repository.TaskRepository {
	return &TaskRepository{}
}

func (r *TaskRepository) FindOne(ctx context.Context, ID int) (*task.Task, error) {
	t := task.NewTask(ID, "sample", "sample")
	return t, nil
}
