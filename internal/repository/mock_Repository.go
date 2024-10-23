// Code generated by mockery v2.46.3. DO NOT EDIT.

package repository

import (
	context "context"

	entity "picadosYa/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

// GetUserByEmail provides a mock function with given fields: ctx, email
func (_m *MockRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByEmail")
	}

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.User); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveUserRole provides a mock function with given fields: ctx, userID, roleID
func (_m *MockRepository) RemoveUserRole(ctx context.Context, userID int64, roleID int64) error {
	ret := _m.Called(ctx, userID, roleID)

	if len(ret) == 0 {
		panic("no return value specified for RemoveUserRole")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = rf(ctx, userID, roleID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveUser provides a mock function with given fields: ctx, email, name, lastname, password, telephone, profile_photo
func (_m *MockRepository) SaveUser(ctx context.Context, email string, name string, lastname string, password string, telephone string, profile_photo string) error {
	ret := _m.Called(ctx, email, name, lastname, password, telephone, profile_photo)

	if len(ret) == 0 {
		panic("no return value specified for SaveUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, string, string) error); ok {
		r0 = rf(ctx, email, name, lastname, password, telephone, profile_photo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveUserRole provides a mock function with given fields: ctx, userID, roleID
func (_m *MockRepository) SaveUserRole(ctx context.Context, userID int64, roleID int64) error {
	ret := _m.Called(ctx, userID, roleID)

	if len(ret) == 0 {
		panic("no return value specified for SaveUserRole")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = rf(ctx, userID, roleID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
