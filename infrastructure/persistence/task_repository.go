package persistence

import (
	"context"
	"fmt"
	"github.com/shotat/ggg/domain/model"
	"github.com/shotat/ggg/domain/repository"
)

type InMemTaskRepository struct {
	ctx context.Context
}

func NewInMemTaskRepository(ctx context.Context) repository.TaskRepository {
	return &InMemTaskRepository{ctx}
}

func (r *InMemTaskRepository) FindOne(ID int) (*model.Task, error) {
	t := model.NewTask(ID, "sample", "sample")
	fmt.Println(r.ctx)
	return t, nil
}
