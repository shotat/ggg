package repository

import "github.com/shotat/ggg/domain/model"

type TaskRepository interface {
	FindOne(ID int) (*model.Task, error)
}
