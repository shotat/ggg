package interfaces

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shotat/ggg/application/usecase"
	"github.com/shotat/ggg/infrastructure/persistence/memory"
	"github.com/shotat/ggg/interfaces/api"
)

type API struct {
	handler http.Handler
}

func NewApi() *API {
	m := mux.NewRouter()
	// FIXME: di
	taskRepository := memory.NewTaskRepository()
	getTask := usecase.NewGetTask(taskRepository)
	t := api.NewTaskHandler(getTask)
	m.HandleFunc("/", t.Index)

	return &API{m}
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.handler.ServeHTTP(w, r)
}
