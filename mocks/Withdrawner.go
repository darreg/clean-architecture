// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/alrund/yp-1-project/internal/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// Withdrawner is an autogenerated mock type for the Withdrawner type
type Withdrawner struct {
	mock.Mock
}

type Withdrawner_Expecter struct {
	mock *mock.Mock
}

func (_m *Withdrawner) EXPECT() *Withdrawner_Expecter {
	return &Withdrawner_Expecter{mock: &_m.Mock}
}

// GetWithdrawn provides a mock function with given fields: ctx, user
func (_m *Withdrawner) GetWithdrawn(ctx context.Context, user *entity.User) (float32, error) {
	ret := _m.Called(ctx, user)

	var r0 float32
	if rf, ok := ret.Get(0).(func(context.Context, *entity.User) float32); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(float32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Withdrawner_GetWithdrawn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetWithdrawn'
type Withdrawner_GetWithdrawn_Call struct {
	*mock.Call
}

// GetWithdrawn is a helper method to define mock.On call
//  - ctx context.Context
//  - user *entity.User
func (_e *Withdrawner_Expecter) GetWithdrawn(ctx interface{}, user interface{}) *Withdrawner_GetWithdrawn_Call {
	return &Withdrawner_GetWithdrawn_Call{Call: _e.mock.On("GetWithdrawn", ctx, user)}
}

func (_c *Withdrawner_GetWithdrawn_Call) Run(run func(ctx context.Context, user *entity.User)) *Withdrawner_GetWithdrawn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.User))
	})
	return _c
}

func (_c *Withdrawner_GetWithdrawn_Call) Return(_a0 float32, _a1 error) *Withdrawner_GetWithdrawn_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewWithdrawner interface {
	mock.TestingT
	Cleanup(func())
}

// NewWithdrawner creates a new instance of Withdrawner. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWithdrawner(t mockConstructorTestingTNewWithdrawner) *Withdrawner {
	mock := &Withdrawner{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
