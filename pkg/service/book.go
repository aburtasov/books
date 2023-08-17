package service

import (
	"github.com/aburtasov/books"
	"github.com/aburtasov/books/pkg/repository"
)

type BookService struct {
	repo repository.Book
}

func NewBookService(repo repository.Book) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetBooks(author books.Author) ([]string, error) {
	// Дополнительная логика не будет реализовываться
	// Сразу передача в репозиторий
	return s.repo.GetBooks(author)
}
