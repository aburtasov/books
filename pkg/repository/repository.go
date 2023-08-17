package repository

import (
	"database/sql"

	"github.com/aburtasov/books"
	_ "github.com/go-sql-driver/mysql"
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

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Book: NewBookMysql(db),
	}
}
