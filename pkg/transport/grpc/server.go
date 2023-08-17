package transport

import (
	"context"
	"errors"
	"fmt"

	book "github.com/aburtasov/books"
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

	titles := make([]string, 0)
	var input book.Author

	input.FirstName = author.FirstName
	input.SecondName = author.SecondName

	if input.FirstName == "" {
		return nil, errors.New("empty first name")
	}

	if input.SecondName == "" {
		return nil, errors.New("empty second name")
	}

	if input.SecondName == "" && input.FirstName == "" {
		return nil, errors.New("empty arguments")
	}

	result, err := s.services.Book.GetBooks(input)

	if err != nil {
		return nil, errors.New("internal server error")
	}

	for i := range result {
		titles = append(titles, result[i].Title)
	}

	fmt.Println("Возврат клиенту")

	return &api.Books{Title: titles}, nil

}

func (s *GRPCServer) GetAuthor(ctx context.Context, book *api.Book) (*api.Author, error) {
	return &api.Author{}, nil
}
func (s *GRPCServer) AddBook(ctx context.Context, book *api.AddBook) (*api.ResponceID, error) {
	return &api.ResponceID{}, nil
}
func (s *GRPCServer) AddAuthor(ctx context.Context, author *api.Author) (*api.ResponceID, error) {
	return &api.ResponceID{}, nil
}
