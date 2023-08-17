package main

import (
	"log"
	"net"

	"github.com/aburtasov/books/api"
	transport "github.com/aburtasov/books/pkg/transport/grpc"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &transport.GRPCServer{}

	api.RegisterBookerServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
