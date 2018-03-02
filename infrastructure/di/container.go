package di

import (
	"context"

	"github.com/shotat/ggg/domain/repository"
	"github.com/shotat/ggg/infrastructure/persistence/memory"
)

func InjectTaskRepository(ctx context.Context) repository.TaskRepository {
	return memory.NewInMemTaskRepository(ctx)
}

func InjectTaskRepository(ctx context.Context) repository.TaskRepository {
	return memory.NewInMemTaskRepository(ctx)
}
