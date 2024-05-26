package service

import (
	"online_lib_api"
	"online_lib_api/internal/storage"
)

type Book interface {
	GetContentBook(book_name_id string) (string, error)
}

type Authorization interface {
	CreateUser(user online_lib_api.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
	Book
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		Authorization: NewAuthService(storage),
		Book: NewBookService(storage),
	}
}