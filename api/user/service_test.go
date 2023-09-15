package user_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

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

func TestServiceCreateUser(t *testing.T) {
	log := config.InitLog()
	mockRepo := new(mocks.UserRepository)

	reqUser := &form.UserForm{
		ID:      xid.New().String(),
		Name:    "John",
		Email:   "john@email.com",
		Address: "Japan",
		Phone:   "(012)-345-6789",
	}

	mockUser := &model.UserModel{
		ID:      reqUser.ID,
		Name:    reqUser.Name,
		Email:   reqUser.Email,
		Address: reqUser.Address,
		Phone:   reqUser.Phone,
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("CreateUser", context.Background(), mockUser).Return(mockUser, nil).Once()
		u := user.NewUserService(log, mockRepo)

		userRow, err := u.CreateUser(context.Background(), reqUser)

		assert.NoError(t, err)
		assert.NotNil(t, userRow)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockRepo.On("CreateUser", context.Background() ,mockUser).Return(nil, errors.New("Unexpected database error")).Once()
		u := user.NewUserService(log, mockRepo)

		userRow, err := u.CreateUser(context.Background(), reqUser)

		assert.Error(t, err)
		assert.Nil(t, userRow)

		mockRepo.AssertExpectations(t)
	})
}

func TestServiceGetUser(t *testing.T) {
	log := config.InitLog()
	mockRepo := new(mocks.UserRepository)

	mockUser := &model.UserModel{
		ID:        xid.New().String(),
		Name:      "John",
		Email:     "John@email.com",
		Address:   "Japan",
		Phone:     "(012)-345-6789",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetUser", context.Background(), mock.AnythingOfType("string"), "").Return(mockUser, nil).Once()

		u := user.NewUserService(log, mockRepo)
		userRow, err := u.GetUser(context.Background(), mockUser.ID, "")

		assert.NoError(t, err)
		assert.NotNil(t, userRow)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed-not-found", func(t *testing.T) {
		mockRepo.On("GetUser", context.Background(), mock.AnythingOfType("string"), "").Return(nil, sql.ErrNoRows).Once()

		u := user.NewUserService(log, mockRepo)
		userRow, err := u.GetUser(context.Background(), mockUser.ID, "")

		assert.Error(t, err)
		assert.Nil(t, userRow)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockRepo.On("GetUser", context.Background(), mock.AnythingOfType("string"), "").Return(nil, errors.New("Unexpected database error")).Once()

		u := user.NewUserService(log, mockRepo)
		userRow, err := u.GetUser(context.Background(), mockUser.ID, "")

		assert.Error(t, err)
		assert.Nil(t, userRow)
	})
}

func TestServiceGetUsers(t *testing.T) {
	log := config.InitLog()
	mockRepo := new(mocks.UserRepository)

	mockUser := &model.UserModel{
		ID:        xid.New().String(),
		Name:      "John",
		Email:     "John@email.com",
		Address:   "Japan",
		Phone:     "(012)-345-6789",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	mockUsers := make([]*model.UserModel, 0)
	mockUsers = append(mockUsers, mockUser)

	filter := map[string]interface{}{}
	filter["limit"] = "10"
	filter["offset"] = "0"
	orderBy := "id DESC"
	deletedNull := "WHERE deleted_at IS NULL"

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetUsers", context.Background(), filter, deletedNull, orderBy, model.UserSelectField).Return(mockUsers, nil).Once()
		mockRepo.On("CountUsers", context.Background(), filter, deletedNull).Return(1, nil).Once()
		u := user.NewUserService(log, mockRepo)

		users, count, err := u.GetUsers(context.Background(), filter, filter, deletedNull, orderBy, model.UserSelectField)

		assert.NoError(t, err)
		assert.NotNil(t, users)
		assert.Equal(t, 1, count)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed-get", func(t *testing.T) {
		mockRepo.On("GetUsers", context.Background(), filter, deletedNull, orderBy, model.UserSelectField).Return(nil, errors.New("Unexpected database error")).Once()
		u := user.NewUserService(log, mockRepo)

		users, count, err := u.GetUsers(context.Background(), filter, filter, deletedNull, orderBy, model.UserSelectField)

		assert.Error(t, err)
		assert.Nil(t, users)
		assert.Equal(t, 0, count)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed-count", func(t *testing.T) {
		mockRepo.On("GetUsers", context.Background(), filter, deletedNull, orderBy, model.UserSelectField).Return(mockUsers, nil).Once()
		mockRepo.On("CountUsers", context.Background(), filter, deletedNull).Return(0, errors.New("Unexpected database error")).Once()

		u := user.NewUserService(log, mockRepo)

		users, count, err := u.GetUsers(context.Background(), filter, filter, deletedNull, orderBy, model.UserSelectField)

		assert.Error(t, err)
		assert.Nil(t, users)
		assert.Equal(t, 0, count)

		mockRepo.AssertExpectations(t)
	})
}

