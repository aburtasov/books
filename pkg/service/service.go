package service

import (
	book "github.com/aburtasov/books"
	"github.com/aburtasov/books/pkg/repository"
)

type Book interface {
	GetBooks(author book.Author) ([]book.Book, error)
	AddBook(book book.Book) (book.Responce, error)
}

type Author interface {
	GetAuthor(bookTitle string) (book.Author, error)
	AddAuthor(author book.Author) (book.Responce, error)
}

type Service struct {
	Book
	Author
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Book:   NewBookService(repos),
		Author: NewAuthorService(repos),
	}
}
