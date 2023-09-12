package user_test

import (
	"context"
	"errors"
	"testing"

	"github.com/alizoubair/go-grpc-server/api/api_struct/form"
	"github.com/alizoubair/go-grpc-server/api/api_struct/model"
	"github.com/alizoubair/go-grpc-server/api/user"
	"github.com/alizoubair/go-grpc-server/api/user/mocks"
	"github.com/alizoubair/go-grpc-server/config"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type Service struct {
	mock.Mock
}

func TestCreateUserService(t *testing.T) {
	log := config.InitLog()
	mockRepo := new(mocks.Repository)

	reqUser := &form.UserForm{
		ID: xid.New().String(),
		Name: "John",
		Email: "john@email.com",
		Address: "Japan",
		Phone: "(012)-345-6789",
	}

	mockUser := &model.UserModel{
		ID: reqUser.ID,
		Name: reqUser.Name,
		Email: reqUser.Email,
		Address: reqUser.Address,
		Phone: reqUser.Phone,
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("CreateUser", mockUser).Return(mockUser, nil).Once()
		u := user.NewUserService(log, mockRepo)

		userRow, err := u.CreateUser(context.Background(), reqUser)

		assert.NoError(t, err)
		assert.NotNil(t, userRow)

		mockRepo.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		mockRepo.On("CreateUser", mockUser).Return(nil, errors.New("Unexpected database error")).Once()
		u := user.NewUserService(log, mockRepo)

		userRow, err := u.CreateUser(context.Background(), reqUser)

		assert.Error(t, err)
		assert.Nil(t, userRow)

		mockRepo.AssertExpectations(t)
	})
}