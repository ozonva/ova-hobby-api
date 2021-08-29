package app

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/rs/zerolog/log"

	desc "github.com/ozonva/ova-hobby-api/pkg/github.com/ozonva/ova-hobby-api/pkg/ova-hobby-api"
)

func NewHobbyAPI() desc.HobbyAPIServer {
	return &desc.UnimplementedHobbyAPIServer{}
}

type hobbyAPI struct {
	desc.HobbyAPIServer
}

func (api *hobbyAPI) CreateHobby(ctx context.Context, request *desc.CreateRequest) (*desc.HobbyReply, error) {
	log.Info().Msg("CreateHobby call")

	return &desc.HobbyReply{}, nil
}
func (api *hobbyAPI) DescribeHobby(ctx context.Context, request *desc.DescribeRequest) (*desc.HobbyReply, error) {
	log.Info().Msg("DescribeHobby call")

	return &desc.HobbyReply{}, nil
}
func (api *hobbyAPI) ListHobby(ctx context.Context, request *desc.ListRequest) (*desc.ListReply, error) {
	log.Info().Msg("ListHobby call")

	return &desc.ListReply{}, nil
}
func (api *hobbyAPI) RemoveHobby(ctx context.Context, request *desc.RemoveRequest) (*emptypb.Empty, error) {
	log.Info().Msg("RemoveHobby call")

	return &emptypb.Empty{}, nil
}
