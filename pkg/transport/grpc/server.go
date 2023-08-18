package transport

import (
	"context"
	"errors"

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

	return &api.Books{Title: titles}, nil

}

func (s *GRPCServer) AddBook(ctx context.Context, input *api.AddBook) (*api.ResponceID, error) {

	var book book.Book
	book.Title = input.Title
	book.Description = input.Description
	book.Author = int(input.AuthorId)

	result, err := s.services.Book.AddBook(book)
	if err != nil {
		return nil, errors.New("internal server error")
	}

	return &api.ResponceID{Value: result.Id}, nil
}

func (s *GRPCServer) GetAuthor(ctx context.Context, book *api.Book) (*api.Author, error) {

	var input = book.Title

	author, err := s.services.Author.GetAuthor(input)
	if err != nil {

		return nil, errors.New("internal server error")
	}

	return &api.Author{FirstName: author.FirstName, SecondName: author.SecondName}, nil
}

func (s *GRPCServer) AddAuthor(ctx context.Context, author *api.Author) (*api.ResponceID, error) {

	var aut book.Author
	aut.FirstName = author.FirstName
	aut.SecondName = author.SecondName

	result, err := s.services.Author.AddAuthor(aut)
	if err != nil {
		return nil, errors.New("internal server error")
	}
	return &api.ResponceID{Value: result.Id}, nil
}
