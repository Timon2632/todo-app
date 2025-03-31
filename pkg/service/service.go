package service

import (
	"github.com/Timon2632/todo-app"
	"github.com/Timon2632/todo-app/pkg/repositoty"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repositoty.Repositoty) *Service {
	return &Service{
		Authorization: NewAuthServise(repos.Authorzation),
	}
}
