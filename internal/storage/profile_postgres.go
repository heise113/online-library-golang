package storage

import (
	"fmt"
	"online_lib_api"

	"github.com/jmoiron/sqlx"
)

type ProfilePostgres struct {
	db *sqlx.DB
}

func NewProfilePostgres(db *sqlx.DB) *ProfilePostgres {
	return &ProfilePostgres{db: db}
}

func (st *ProfilePostgres) GetProfileData(user_id int) (online_lib_api.Profile, error) {
	var profile_data online_lib_api.Profile
	query := "SELECT name, username FROM users WHERE id=$1"
	err := st.db.Get(&profile_data, query, user_id)
	if err != nil {
		fmt.Println(err)
	}

	var my_books_id []int
	query = "SELECT book_id FROM users_books WHERE user_id=$1"
	err = st.db.Select(&my_books_id, query, user_id)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(my_books_id); i++ {
		var book online_lib_api.Book
		err = st.db.Get(&book, "SELECT * FROM books WHERE id=$1", my_books_id[i])
		if err != nil {
			fmt.Println(err)
		}
		profile_data.Mybooks = append(profile_data.Mybooks, book)
		fmt.Println("my books: ", profile_data.Mybooks)
	}

	fmt.Println("profile_data: ", profile_data)
	return profile_data, err
}

func (st *ProfilePostgres) AddBook(user_id int, book_id int) error {
	_, err := st.db.Exec("INSERT INTO users_books (user_id, book_id) VALUES ($1, $2)", user_id, book_id)
	return err
}

func (st *ProfilePostgres) DeleteBook(user_id int, book_id int) error {
	_, err := st.db.Exec("DELETE FROM users_books WHERE user_id=$1 AND book_id=$2", user_id, book_id)
	return err
}
