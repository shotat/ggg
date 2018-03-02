package usecase

import (
	"fmt"

	"github.com/shotat/ggg/application/input"
	"github.com/shotat/ggg/application/output"
	"github.com/shotat/ggg/infrastructure/persistence"
)

func NewGetTaskUseCase() *GetTaskUseCase {
	return &GetTaskUseCase{}
}

type GetTaskUseCase struct{}

func (u *GetTaskUseCase) Execute(q *input.Query) (*output.TaskDetail, error) {
	query := q.Q
	fmt.Println(query)
	// FIXME DI
	r := persistence.NewInMemTaskRepository()
	t, err := r.FindOne(10)
	if err != nil {
		return nil, err
	}
	return &output.TaskDetail{t.Name}, nil
}
