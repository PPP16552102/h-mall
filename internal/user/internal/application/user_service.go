package application

import (
	"context"

	"github.com/h-mall/user/api"
	"github.com/h-mall/user/internal/domain"
)

type UserService struct {
	api.UnimplementedUserServiceServer
}

func NewUserService() *UserService  {
	return &UserService{}
}

func (u *UserService) GetUser(ctx context.Context, req *api.GetUserRequest) (*api.GetUserResponse, error)  {
	user := domain.FindUserByID(req.GetId())
	if user == nil {
		return nil, nil
	}

	return &api.GetUserResponse{
		Id: user.ID,
		Name: user.Name,
		Email: user.Email,
	}, nil
}