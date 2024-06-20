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

func (st *BookPostgres) GetBooks(filters map[string]interface{}) ([]online_lib_api.Book, error) {
	var books []online_lib_api.Book
	var book online_lib_api.Book
	var books_id []int
	var err error

	filters["genres"] = int(filters["genres"].(float64))

	if filters["genres"] == -1 {
		if filters["filter"] == "new" {
			err = st.db.Select(&books, "SELECT * FROM books ORDER BY id DESC")
			if err != nil {
				fmt.Println(err)
			}
		} else if filters["filter"] == "popular" {
			err = st.db.Select(&books, "SELECT * FROM books ORDER BY likes DESC")
			if err != nil {
				fmt.Println(err)
			}
		}
	} else {
		if filters["filter"] == "new" {
			err = st.db.Select(&books_id, "SELECT book_id FROM books_genres WHERE genre_id=$1 ORDER BY book_id DESC", filters["genres"])
			if err != nil {
				fmt.Println(err)
			}
			for i := 0; i < len(books_id); i++ {
				err = st.db.Get(&book, "SELECT * FROM books WHERE id=$1", books_id[i])
				if err != nil {
					fmt.Println(err)
				}
				books = append(books, book)
			}
		} else if filters["filter"] == "popular" {
			err = st.db.Select(&books_id, "SELECT book_id FROM books_genres WHERE genre_id=$1 ORDER BY book_id DESC", filters["genres"])
			if err != nil {
				fmt.Println(err)
			}
			for i := 0; i < len(books_id); i++ {
				err = st.db.Get(&book, "SELECT * FROM books WHERE id=$1", books_id[i])
				if err != nil {
					fmt.Println(err)
				}
				books = append(books, book)
			}
		}
	}
	return books, err
}

func (st *BookPostgres) GetAboutBook(book_name_id string) (online_lib_api.Book, error) {
	var about_book online_lib_api.Book

	// var book_genres []online_lib_api.Genre
	var book_id int
	err := st.db.Get(&book_id, "SELECT id FROM books WHERE id_name=$1", book_name_id)
	if err != nil {
		fmt.Println(err)
	}
	
	var genres_id []int
	err = st.db.Select(&genres_id, "SELECT genre_id FROM books_genres WHERE book_id=$1", book_id)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(genres_id); i++ {
		var genre online_lib_api.Genre
		err = st.db.Get(&genre, "SELECT * FROM genres WHERE id=$1", genres_id[i])
		if err != nil {
			fmt.Println(err)
		}
		about_book.BookGenres = append(about_book.BookGenres, genre)
	}
	
	query := "SELECT * FROM books WHERE id_name=$1"
	err = st.db.Get(&about_book, query, book_name_id)

	return about_book, err
}

func (st *BookPostgres) GetPopularGenres() ([]online_lib_api.Genre, error) {
	var popular_genres []online_lib_api.Genre
	query := "SELECT * FROM genres"
	err := st.db.Select(&popular_genres, query)
	return popular_genres, err
}

func (st *BookPostgres) SearchBooks(param string) ([]online_lib_api.Book, error) {
	var books []online_lib_api.Book
	query := "SELECT * FROM books WHERE book_name LIKE '%" + param + "%' OR book_author LIKE '%" + param + "%'"
	err := st.db.Select(&books, query)
	return books, err
}