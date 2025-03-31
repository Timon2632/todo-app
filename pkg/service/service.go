package service

import (
	"github.com/Timon2632/todo-app"
	"github.com/Timon2632/todo-app/pkg/repositoty"
)

type Authorzation interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorzation
	TodoList
	TodoItem
}

func NewService(repos *repositoty.Repositoty) *Service {
	return &Service{
		Authorzation: NewAuthServise(repos.Authorzation),
	}
}
