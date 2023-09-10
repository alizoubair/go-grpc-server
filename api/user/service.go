package user

import (
	"context"
	"log"

	"github.com/alizoubair/go-grpc-server/api/api_struct/form"
	"github.com/alizoubair/go-grpc-server/api/api_struct/model"
)

type UserService interface {
	CreateUser(ctx context.Context, req *form.UserForm) (*model.UserModel, error)
}

type userService struct {
	repository UserRepository
}

func NewUserService(r UserRepository) UserService {
	return &userService{
		repository: r,
	}
}

func (s *userService) CreateUser(ctx context.Context, req *form.UserForm) (*model.UserModel, error) {
	userReq := &model.UserModel{
		ID:      req.ID,
		Name:    req.Name,
		Email:   req.Email,
		Address: req.Address,
		Phone:   req.Phone,
	}

	user, err := s.repository.CreateUser(ctx, userReq)
	if err != nil {
		log.Printf("can't create user: %s", err.Error())
		return nil, err
	}

	return user, nil
}
