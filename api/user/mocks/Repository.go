package mocks

import (
	"context"

	"github.com/alizoubair/go-grpc-server/api/api_struct/model"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (m *Repository) CreateUser(ctx context.Context, userReq *model.UserModel) (*model.UserModel, error) {
	ret := m.Called(userReq)

	var r0 *model.UserModel
	if rf, ok := ret.Get(0).(func(*model.UserModel) *model.UserModel); ok {
		r0 = rf(userReq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.UserModel) error); ok {
		r1 = rf(userReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (*Repository) GetUser(ctx context.Context, id string, selectField string) (*model.UserModel, error) {
	panic("unimplemented")
}

func (*Repository) GetUsers(ctx context.Context, filter map[string]interface{}, where string, orderBy string, selectField string) ([]*model.UserModel, error) {
	panic("unimplemented")
}

func (*Repository) UpdateUser(ctx context.Context, user *model.UserModel) (*model.UserModel, error) {
	panic("unimplemented")
}

func (*Repository) DeleteUser(ctx context.Context, id string) error {
	panic("unimplemented")
}