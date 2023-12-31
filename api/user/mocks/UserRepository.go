// Code generated by mockery v2.33.2. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/alizoubair/go-grpc-server/api/api_struct/model"
	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// CountUsers provides a mock function with given fields: ctx, filter, where
func (_m *UserRepository) CountUsers(ctx context.Context, filter map[string]interface{}, where string) (int, error) {
	ret := _m.Called(ctx, filter, where)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}, string) (int, error)); ok {
		return rf(ctx, filter, where)
	}
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}, string) int); ok {
		r0 = rf(ctx, filter, where)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, map[string]interface{}, string) error); ok {
		r1 = rf(ctx, filter, where)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: ctx, _a1
func (_m *UserRepository) CreateUser(ctx context.Context, _a1 *model.UserModel) (*model.UserModel, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *model.UserModel
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.UserModel) (*model.UserModel, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.UserModel) *model.UserModel); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserModel)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.UserModel) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUser provides a mock function with given fields: ctx, id
func (_m *UserRepository) DeleteUser(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUser provides a mock function with given fields: ctx, id, selectField
func (_m *UserRepository) GetUser(ctx context.Context, id string, selectField string) (*model.UserModel, error) {
	ret := _m.Called(ctx, id, selectField)

	var r0 *model.UserModel
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*model.UserModel, error)); ok {
		return rf(ctx, id, selectField)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *model.UserModel); ok {
		r0 = rf(ctx, id, selectField)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserModel)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, id, selectField)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsers provides a mock function with given fields: ctx, filter, where, orderBy, selectField
func (_m *UserRepository) GetUsers(ctx context.Context, filter map[string]interface{}, where string, orderBy string, selectField string) ([]*model.UserModel, error) {
	ret := _m.Called(ctx, filter, where, orderBy, selectField)

	var r0 []*model.UserModel
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}, string, string, string) ([]*model.UserModel, error)); ok {
		return rf(ctx, filter, where, orderBy, selectField)
	}
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}, string, string, string) []*model.UserModel); ok {
		r0 = rf(ctx, filter, where, orderBy, selectField)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.UserModel)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, map[string]interface{}, string, string, string) error); ok {
		r1 = rf(ctx, filter, where, orderBy, selectField)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, _a1
func (_m *UserRepository) UpdateUser(ctx context.Context, _a1 *model.UserModel) (*model.UserModel, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *model.UserModel
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.UserModel) (*model.UserModel, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.UserModel) *model.UserModel); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserModel)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.UserModel) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
