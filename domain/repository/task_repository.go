package repository

import (
	"context"

	"github.com/shotat/ggg/domain/model"
)

type TaskRepository interface {
	FindOne(ctx context.Context, ID int) (*model.Task, error)
}
