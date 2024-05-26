package storage

import (
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