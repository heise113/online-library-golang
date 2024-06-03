package online_lib_api

type Book struct {
	Id int `json:"id" db:"id"`
	IdName string `json:"id_name" db:"id_name"`
	BookName string `json:"book_name" db:"book_name"`
	BookAuthor string `json:"book_author" db:"book_author"`
	BookImage string `json:"book_image" db:"book_image"`
	BookGenres []Genre `json:"book_genres" db:"genre"`
	BookDescription string `json:"book_description" db:"book_description"`
}