package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	book "github.com/aburtasov/books"
	_ "github.com/go-sql-driver/mysql"
)

type BookMysql struct {
	db *sql.DB
}

func NewBookMysql(db *sql.DB) *BookMysql {
	return &BookMysql{db: db}
}

func (r *BookMysql) GetBooks(author book.Author) ([]book.Book, error) {
	var id int64
	var bks []book.Book
	var bok book.Book

	query1 := fmt.Sprintf("SELECT id FROM %s WHERE first_name='%s' AND second_name='%s'", authorsTable, author.FirstName, author.SecondName)

	err := r.db.QueryRow(query1).Scan(&id)
	if err != nil {
		return nil, err
	}

	query2 := fmt.Sprintf("SELECT * FROM %s WHERE author_id=%d", booksTable, id)

	rows, err := r.db.Query(query2)
	if err != nil {
		return nil, err
	}
	for rows.Next() {

		err = rows.Scan(&bok.Id, &bok.Title, &bok.Description, &bok.Author)
		if err != nil {

			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		bks = append(bks, bok)
	}

	return bks, nil
}

func (r *BookMysql) AddBook(bk book.Book) (book.Responce, error) {

	var id int

	tx, err := r.db.Begin()
	if err != nil {
		return book.Responce{}, err
	}
	query1 := fmt.Sprintf("INSERT INTO %s (title, description,author_id) VALUES ('%s','%s','%d')", booksTable, bk.Title, bk.Description, bk.Author)

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
