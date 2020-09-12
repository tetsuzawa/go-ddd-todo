package registry

import (
	"database/sql"

	"github.com/tetsuzawa/godddtodo/application/usecase"
	"github.com/tetsuzawa/godddtodo/domain/repository"
	"github.com/tetsuzawa/godddtodo/infrastructure/persistence/datastore"
	"github.com/tetsuzawa/godddtodo/interfaces/api/server/handler"
)

type Registry interface {
	NewTodoRepository() repository.TodoRepository
	NewTodoUseCase() usecase.TodoUseCase
	NewAppHandler() handler.AppHandler
}

type registryImpl struct {
	db             *sql.DB
	todoRepository repository.TodoRepository
	todoUseCase    usecase.TodoUseCase
	appHandler     handler.AppHandler
}

func NewRegistry(db *sql.DB) Registry {
	r := &registryImpl{db:db}
	r.todoRepository = r.NewTodoRepository()
	r.todoUseCase = r.NewTodoUseCase()
	r.appHandler = r.NewAppHandler()
	return r
}

func (r *registryImpl) NewTodoRepository() repository.TodoRepository {
	if r.todoRepository == nil {
		r.todoRepository = datastore.NewTodoRepository(r.db)
	}
	return r.todoRepository
}

func (r *registryImpl) NewTodoUseCase() usecase.TodoUseCase {
	if r.todoUseCase == nil {
		r.todoUseCase = usecase.NewTodoUseCase(r.todoRepository)
	}
	return r.todoUseCase
}

func (r *registryImpl) NewAppHandler() handler.AppHandler {
	if r.appHandler == nil{
		r.appHandler = handler.NewAppHandler(r.todoUseCase)
	}
	return r.appHandler
}
