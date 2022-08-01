// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/alrund/yp-1-project/internal/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// OrderChanger is an autogenerated mock type for the OrderChanger type
type OrderChanger struct {
	mock.Mock
}

type OrderChanger_Expecter struct {
	mock *mock.Mock
}

func (_m *OrderChanger) EXPECT() *OrderChanger_Expecter {
	return &OrderChanger_Expecter{mock: &_m.Mock}
}

// Change provides a mock function with given fields: ctx, order
func (_m *OrderChanger) Change(ctx context.Context, order *entity.Order) error {
	ret := _m.Called(ctx, order)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Order) error); ok {
		r0 = rf(ctx, order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrderChanger_Change_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Change'
type OrderChanger_Change_Call struct {
	*mock.Call
}

// Change is a helper method to define mock.On call
//  - ctx context.Context
//  - order *entity.Order
func (_e *OrderChanger_Expecter) Change(ctx interface{}, order interface{}) *OrderChanger_Change_Call {
	return &OrderChanger_Change_Call{Call: _e.mock.On("Change", ctx, order)}
}

func (_c *OrderChanger_Change_Call) Run(run func(ctx context.Context, order *entity.Order)) *OrderChanger_Change_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.Order))
	})
	return _c
}

func (_c *OrderChanger_Change_Call) Return(_a0 error) *OrderChanger_Change_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewOrderChanger interface {
	mock.TestingT
	Cleanup(func())
}

// NewOrderChanger creates a new instance of OrderChanger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOrderChanger(t mockConstructorTestingTNewOrderChanger) *OrderChanger {
	mock := &OrderChanger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}