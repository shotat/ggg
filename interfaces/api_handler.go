package interfaces

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shotat/ggg/interfaces/api"
)

type ApiHandler struct {
	handler http.Handler
}

func NewApiHandler() *ApiHandler {
	m := mux.NewRouter()
	t := api.NewTaskHandler()
	m.HandleFunc("/", t.Index)
	return &ApiHandler{m}
}

func (h *ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
}
