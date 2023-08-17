package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	booksTable   = "books"
	authorsTable = "authors"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func NewMysqlDB(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn(cfg))

	if err != nil {
		fmt.Println("Нет подключения!")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Нет пинга!")
		return nil, err
	}

	fmt.Println("Подключение к базе все ОК!")
	return db, nil

}

func dsn(cfg Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.DBName)
}
