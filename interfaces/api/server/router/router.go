package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shotat/ggg/application/usecase"
)

type Router interface {
	http.Handler
}

func NewRouter(u usecase.AppUseCase) Router {
	r := mux.NewRouter()
	r.HandleFunc("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(400)
			return
		}
		t, err := u.GetTask(ctx, id)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		json.NewEncoder(w).Encode(t)
	})
	return r
}
