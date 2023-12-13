package service

import (
	"github.com/Nothinglikethat/todo-app"
	"github.com/Nothinglikethat/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoLists) (int, error)
	GetAll(userId int) ([]todo.TodoLists, error)
	GetById(userIdm, listId int) (todo.TodoLists, error)
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoLostService(repos.TodoList),
	}
}
