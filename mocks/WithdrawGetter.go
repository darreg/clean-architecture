// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/alrund/yp-1-project/internal/domain/entity"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// WithdrawGetter is an autogenerated mock type for the WithdrawGetter type
type WithdrawGetter struct {
	mock.Mock
}

type WithdrawGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *WithdrawGetter) EXPECT() *WithdrawGetter_Expecter {
	return &WithdrawGetter_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, withdrawID
func (_m *WithdrawGetter) Get(ctx context.Context, withdrawID uuid.UUID) (*entity.Withdraw, error) {
	ret := _m.Called(ctx, withdrawID)

	var r0 *entity.Withdraw
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *entity.Withdraw); ok {
		r0 = rf(ctx, withdrawID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Withdraw)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, withdrawID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithdrawGetter_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type WithdrawGetter_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//  - ctx context.Context
//  - withdrawID uuid.UUID
func (_e *WithdrawGetter_Expecter) Get(ctx interface{}, withdrawID interface{}) *WithdrawGetter_Get_Call {
	return &WithdrawGetter_Get_Call{Call: _e.mock.On("Get", ctx, withdrawID)}
}

func (_c *WithdrawGetter_Get_Call) Run(run func(ctx context.Context, withdrawID uuid.UUID)) *WithdrawGetter_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *WithdrawGetter_Get_Call) Return(_a0 *entity.Withdraw, _a1 error) *WithdrawGetter_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewWithdrawGetter interface {
	mock.TestingT
	Cleanup(func())
}

// NewWithdrawGetter creates a new instance of WithdrawGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWithdrawGetter(t mockConstructorTestingTNewWithdrawGetter) *WithdrawGetter {
	mock := &WithdrawGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
