package storage

import (
	"online_lib_api"

	"github.com/jmoiron/sqlx"
)

type Book interface {
	GetContentBook(book_id string) string
}

type Authorization interface {
	CreateUser(user online_lib_api.User) (int, error)
	GetUser(username, password string) (online_lib_api.User, error)
}

type Storage struct {
	Authorization
	Book
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		Authorization: NewAuthPostgres(db),
		Book: NewBookPostgres(db),
	}
}