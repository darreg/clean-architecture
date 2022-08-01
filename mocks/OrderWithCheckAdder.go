// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/alrund/yp-1-project/internal/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// OrderWithCheckAdder is an autogenerated mock type for the OrderWithCheckAdder type
type OrderWithCheckAdder struct {
	mock.Mock
}

type OrderWithCheckAdder_Expecter struct {
	mock *mock.Mock
}

func (_m *OrderWithCheckAdder) EXPECT() *OrderWithCheckAdder_Expecter {
	return &OrderWithCheckAdder_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: ctx, order
func (_m *OrderWithCheckAdder) Add(ctx context.Context, order *entity.Order) error {
	ret := _m.Called(ctx, order)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Order) error); ok {
		r0 = rf(ctx, order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrderWithCheckAdder_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type OrderWithCheckAdder_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//  - ctx context.Context
//  - order *entity.Order
func (_e *OrderWithCheckAdder_Expecter) Add(ctx interface{}, order interface{}) *OrderWithCheckAdder_Add_Call {
	return &OrderWithCheckAdder_Add_Call{Call: _e.mock.On("Add", ctx, order)}
}

func (_c *OrderWithCheckAdder_Add_Call) Run(run func(ctx context.Context, order *entity.Order)) *OrderWithCheckAdder_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.Order))
	})
	return _c
}

func (_c *OrderWithCheckAdder_Add_Call) Return(_a0 error) *OrderWithCheckAdder_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

// Get provides a mock function with given fields: ctx, number
func (_m *OrderWithCheckAdder) Get(ctx context.Context, number string) (*entity.Order, error) {
	ret := _m.Called(ctx, number)

	var r0 *entity.Order
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Order); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrderWithCheckAdder_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type OrderWithCheckAdder_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//  - ctx context.Context
//  - number string
func (_e *OrderWithCheckAdder_Expecter) Get(ctx interface{}, number interface{}) *OrderWithCheckAdder_Get_Call {
	return &OrderWithCheckAdder_Get_Call{Call: _e.mock.On("Get", ctx, number)}
}

func (_c *OrderWithCheckAdder_Get_Call) Run(run func(ctx context.Context, number string)) *OrderWithCheckAdder_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *OrderWithCheckAdder_Get_Call) Return(_a0 *entity.Order, _a1 error) *OrderWithCheckAdder_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewOrderWithCheckAdder interface {
	mock.TestingT
	Cleanup(func())
}

// NewOrderWithCheckAdder creates a new instance of OrderWithCheckAdder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOrderWithCheckAdder(t mockConstructorTestingTNewOrderWithCheckAdder) *OrderWithCheckAdder {
	mock := &OrderWithCheckAdder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}