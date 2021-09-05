package main

import (
	"net"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	api "github.com/ozonva/ova-hobby-api/internal/app"
	"github.com/ozonva/ova-hobby-api/internal/db"
	"github.com/ozonva/ova-hobby-api/internal/repo"
	desc "github.com/ozonva/ova-hobby-api/pkg/github.com/ozonva/ova-hobby-api/pkg/ova-hobby-api"
)

func connectToDB() *sqlx.DB {
	pgUser := os.Getenv("POSTGRES_USER")
	pgPassword := os.Getenv("POSTGRES_PASSWORD")
	pgDB := os.Getenv("POSTGRES_DB")

	if pgUser == "" {
		log.Fatal().Msg("env variable POSTGRES_USER must be set")
	}
	if pgPassword == "" {
		log.Fatal().Msg("env variable POSTGRES_PASSWORD must be set")
	}
	if pgDB == "" {
		log.Fatal().Msg("env variable POSTGRES_DB must be set")
	}

	config := db.NewConfigDB("localhost", 5432, pgUser, pgPassword, pgDB, false)
	dbConnect, err := db.Connect(config)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	return dbConnect
}

func run(r repo.Repo) error {
	grpcServerEndpoint := os.Getenv("GRPC_SERVER_ADDRESS")
	if grpcServerEndpoint == "" {
		log.Fatal().Msg("env variable GRPC_SERVER_ADDRESS must be set")
	}

	listen, err := net.Listen("tcp", grpcServerEndpoint)
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.ChainUnaryInterceptor(repo.NewInterceptorWithRepo(r)))
	desc.RegisterHobbyAPIServer(s, api.NewHobbyAPI())

	log.Info().Msg("Server is running")
	if err := s.Serve(listen); err != nil {
		log.Fatal().Msgf("failed to serve: %v", err)
	}

	return nil
}

func main() {
	dbConn := connectToDB()
	if err := run(repo.NewRepo(dbConn)); err != nil {
		log.Fatal().Msg(err.Error())
	}
}
