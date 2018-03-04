package inmem

import (
	"context"

	"github.com/shotat/ggg/domain/model"
	"github.com/shotat/ggg/domain/repository"
)

type TaskRepository struct{}

func NewTaskRepository() repository.TaskRepository {
	return &TaskRepository{}
}

func (r *TaskRepository) FindOne(ctx context.Context, ID int) (*model.Task, error) {
	t := model.NewTask(ID, "sample", "sample")
	return t, nil
}
