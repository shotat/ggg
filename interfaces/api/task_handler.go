package api

import (
	"encoding/json"
	"net/http"

	"github.com/shotat/ggg/application/input"
	"github.com/shotat/ggg/application/usecase"
)

type TaskHandler struct {
	*usecase.GetTask
}

func NewTaskHandler(uc *usecase.GetTask) *TaskHandler {
	return &TaskHandler{uc}
}

func (h *TaskHandler) Index(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	i := input.NewGetTask(100)
	if o, err := h.GetTask.Execute(ctx, i); err != nil {
		w.WriteHeader(500)
	} else {
		json.NewEncoder(w).Encode(o)
	}
}
