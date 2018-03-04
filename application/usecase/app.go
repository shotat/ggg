package usecase

type AppUseCase interface {
	TaskUseCase
	// XxxUseCase
	// YyyUseCase
	// ...
}

type appUseCase struct {
	TaskUseCase
}

func NewAppUseCase(taskUseCase TaskUseCase) AppUseCase {
	return &appUseCase{
		TaskUseCase: taskUseCase,
	}
}
