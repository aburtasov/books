package repository

import (
	"database/sql"
	"fmt"

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
			fmt.Println("Тут", err)
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		bks = append(bks, bok)
	}
	fmt.Println("ОКЕЙ")
	return bks, nil
}
