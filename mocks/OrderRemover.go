// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// OrderRemover is an autogenerated mock type for the OrderRemover type
type OrderRemover struct {
	mock.Mock
}

type OrderRemover_Expecter struct {
	mock *mock.Mock
}

func (_m *OrderRemover) EXPECT() *OrderRemover_Expecter {
	return &OrderRemover_Expecter{mock: &_m.Mock}
}

// Remove provides a mock function with given fields: ctx, number
func (_m *OrderRemover) Remove(ctx context.Context, number string) error {
	ret := _m.Called(ctx, number)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, number)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrderRemover_Remove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Remove'
type OrderRemover_Remove_Call struct {
	*mock.Call
}

// Remove is a helper method to define mock.On call
//  - ctx context.Context
//  - number string
func (_e *OrderRemover_Expecter) Remove(ctx interface{}, number interface{}) *OrderRemover_Remove_Call {
	return &OrderRemover_Remove_Call{Call: _e.mock.On("Remove", ctx, number)}
}

func (_c *OrderRemover_Remove_Call) Run(run func(ctx context.Context, number string)) *OrderRemover_Remove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *OrderRemover_Remove_Call) Return(_a0 error) *OrderRemover_Remove_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewOrderRemover interface {
	mock.TestingT
	Cleanup(func())
}

// NewOrderRemover creates a new instance of OrderRemover. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOrderRemover(t mockConstructorTestingTNewOrderRemover) *OrderRemover {
	mock := &OrderRemover{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
