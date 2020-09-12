package usecase

import (
	"context"

	"github.com/tetsuzawa/godddtodo/domain/model"
	"github.com/tetsuzawa/godddtodo/domain/repository"
)

type TodoUseCase interface {
	Show(ctx context.Context, todoID int64) (*model.Todo, error)
	Create(ctx context.Context, title, text string) (*model.Todo, error)
}

type todoUseCase struct {
	repository.TodoRepository
}

func NewTodoUseCase(r repository.TodoRepository) TodoUseCase {
	return &todoUseCase{r}
}

func (u *todoUseCase) Show(ctx context.Context, todoID int64) (*model.Todo, error) {
	todo, err := u.TodoRepository.Find(ctx, todoID)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (u *todoUseCase) Create(ctx context.Context, title, text string) (*model.Todo, error) {
	todo := &model.Todo{
		Title: title,
		Text:  text,
	}
	todo, err := u.TodoRepository.Create(ctx, todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
