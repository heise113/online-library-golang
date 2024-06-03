package storage

import (
	"online_lib_api"

	"github.com/jmoiron/sqlx"
)

type Book interface {
	GetContentBook(book_id string) string
	GetAllBooks() ([]online_lib_api.Book, error)
	GetAboutBook(book_name_id string) (online_lib_api.Book, error)
	GetPopularGenres() ([]online_lib_api.Genre, error)
}

type Authorization interface {
	CreateUser(user online_lib_api.User) (int, error)
	GetUser(username, password string) (online_lib_api.User, error)
}

type Profile interface {
	GetProfileData(user_id int) (online_lib_api.Profile, error)
	AddBook(user_id int, book_id int) error
	DeleteBook(user_id int, book_id int) error
}

type Storage struct {
	Authorization
	Book
	Profile
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		Authorization: NewAuthPostgres(db),
		Book: NewBookPostgres(db),
		Profile: NewProfilePostgres(db),
	}
}