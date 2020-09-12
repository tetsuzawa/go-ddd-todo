package handler

import "github.com/tetsuzawa/godddtodo/application/usecase"

type AppHandler interface {
	TodoHandler
}

type appHandler struct {
	TodoHandler
}

func NewAppHandler(u usecase.TodoUseCase) AppHandler {
	a := &appHandler{
		TodoHandler: NewTodoHandler(u),
	}
	return a
}
