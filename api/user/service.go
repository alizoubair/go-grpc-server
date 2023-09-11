package user

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/alizoubair/go-grpc-server/api/api_struct/form"
	"github.com/alizoubair/go-grpc-server/api/api_struct/model"
)

type UserService interface {
	CreateUser(ctx context.Context, req *form.UserForm) (*model.UserModel, error)
	GetUser(ctx context.Context, id, selectField string) (*model.UserModel, error)
	UpdateUser(ctx context.Context, req *form.UserForm, id string) (*model.UserModel, error)
	DeleteUser(ctx context.Context, id string) error
}

type userService struct {
	repository UserRepository
}

var _ UserService = (*userService)(nil)

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

func (s *userService) GetUser(ctx context.Context, id, selectField string) (*model.UserModel, error) {
	user, err := s.repository.GetUser(ctx, id, selectField)
	if err == sql.ErrNoRows {
		return nil, errors.New("User not found")
	}
	if err != nil {
		log.Printf("can't get user: %s with id %v", err.Error(), id)
		return nil, err
	}

	return user, nil
}

func (s *userService) UpdateUser(ctx context.Context, req *form.UserForm, id string) (*model.UserModel, error) {
	user := &model.UserModel{
		ID:      id,
		Name:    req.Name,
		Email:   req.Email,
		Address: req.Address,
		Phone:   req.Phone,
	}

	user, err := s.repository.UpdateUser(ctx, user)
	if err != nil {
		return nil, errors.New("Something went wrong. Please try again later")
	}

	return user, nil
}


func (s *userService) DeleteUser(ctx context.Context, id string) error {
	_, err := s.repository.GetUser(ctx, id, "id")
	if err != sql.ErrNoRows {
		return errors.New("User not found")
	}

	if err != nil {
		return errors.New("Something went wrong. Please try again later")
	}

	err = s.repository.DeleteUser(ctx, id)
	if err != nil {
		return errors.New("Something went wrong. Please try again later")
	}

	return nil
}