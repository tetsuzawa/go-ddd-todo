package repository

import (
	"context"

	"github.com/tetsuzawa/godddtodo/domain/model"
)

type TodoRepository interface {
	Find(ctx context.Context, todoID int64) (*model.Todo, error)
	Create(ctx context.Context, todo *model.Todo) (*model.Todo, error)
}
