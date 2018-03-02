package interfaces

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shotat/ggg/interfaces/handler"
)

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m := mux.NewRouter()
	t := &handler.TaskHandler{}
	m.HandleFunc("/", t.Index)
	m.ServeHTTP(w, r)
}
