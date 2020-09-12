package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tetsuzawa/godddtodo/application/usecase"
	"github.com/tetsuzawa/godddtodo/domain/model"
	"log"
	"net/http"
	"strconv"
)

type TodoHandler interface {
	ShowTodo(w http.ResponseWriter, r *http.Request) (int, interface{}, error)
	CreateTodo(w http.ResponseWriter, r *http.Request) (int, interface{}, error)
}

type todoHandler struct {
	useCase usecase.TodoUseCase
}

func NewTodoHandler(u usecase.TodoUseCase) TodoHandler {
	return &todoHandler{u}
}

func (h *todoHandler) ShowTodo(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	todoID, err := strconv.ParseInt(vars["todo_id"], 10, 64)
	if err != nil {
		log.Println(err)
		return http.StatusBadRequest, nil, errors.New("todo_id is invalid")
	}

	todo, err := h.useCase.Show(r.Context(), todoID)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	if todo == nil {
		return http.StatusNotFound, nil, errors.New("todo not found")
	}
	return http.StatusOK, todo, nil
}

func (h *todoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	todo := &model.Todo{}
	if err := json.NewDecoder(r.Body).Decode(todo); err != nil {
		fmt.Println(err)
		return http.StatusBadRequest, nil, errors.New("request body is invalid")
	}
	todo, err := h.useCase.Create(r.Context(), todo.Title, todo.Text)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusCreated, todo, nil
}
