package service

import "context"

type TodoService interface {
	DoSomething(ctx context.Context, foo int) error
}
