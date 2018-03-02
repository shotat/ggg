package handler

import (
	"io"
	"net/http"

	"github.com/shotat/ggg/application/input"
	"github.com/shotat/ggg/application/usecase"
)

type TaskHandler struct{}

func (h *TaskHandler) Index(w http.ResponseWriter, r *http.Request) {
	u := usecase.NewGetTaskUseCase()
	i := &input.Query{"TODO get params from *r"}
	if o, err := u.Execute(i); err != nil {
		w.WriteHeader(500)
	} else {
		io.WriteString(w, o.Name)
	}
}
