package api

import (
	"encoding/json"
	"net/http"

	"github.com/shotat/ggg/application/input"
	"github.com/shotat/ggg/application/usecase"
	"github.com/shotat/ggg/infrastructure/persistence/memory"
)

type TaskHandler struct{}

func NewTaskHandler() *TaskHandler {
	return &TaskHandler{}
}

func (h *TaskHandler) Index(w http.ResponseWriter, r *http.Request) {
	repo := memory.NewInMemTaskRepository(r.Context())

	u := usecase.NewGetTask(repo)
	i := &input.GetTask{100}
	if o, err := u.Execute(i); err != nil {
		w.WriteHeader(500)
	} else {
		json.NewEncoder(w).Encode(o)
	}
}
