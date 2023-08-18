package repository

import (
	"database/sql"

	book "github.com/aburtasov/books"
	_ "github.com/go-sql-driver/mysql"
)

type Book interface {
	GetBooks(author book.Author) ([]book.Book, error)
	AddBook(book book.Book) (book.Responce, error)
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
