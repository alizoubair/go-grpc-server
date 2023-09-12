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
	mockService := new(mocks.Service)
	mockService.On("CreateUser", mock.AnythingOfType("*form.UserForm")).Return(nil, errors.New("Unexpected database error"))
	usrCtrl := usrGrpc.NewUserServerGrpc(server, mockService)

	createUserRes, err := usrCtrl.CreateUser(context.Background(), &proto.CreateUserRequest{
		Name: "John",
		Phone: "(012)-345-6789",
		Email: "john@email.com",
		Address: "Japan",
	})
	assert.Nil(t, createUserRes)
	assert.Error(t, err)
}

func TestDeliveryCreateSuccess(t *testing.T) {
	user := &model.UserModel{
		ID: "id",
		Name: "John",
		Email: "john@email.com",
		Address: "Japan",
		Phone: "(012)-345-6789",
	}

	server := grpc.NewServer()
	mocksService := new(mocks.Service)
	mocksService.On("CreateUser", mock.AnythingOfType("*form.UserForm")).Return(user, nil)
	usrCtrl := usrGrpc.NewUserServerGrpc(server, mocksService)

	createUserReq, err :=usrCtrl.CreateUser(context.Background(), &proto.CreateUserRequest{
		Name: user.Name,
		Email: user.Email,
		Address: user.Address,
		Phone: user.Phone,
	})

	assert.NotNil(t, createUserReq)
	assert.NoError(t, err)
}