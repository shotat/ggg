package usecase

import (
	"context"

	"github.com/shotat/ggg/application/input"
	"github.com/shotat/ggg/application/output"
	"github.com/shotat/ggg/domain/repository"
)

func NewGetTask(t repository.TaskRepository) *GetTask {
	return &GetTask{t}
}

type GetTask struct {
	t repository.TaskRepository
}

func (u *GetTask) Execute(ctx context.Context, i *input.GetTask) (*output.GetTask, error) {
	id := i.ID
	t, err := u.t.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}
	return &output.GetTask{t}, nil
}
