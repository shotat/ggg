package server

import (
	"net/http"

	"github.com/shotat/ggg/interfaces/api/server/router"
)

type Server interface {
	Serve() error
}

type server struct {
	router.Router
}

func NewServer(r router.Router) Server {
	return &server{r}
}

func (s *server) Serve() error {
	return http.ListenAndServe(":8080", s.Router)
}
