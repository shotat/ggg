package input

func NewGetTask(ID int) *GetTask {
	return &GetTask{ID}
}

type GetTask struct {
	ID int
}
