package service

import (
	"online_lib_api"
	"online_lib_api/internal/storage"
)

type Book interface {
	GetContentBook(book_name_id string) (string, error)
	GetBooks(filters map[string]interface{}) ([]online_lib_api.Book, error)
	GetAboutBook(book_name_id string) (online_lib_api.Book, error)
	GetPopularGenres() ([]online_lib_api.Genre, error)
	SearchBooks(param string) ([]online_lib_api.Book, error)
}

type Authorization interface {
	CreateUser(user online_lib_api.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Profile interface {
	GetProfileData(user_id int) (online_lib_api.Profile, error)
	AddBook(user_id int, book_id int) error
	DeleteBook(user_id int, book_id int) error
}

type Service struct {
	Authorization
	Book
	Profile
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		Authorization: NewAuthService(storage),
		Book: NewBookService(storage),
		Profile: NewProfileService(storage),
	}
}