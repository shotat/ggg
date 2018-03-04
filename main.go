package main

import (
	"github.com/shotat/ggg/application/usecase"
	"github.com/shotat/ggg/infrastructure/persistence/inmem"
	"github.com/shotat/ggg/interfaces/api/server"
	"github.com/shotat/ggg/interfaces/api/server/router"
)

func main() {
	//
	taskRepository := inmem.NewTaskRepository()
	taskUseCase := usecase.NewTaskUseCase(taskRepository)

	appUseCase := usecase.NewAppUseCase(taskUseCase)

	router := router.NewRouter(appUseCase)
	s := server.NewServer(router)
	s.Serve()
}
