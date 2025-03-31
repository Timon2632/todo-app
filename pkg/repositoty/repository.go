package repositoty

import (
	"github.com/Timon2632/todo-app"
	"github.com/jmoiron/sqlx"
)

type Authorzation interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repositoty struct {
	Authorzation
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repositoty {
	return &Repositoty{
		Authorzation: NewAuthPostgres(db),
	}
}
