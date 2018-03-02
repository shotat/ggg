package repository

import "github.com/shotat/ggg/domain/model/task"

type TaskRepository interface {
	FindOne(ID int) (*task.Task, error)
}
