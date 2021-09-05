package app

import (
	"context"

	"github.com/google/uuid"
	"github.com/ozonva/ova-hobby-api/internal/repo"
	"github.com/ozonva/ova-hobby-api/pkg/models"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/ozonva/ova-hobby-api/pkg/github.com/ozonva/ova-hobby-api/pkg/ova-hobby-api"
)

func NewHobbyAPI() desc.HobbyAPIServer {
	return &desc.UnimplementedHobbyAPIServer{}
}

type hobbyAPI struct {
	desc.HobbyAPIServer
}

func (api *hobbyAPI) CreateHobby(ctx context.Context, request *desc.CreateRequest) (*desc.HobbyReply, error) {
	newHobby := models.NewHobby(request.GetName(), request.GetUserId(), models.HobbyKind(request.GetHobbyKind()))
	err := repo.FromContext(ctx).AddHobby(newHobby)

	if err != nil {
		log.Error().Msgf("unable to create a hobby, %v", request)
		return nil, err
	}

	log.Debug().Msg("CreateHobby successful")
	return &desc.HobbyReply{
		Uuid: newHobby.ID.String(), Name: newHobby.Name, HobbyKind: uint32(newHobby.Kind), UserId: newHobby.UserID,
	}, nil
}
func (api *hobbyAPI) DescribeHobby(ctx context.Context, request *desc.DescribeRequest) (*desc.HobbyReply, error) {
	var hobby *models.Hobby
	id, err := uuid.Parse(request.Uuid)
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}

	hobby, err = repo.FromContext(ctx).DescribeHobby(id)
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}
	log.Debug().Msg("DescribeHobby successful")
	return &desc.HobbyReply{
		Uuid: hobby.ID.String(), Name: hobby.Name, HobbyKind: uint32(hobby.Kind), UserId: hobby.UserID,
	}, nil
}
func (api *hobbyAPI) ListHobby(ctx context.Context, request *desc.ListRequest) (*desc.ListReply, error) {
	hobbies, err := repo.FromContext(ctx).ListHobbies(
		uint64(request.GetLimit()), uint64(request.GetOffset()),
	)
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}

	var listReply []*desc.HobbyReply
	for _, hobby := range hobbies {
		hobbyReply := &desc.HobbyReply{
			Uuid: hobby.ID.String(), Name: hobby.Name, HobbyKind: uint32(hobby.Kind), UserId: hobby.UserID,
		}
		listReply = append(listReply, hobbyReply)
	}
	log.Debug().Msg("ListHobby successful")
	return &desc.ListReply{Members: listReply}, nil
}
func (api *hobbyAPI) RemoveHobby(ctx context.Context, request *desc.RemoveRequest) (*emptypb.Empty, error) {
	id, err := uuid.Parse(request.Uuid)
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}

	err = repo.FromContext(ctx).RemoveHobby(id)
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}
	log.Debug().Msg("RemoveHobby successful")
	return &emptypb.Empty{}, nil
}
