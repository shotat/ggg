package interfaces

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shotat/ggg/interfaces/handler"
)

type AppHandler struct {
	handler http.Handler
}

func NewAppHandler() *AppHandler {
	m := mux.NewRouter()
	t := handler.NewTaskHandler()
	m.HandleFunc("/", t.Index)
	return &AppHandler{m}
}

func (h *AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
}
