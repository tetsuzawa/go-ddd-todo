package datastore

import (
	"context"
	"database/sql"

	"github.com/tetsuzawa/godddtodo/domain/model"
	"github.com/tetsuzawa/godddtodo/domain/repository"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) repository.TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) Find(ctx context.Context, todoID int64) (*model.Todo, error) {
	todo := &model.Todo{}
	err := r.db.QueryRowContext(ctx, "SELECT id, title, text from todo WHERE id = ?", todoID).Scan(&todo.ID, &todo.Title, &todo.Text)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return todo, nil
	}
}

func (r *TodoRepository) Create(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	stmt, err := r.db.PrepareContext(ctx, "INSERT INTO todo (title, text) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}

	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	result, err := stmt.ExecContext(ctx, todo.Title, todo.Text)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	todo.ID = id
	return todo, nil
}
