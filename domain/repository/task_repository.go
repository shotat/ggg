package repository

import (
	"context"

	"github.com/shotat/ggg/domain/model/task"
)

type TaskRepository interface {
	FindOne(ctx context.Context, ID int) (*task.Task, error)
}
