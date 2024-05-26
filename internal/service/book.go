package service

import (
	"online_lib_api/internal/storage"
	"os"
	"io"
	"fmt"
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