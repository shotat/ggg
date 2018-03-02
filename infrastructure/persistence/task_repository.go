package persistence

import (
	"github.com/shotat/ggg/domain/model"
	"github.com/shotat/ggg/domain/repository"
)

type InMemTaskRepository struct {
}

func NewInMemTaskRepository() repository.TaskRepository {
	return &InMemTaskRepository{}
}

func (r *InMemTaskRepository) FindOne(ID int) (*model.Task, error) {
	t := model.NewTask(ID, "sample", "sample")
	return t, nil
}
