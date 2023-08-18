package service

import (
	book "github.com/aburtasov/books"
	"github.com/aburtasov/books/pkg/repository"
)

type AuthorService struct {
	repo repository.Author
}

func NewAuthorService(repo repository.Author) *AuthorService {
	return &AuthorService{repo: repo}
}

func (s *AuthorService) GetAuthor(bookTitle string) (book.Author, error) {
	// Дополнительная логика не будет реализовываться
	// Сразу передача в репозиторий
	return s.repo.GetAuthor(bookTitle)
}

func (s *AuthorService) AddAuthor(author book.Author) (book.Responce, error) {
	// Дополнительная логика не будет реализовываться
	// Сразу передача в репозиторий
	return s.repo.AddAuthor(author)
}
