package repository

import (
	"github.com/aburtasov/books"
	"github.com/jmoiron/sqlx"
)

type Book interface {
	GetBooks(author books.Author) ([]string, error)
}

type Author interface {
}

type Repository struct {
	Book
	Author
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Book: NewBookMysql(db),
	}
}
