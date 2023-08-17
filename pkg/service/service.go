package service

import (
	book "github.com/aburtasov/books"
	"github.com/aburtasov/books/pkg/repository"
)

type Book interface {
	GetBooks(author book.Author) ([]book.Book, error)
}

type Author interface {
}

type Service struct {
	Book
	Author
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Book: NewBookService(repos),
	}
}
