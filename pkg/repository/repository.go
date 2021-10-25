package repository

import (
	"github.com/Red-373/todo"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList struct {
}

type TodoItem struct {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
