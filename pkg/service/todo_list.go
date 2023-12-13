package service

import (
	"github.com/Nothinglikethat/todo-app"
	"github.com/Nothinglikethat/todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoLostService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list todo.TodoLists) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]todo.TodoLists, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listId int) (todo.TodoLists, error) {
	return s.repo.GetById(userId, listId)
}
