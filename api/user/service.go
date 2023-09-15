package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/alizoubair/go-grpc-server/api/api_struct/form"
	"github.com/alizoubair/go-grpc-server/api/api_struct/model"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	CreateUser(ctx context.Context, req *form.UserForm) (*model.UserModel, error)
	GetUser(ctx context.Context, id, selectField string) (*model.UserModel, error)
	GetUsers(ctx context.Context, filter, filterCount map[string]interface{}, where, orderBy, selectField string) ([]*model.UserModel, int, error)
	UpdateUser(ctx context.Context, req *form.UserForm, id string) (*model.UserModel, error)
	DeleteUser(ctx context.Context, id string) error
}

type userService struct {
	log        *logrus.Entry
	repository UserRepository
}

var _ UserService = (*userService)(nil)

func NewUserService(log *logrus.Entry, r UserRepository) UserService {
	return &userService{
		log:        log,
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
		s.log.Errorf("can't create user: %s", err.Error())
		return nil, errors.New("Something went wrong. Please try again later")
	}

	return user, nil
}

func (s *userService) GetUser(ctx context.Context, id, selectField string) (*model.UserModel, error) {
	user, err := s.repository.GetUser(ctx, id, selectField)
	if err == sql.ErrNoRows {
		return nil, errors.New("User not found")
	}
	if err != nil {
		s.log.Errorf("can't get user: %s with id %v", err.Error(), id)
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUsers(ctx context.Context, filter, filterCount map[string]interface{}, where, orderBy, selectField string) ([]*model.UserModel, int, error) {
	users, err := s.repository.GetUsers(ctx, filter, where, orderBy, selectField)
	if err != nil {
		s.log.Errorf("can't get users: %s", err.Error())
		return nil, 0, errors.New("Something went wrong Please try again later")
	}

	count, err := s.repository.CountUsers(ctx, filterCount, where)
	if err != nil {
		s.log.Errorf("can't count users: %s", err.Error())
		return nil, 0, errors.New("Something went wrong. Please try again later")
	}

	return users, count, nil
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
		s.log.Errorf("can't update user: %s with id %v", err.Error(), id)
		return nil, errors.New("Something went wrong. Please try again later")
	}

	return user, nil
}

func (s *userService) DeleteUser(ctx context.Context, id string) error {
	_, err := s.repository.GetUser(ctx, id, "id")
	if err == sql.ErrNoRows {
		return errors.New("User not found")
	}

	if err != nil {
		s.log.Errorf("can't get user: %s with id %v", err.Error(), id)
		return errors.New("Something went wrong. Please try again later")
	}

	err = s.repository.DeleteUser(ctx, id)
	if err != nil {
		s.log.Errorf("can't delete user: %s", err.Error())
		return errors.New("Something went wrong. Please try again later")
	}

	return nil
}
