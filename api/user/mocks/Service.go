package mocks

import (
	"context"

	"github.com/alizoubair/go-grpc-server/api/api_struct/form"
	"github.com/alizoubair/go-grpc-server/api/api_struct/model"
	"github.com/stretchr/testify/mock"
)

type Service struct {
	mock.Mock
}

func (m *Service) CreateUser(ctx context.Context, req *form.UserForm) (*model.UserModel, error) {
	ret := m.Called(req)

	var r0 *model.UserModel
	if rf, ok := ret.Get(0).(func(*form.UserForm) *model.UserModel); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*form.UserForm) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *Service) GetUser(ctx context.Context, id string, selectField string) (*model.UserModel, error) {
	panic("unimplemented")
}

func (m *Service) UpdateUser(ctx context.Context, req *form.UserForm, id string) (*model.UserModel, error) {
	panic("unimplemented")
}

func (m *Service) DeleteUser(ctx context.Context, id string) error {
	panic("unimplemented")
}