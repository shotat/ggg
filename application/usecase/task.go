package usecase

import (
	"context"

	"github.com/shotat/ggg/domain/model"
	"github.com/shotat/ggg/domain/repository"
)

type TaskUseCase interface {
	GetTask(ctx context.Context, id int) (*model.Task, error)
}

// impl
func NewTaskUseCase(t repository.TaskRepository) TaskUseCase {
	return &taskUseCase{t}
}

type taskUseCase struct {
	t repository.TaskRepository
}

func (u *taskUseCase) GetTask(ctx context.Context, id int) (*model.Task, error) {
	t, err := u.t.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}
	return t, nil
}