func TestServiceUpdateUser(t *testing.T) {
	log := config.InitLog()
	mockRepo := new(mocks.UserRepository)

	reqUser := &form.UserForm{
		Name:    "John",
		Email:   "John@email.com",
		Address: "Japan",
		Phone:   "(012)-345-6789",
	}

	mockUser := &model.UserModel{
		ID:      xid.New().String(),
		Name:    reqUser.Name,
		Email:   reqUser.Email,
		Address: reqUser.Address,
		Phone:   reqUser.Phone,
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("UpdateUser", context.Background(), mockUser).Return(mockUser, nil).Once()

		u := user.NewUserService(log, mockRepo)
		userRow, err := u.UpdateUser(context.Background(), reqUser, mockUser.ID)

		assert.NoError(t, err)
		assert.NotNil(t, userRow)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockRepo.On("UpdateUser", context.Background(), mockUser).Return(nil, errors.New("Unexpected database error")).Once()

		u := user.NewUserService(log, mockRepo)
		userRow, err := u.UpdateUser(context.Background(), reqUser, mockUser.ID)

		assert.Error(t, err)
		assert.Nil(t, userRow)

		mockRepo.AssertExpectations(t)
	})
}

func TestServiceDeleteUser(t *testing.T) {
	log := config.InitLog()
	mockRepo := new(mocks.UserRepository)

	mockUser := &model.UserModel{
		ID:        xid.New().String(),
		Name:      "John",
		Email:     "John@email.com",
		Address:   "Japan",
		Phone:     "(012)-345-6789",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetUser", context.Background(), mock.AnythingOfType("string"), "id").Return(mockUser, nil).Once()
		mockRepo.On("DeleteUser", context.Background(), mock.AnythingOfType("string")).Return(nil).Once()

		u := user.NewUserService(log, mockRepo)
		err := u.DeleteUser(context.Background(), mockUser.ID)

		assert.NoError(t, err)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed-delete", func(t *testing.T) {
		mockRepo.On("GetUser", context.Background(), mock.AnythingOfType("string"), "id").Return(mockUser, nil).Once()
		mockRepo.On("DeleteUser", context.Background(), mock.AnythingOfType("string")).Return(errors.New("Unexpected database error")).Once()

		u := user.NewUserService(log, mockRepo)
		err := u.DeleteUser(context.Background(), mockUser.ID)

		assert.Error(t, err)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed-get-not-found", func(t *testing.T) {
		mockRepo.On("GetUser", context.Background(), mock.AnythingOfType("string"), "id").Return(mockUser, sql.ErrNoRows).Once()

		u := user.NewUserService(log, mockRepo)
		err := u.DeleteUser(context.Background(), mockUser.ID)

		assert.Error(t, err)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed-get", func(t *testing.T) {
		mockRepo.On("GetUser", context.Background(), mock.AnythingOfType("string"), "id").Return(nil, errors.New("Unexpected database error")).Once()

		u := user.NewUserService(log, mockRepo)
		err := u.DeleteUser(context.Background(), mockUser.ID)

		assert.Error(t, err)

		mockRepo.AssertExpectations(t)
	})
}
