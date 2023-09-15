package grpc_test

import (
	"context"
	"errors"
	"testing"

	"github.com/alizoubair/go-grpc-server/api/api_struct/model"
	usrGrpc "github.com/alizoubair/go-grpc-server/api/user/grpc"
	"github.com/alizoubair/go-grpc-server/api/user/mocks"
	"github.com/alizoubair/go-grpc-server/api/user/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

func TestDeliveryCreateFail(t *testing.T) {
	server := grpc.NewServer()
	mockService := new(mocks.UserService)
	mockService.On("CreateUser", context.Background() ,mock.AnythingOfType("*form.UserForm")).Return(nil, errors.New("Unexpected database error"))
	usrCtrl := usrGrpc.NewUserServerGrpc(server, mockService)

	createUserRes, err := usrCtrl.CreateUser(context.Background(), &proto.CreateUserRequest{
		Name:    "John",
		Phone:   "(012)-345-6789",
		Email:   "john@email.com",
		Address: "Japan",
	})
	
	assert.Nil(t, createUserRes)
	assert.Error(t, err)
}

func TestDeliveryCreateSuccess(t *testing.T) {
	user := &model.UserModel{
		ID:      "id",
		Name:    "John",
		Email:   "john@email.com",
		Address: "Japan",
		Phone:   "(012)-345-6789",
	}

	server := grpc.NewServer()
	mockService := new(mocks.UserService)
	mockService.On("CreateUser", context.Background(), mock.AnythingOfType("*form.UserForm")).Return(user, nil)
	usrCtrl := usrGrpc.NewUserServerGrpc(server, mockService)

	createUserReq, err := usrCtrl.CreateUser(context.Background(), &proto.CreateUserRequest{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	})

	assert.NotNil(t, createUserReq)
	assert.NoError(t, err)
}

func TestDeliveryGetUserFail(t *testing.T) {
	server := grpc.NewServer()
	mockService := new(mocks.UserService)
	mockService.On("GetUser", context.Background(), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, errors.New("Unexpected database error"))
	usrCtrl := usrGrpc.NewUserServerGrpc(server, mockService)

	detailUserRes, err := usrCtrl.GetUser(context.Background(), &proto.GetUserRequest{
		Id: "id",
	})

	assert.Error(t, err)
	assert.Nil(t, detailUserRes)
}

func TestDeliveryGetUserSuccess(t *testing.T) {
	user := &model.UserModel{
		ID:      "id",
		Name:    "John",
		Email:   "John@email.com",
		Address: "Japan",
		Phone:   "(012)-345-6789",
	}

	server := grpc.NewServer()
	mockService := new(mocks.UserService)
	mockService.On("GetUser", context.Background(), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(user, nil)
	usrCtrl := usrGrpc.NewUserServerGrpc(server, mockService)

	detailUserRes, err := usrCtrl.GetUser(context.Background(), &proto.GetUserRequest{
		Id: "id",
	})

	assert.NoError(t, err)
	assert.NotNil(t, detailUserRes)
}

func TestDeliveryUpdateFail(t *testing.T) {
	user := &model.UserModel{
		ID:      "id",
		Name:    "John",
		Email:   "John@email.com",
		Address: "Japan",
		Phone:   "(012)-345-6789",
	}

	server := grpc.NewServer()
	mockService := new(mocks.UserService)
	mockService.On("GetUser", context.Background(), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(user, nil)
	mockService.On("UpdateUser", context.Background(), mock.AnythingOfType("*form.UserForm"), mock.AnythingOfType("string")).Return(nil, errors.New("Unexpected database error"))
	usrCtrl := usrGrpc.NewUserServerGrpc(server, mockService)

	updatedUserRes, err := usrCtrl.UpdateUser(context.Background(), &proto.UpdateUserRequest{
		Id:      "id",
		Name:    "John",
		Email:   "John@email.com",
		Address: "Japan",
		Phone:   "(012)-345-6789",
	})

	assert.Nil(t, updatedUserRes)
	assert.Error(t, err)
}

func TestDeliveryUpdateSuccess(t *testing.T) {
	user := &model.UserModel{
		ID:      "id",
		Name:    "John",
		Email:   "John@email.com",
		Address: "John@email.com",
		Phone:   "(012)-345-6789",
	}

	server := grpc.NewServer()
	mockService := new(mocks.UserService)
	mockService.On("GetUser", context.Background(), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(user, nil)
	mockService.On("UpdateUser", context.Background(), mock.AnythingOfType("*form.UserForm"), mock.AnythingOfType("string")).Return(user, nil)
	usrCtrl := usrGrpc.NewUserServerGrpc(server, mockService)

	updatedUserRes, err := usrCtrl.UpdateUser(context.Background(), &proto.UpdateUserRequest{
		Id:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	})

	assert.NoError(t, err)
	assert.NotNil(t, updatedUserRes)
}

func TestDeliveryDeleteFail(t *testing.T) {
	server := grpc.NewServer()
	mockService := new(mocks.UserService)
	mockService.On("DeleteUser", context.Background(), mock.AnythingOfType("string")).Return(errors.New("Unexpected database error"))
	usrCtrl := usrGrpc.NewUserServerGrpc(server, mockService)

	_, err := usrCtrl.DeleteUser(context.Background(), &proto.DeleteUserRequest{
		Id: "id",
	})

	assert.Error(t, err)
}

func TestDeliveryDeleteSucess(t *testing.T) {
	server := grpc.NewServer()
	mockService := new(mocks.UserService)
	mockService.On("DeleteUser", context.Background(), mock.AnythingOfType("string")).Return(nil)
	usrCtrl := usrGrpc.NewUserServerGrpc(server, mockService)

	_, err := usrCtrl.DeleteUser(context.Background(), &proto.DeleteUserRequest{
		Id: "id",
	})

	assert.NoError(t, err)
}
