package transport

import (
	"context"
	"errors"

	"github.com/aburtasov/books"
	"github.com/aburtasov/books/api"
	"github.com/aburtasov/books/pkg/service"
)

type GRPCServer struct {
	services service.Service
	api.UnimplementedBookerServer
}

func NewGRPCServer(services service.Service) *GRPCServer {
	return &GRPCServer{
		services: services,
	}
}

func (s *GRPCServer) GetBooks(ctx context.Context, author *api.Author) (*api.Books, error) {

	var input books.Author
	input.FirstName = author.FirstName
	input.SecondName = author.SecondName

	if input.SecondName == "" {
		return &api.Books{}, errors.New("empty second name")
	}

	books, err := s.services.Book.GetBooks(input)
	if err != nil {
		return &api.Books{}, errors.New("internal server error")
	}

	return &api.Books{Title: books}, nil

}

func (s *GRPCServer) GetAuthor(ctx context.Context, books *api.Books) (*api.Author, error) {
	return &api.Author{}, nil
}
func (s *GRPCServer) AddBook(ctx context.Context, books *api.Book) (*api.ResponceID, error) {
	return &api.ResponceID{}, nil
}
func (s *GRPCServer) AddAuthor(ctx context.Context, author *api.Author) (*api.ResponceID, error) {
	return &api.ResponceID{}, nil
}
