package service

import "github.com/Red-373/todo/pkg/repository"

type Authorization struct {
}
type TodoList struct {
}
type TodoItem struct {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
