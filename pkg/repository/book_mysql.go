package repository

import (
	"fmt"

	"github.com/aburtasov/books"
	"github.com/jmoiron/sqlx"
)

type BookMysql struct {
	db *sqlx.DB
}

func NewBookMysql(db *sqlx.DB) *BookMysql {
	return &BookMysql{db: db}
}

func (r *BookMysql) GetBooks(author books.Author) ([]string, error) {
	var books []string
	var id int

	query1 := fmt.Sprintf("SELECT id FROM %s WHERE first_name=%s AND second_name=%s", authorsTable, author.FirstName, author.SecondName)
	err := r.db.Get(&id, query1)
	if err != nil {
		return nil, err
	}

	query2 := fmt.Sprintf("SELECT bt.title FROM %s bt WHERE author_id=%d", booksTable, id)
	if err := r.db.Select(&books, query2); err != nil {
		return nil, err
	}

	return books, nil
}
