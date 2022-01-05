// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import users "currency-exchange/business/users"

// UserUseCaseInterface is an autogenerated mock type for the UserUseCaseInterface type
type UserUseCaseInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, register
func (_m *UserUseCaseInterface) Create(ctx context.Context, register *users.Domain) (users.Domain, error) {
	ret := _m.Called(ctx, register)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *users.Domain) users.Domain); ok {
		r0 = rf(ctx, register)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *users.Domain) error); ok {
		r1 = rf(ctx, register)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *UserUseCaseInterface) Delete(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *UserUseCaseInterface) GetAll(ctx context.Context) ([]users.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []users.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []users.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: ctx, id
func (_m *UserUseCaseInterface) GetById(ctx context.Context, id int) (users.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int) users.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: domain, ctx
func (_m *UserUseCaseInterface) Login(domain users.Domain, ctx context.Context) (users.Domain, error) {
	ret := _m.Called(domain, ctx)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(users.Domain, context.Context) users.Domain); ok {
		r0 = rf(domain, ctx)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(users.Domain, context.Context) error); ok {
		r1 = rf(domain, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}