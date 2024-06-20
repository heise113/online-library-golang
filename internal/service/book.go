package service

import (
	"fmt"
	"io"
	"online_lib_api"
	"online_lib_api/internal/storage"
	"os"
)

type BookService struct {
	storage storage.Book
}

func NewBookService(storage storage.Book) *BookService {
	return &BookService{storage: storage}
}

func (s *BookService) GetContentBook(book_name_id string) (string, error) {
	file, err := os.Open(fmt.Sprintf("%s.html", book_name_id))
	if err != nil{
		fmt.Println(err) 
		return "", err
	}
	defer file.Close() 
		
	data := make([]byte, 64)
	var book string
		
	for{
		n, err := file.Read(data)
		if err == io.EOF{   // если конец файла
			break           // выходим из цикла
		}
		book += string(data[:n])
	}
	return book, nil
}

func (s *BookService) GetBooks(filters map[string]interface{}) ([]online_lib_api.Book, error) {
	all_books, err := s.storage.GetBooks(filters)
	return all_books, err
}

func (s *BookService) GetAboutBook(book_name_id string) (online_lib_api.Book, error) {
	about_book, err := s.storage.GetAboutBook(book_name_id)
	return about_book, err
}

func (s *BookService) GetPopularGenres() ([]online_lib_api.Genre, error) {
	popular_genres, err := s.storage.GetPopularGenres()
	return popular_genres, err
}

func (s *BookService) SearchBooks(param string) ([]online_lib_api.Book, error) {
	books, err := s.storage.SearchBooks(param)
	return books, err
}