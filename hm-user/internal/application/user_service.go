package application

import (
	"context"

	"github.com/h-mall/hm-user/internal/domain"
	v1 "github.com/h-mall/proto-repo/api/user/v1"
)

type UserService struct {
	v1.UnimplementedUserServiceServer
}

func NewUserService() *UserService  {
	return &UserService{}
}

func (u *UserService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserResponse, error)  {
	user := domain.FindUserByID(req.GetId())
	if user == nil {
		return nil, nil
	}

	return &v1.GetUserResponse{
		Id: user.ID,
		Name: user.Name,
		Email: user.Email,
	}, nil
}