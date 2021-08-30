package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	api "github.com/ozonva/ova-hobby-api/internal/app"
	desc "github.com/ozonva/ova-hobby-api/pkg/github.com/ozonva/ova-hobby-api/pkg/ova-hobby-api"
)

const (
	grpcServerEndpoint = "localhost:8082"
)

func run() error {
	listen, err := net.Listen("tcp", grpcServerEndpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	desc.RegisterHobbyAPIServer(s, api.NewHobbyAPI())

	log.Println("Server is running")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
