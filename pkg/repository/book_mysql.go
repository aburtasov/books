package repository

import (
	"database/sql"
	"fmt"

	"github.com/aburtasov/books"
	_ "github.com/go-sql-driver/mysql"
)

type BookMysql struct {
	db *sql.DB
}

func NewBookMysql(db *sql.DB) *BookMysql {
	return &BookMysql{db: db}
}

func (r *BookMysql) GetBooks(author books.Author) ([]string, error) {
	var id int64
	var books []string
	fmt.Println("Дошли до сюда")

	query1 := fmt.Sprintf("SELECT id FROM %s WHERE first_name=$1 AND second_name=$2", authorsTable)
	fmt.Println("Дошли до сюда 1")
	result, err := r.db.Exec(query1, author.FirstName, author.SecondName)
	if err != nil {
		return nil, err
	}
	id, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}

	fmt.Println("Дошли до сюда 2")

	query2 := fmt.Sprintf("SELECT title FROM %s WHERE author_id=$1", booksTable)
	rows, err := r.db.Query(query2, id)
	if err != nil {
		fmt.Println("Тут 3")
		return nil, err
	}

	var book string

	for rows.Next() {

		err = rows.Scan(&book)
		if err != nil {
			fmt.Println("Тут 4")
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		books = append(books, book)
	}

	return books, nil
}
