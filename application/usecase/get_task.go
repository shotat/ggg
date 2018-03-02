package usecase

import (
	"fmt"

	"github.com/shotat/ggg/application/input"
	"github.com/shotat/ggg/application/output"
	"github.com/shotat/ggg/domain/repository"
)

func NewGetTaskUseCase(t repository.TaskRepository) *GetTaskUseCase {
	return &GetTaskUseCase{t}
}

type GetTaskUseCase struct {
	t repository.TaskRepository
}

func (u *GetTaskUseCase) Execute(q *input.Query) (*output.TaskDetail, error) {
	query := q.Q
	fmt.Println(query)
	t, err := u.t.FindOne(10)
	if err != nil {
		return nil, err
	}
	return &output.TaskDetail{t.Name}, nil
}
