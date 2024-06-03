package storage

import (
	"fmt"
	"online_lib_api"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (st *AuthPostgres) CreateUser(user online_lib_api.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO users (name, username, password_hash) values ($1, $2, $3) RETURNING id")

	row := st.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (st *AuthPostgres) GetUser(username, password string) (online_lib_api.User, error) {
	var user online_lib_api.User
	query := fmt.Sprintf("SELECT id FROM users WHERE username=$1 AND password_hash=$2")
	err := st.db.Get(&user, query, username, password)
	return user, err
}