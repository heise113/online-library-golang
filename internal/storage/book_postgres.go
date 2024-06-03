package storage

import (
	"fmt"
	"online_lib_api"

	"github.com/jmoiron/sqlx"
)

type BookPostgres struct {
	db *sqlx.DB
}

func NewBookPostgres(db *sqlx.DB) *BookPostgres {
	return &BookPostgres{db: db}
}

func (st *BookPostgres) GetContentBook(book_id string) string {
	return "Она пришла под утро..."
}

func (st *BookPostgres) GetAllBooks() ([]online_lib_api.Book, error) {
	var all_books []online_lib_api.Book
	query := "SELECT * FROM books"
	err := st.db.Select(&all_books, query)
	return all_books, err
}

func (st *BookPostgres) GetAboutBook(book_name_id string) (online_lib_api.Book, error) {
	var about_book online_lib_api.Book

	// var book_genres []online_lib_api.Genre
	var book_id int
	err := st.db.Get(&book_id, "SELECT id FROM books WHERE id_name=$1", book_name_id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("book_id", book_id)
	
	var genres_id []int
	err = st.db.Select(&genres_id, "SELECT genre_id FROM books_genres WHERE book_id=$1", book_id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("genres_id", genres_id)

	for i := 0; i < len(genres_id); i++ {
		var genre online_lib_api.Genre
		err = st.db.Get(&genre, "SELECT * FROM genres WHERE id=$1", genres_id[i])
		if err != nil {
			fmt.Println(err)
		}
		about_book.BookGenres = append(about_book.BookGenres, genre)
		fmt.Println("bg: ", about_book.BookGenres)
	}
	
	query := "SELECT * FROM books WHERE id_name=$1"
	err = st.db.Get(&about_book, query, book_name_id)
	fmt.Println("about_book: ", about_book)

	return about_book, err
}

func (st *BookPostgres) GetPopularGenres() ([]online_lib_api.Genre, error) {
	var popular_genres []online_lib_api.Genre
	query := "SELECT * FROM genres"
	err := st.db.Select(&popular_genres, query)
	return popular_genres, err
}