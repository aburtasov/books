package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	book "github.com/aburtasov/books"
)

type AuthorMysql struct {
	db *sql.DB
}

func NewAuthorMysql(db *sql.DB) *AuthorMysql {
	return &AuthorMysql{db: db}
}

func (r *AuthorMysql) GetAuthor(bookTitle string) (book.Author, error) {

	var id int
	var author book.Author

	query1 := fmt.Sprintf("SELECT author_id FROM %s WHERE title='%s'", booksTable, bookTitle)

	if err := r.db.QueryRow(query1).Scan(&id); err != nil {

		return book.Author{}, err
	}

	query2 := fmt.Sprintf("SELECT first_name,second_name FROM %s WHERE id='%d'", authorsTable, id)

	if err := r.db.QueryRow(query2).Scan(&author.FirstName, &author.SecondName); err != nil {

		return book.Author{}, err
	}

	return author, nil

}

func (r *AuthorMysql) AddAuthor(author book.Author) (book.Responce, error) {
	var id int

	tx, err := r.db.Begin()
	if err != nil {
		return book.Responce{}, err
	}
	query1 := fmt.Sprintf("INSERT INTO %s (first_name, second_name) VALUES ('%s','%s')", authorsTable, author.FirstName, author.SecondName)

	tx.QueryRow(query1)

	err = tx.QueryRow("SELECT LAST_INSERT_ID();").Scan(&id)
	if err != nil {
		tx.Rollback()
		return book.Responce{}, err
	}

	tx.Commit()

	strId := strconv.Itoa(id)

	return book.Responce{Id: strId}, nil
}
