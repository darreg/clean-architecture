// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/alrund/yp-1-project/internal/domain/entity"
	mock "github.com/stretchr/testify/mock"

	sql "database/sql"
)

// OrderRepository is an autogenerated mock type for the OrderRepository type
type OrderRepository struct {
	mock.Mock
}

type OrderRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *OrderRepository) EXPECT() *OrderRepository_Expecter {
	return &OrderRepository_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: ctx, order
func (_m *OrderRepository) Add(ctx context.Context, order *entity.Order) error {
	ret := _m.Called(ctx, order)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Order) error); ok {
		r0 = rf(ctx, order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrderRepository_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type OrderRepository_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//  - ctx context.Context
//  - order *entity.Order
func (_e *OrderRepository_Expecter) Add(ctx interface{}, order interface{}) *OrderRepository_Add_Call {
	return &OrderRepository_Add_Call{Call: _e.mock.On("Add", ctx, order)}
}

func (_c *OrderRepository_Add_Call) Run(run func(ctx context.Context, order *entity.Order)) *OrderRepository_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.Order))
	})
	return _c
}

func (_c *OrderRepository_Add_Call) Return(_a0 error) *OrderRepository_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

// Change provides a mock function with given fields: ctx, order
func (_m *OrderRepository) Change(ctx context.Context, order *entity.Order) error {
	ret := _m.Called(ctx, order)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Order) error); ok {
		r0 = rf(ctx, order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrderRepository_Change_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Change'
type OrderRepository_Change_Call struct {
	*mock.Call
}

// Change is a helper method to define mock.On call
//  - ctx context.Context
//  - order *entity.Order
func (_e *OrderRepository_Expecter) Change(ctx interface{}, order interface{}) *OrderRepository_Change_Call {
	return &OrderRepository_Change_Call{Call: _e.mock.On("Change", ctx, order)}
}

func (_c *OrderRepository_Change_Call) Run(run func(ctx context.Context, order *entity.Order)) *OrderRepository_Change_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.Order))
	})
	return _c
}

func (_c *OrderRepository_Change_Call) Return(_a0 error) *OrderRepository_Change_Call {
	_c.Call.Return(_a0)
	return _c
}

// ExecContext provides a mock function with given fields: ctx, db, query, args
func (_m *OrderRepository) ExecContext(ctx context.Context, db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, db, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 sql.Result
	if rf, ok := ret.Get(0).(func(context.Context, *sql.DB, string, ...interface{}) sql.Result); ok {
		r0 = rf(ctx, db, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sql.DB, string, ...interface{}) error); ok {
		r1 = rf(ctx, db, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrderRepository_ExecContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecContext'
type OrderRepository_ExecContext_Call struct {
	*mock.Call
}

// ExecContext is a helper method to define mock.On call
//  - ctx context.Context
//  - db *sql.DB
//  - query string
//  - args ...interface{}
func (_e *OrderRepository_Expecter) ExecContext(ctx interface{}, db interface{}, query interface{}, args ...interface{}) *OrderRepository_ExecContext_Call {
	return &OrderRepository_ExecContext_Call{Call: _e.mock.On("ExecContext",
		append([]interface{}{ctx, db, query}, args...)...)}
}

func (_c *OrderRepository_ExecContext_Call) Run(run func(ctx context.Context, db *sql.DB, query string, args ...interface{})) *OrderRepository_ExecContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(*sql.DB), args[2].(string), variadicArgs...)
	})
	return _c
}

func (_c *OrderRepository_ExecContext_Call) Return(_a0 sql.Result, _a1 error) *OrderRepository_ExecContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Get provides a mock function with given fields: ctx, number
func (_m *OrderRepository) Get(ctx context.Context, number string) (*entity.Order, error) {
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

// OrderRepository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type OrderRepository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//  - ctx context.Context
//  - number string
func (_e *OrderRepository_Expecter) Get(ctx interface{}, number interface{}) *OrderRepository_Get_Call {
	return &OrderRepository_Get_Call{Call: _e.mock.On("Get", ctx, number)}
}

func (_c *OrderRepository_Get_Call) Run(run func(ctx context.Context, number string)) *OrderRepository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *OrderRepository_Get_Call) Return(_a0 *entity.Order, _a1 error) *OrderRepository_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetAllByUser provides a mock function with given fields: ctx, user
func (_m *OrderRepository) GetAllByUser(ctx context.Context, user *entity.User) ([]*entity.Order, error) {
	ret := _m.Called(ctx, user)

	var r0 []*entity.Order
	if rf, ok := ret.Get(0).(func(context.Context, *entity.User) []*entity.Order); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrderRepository_GetAllByUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllByUser'
type OrderRepository_GetAllByUser_Call struct {
	*mock.Call
}

// GetAllByUser is a helper method to define mock.On call
//  - ctx context.Context
//  - user *entity.User
func (_e *OrderRepository_Expecter) GetAllByUser(ctx interface{}, user interface{}) *OrderRepository_GetAllByUser_Call {
	return &OrderRepository_GetAllByUser_Call{Call: _e.mock.On("GetAllByUser", ctx, user)}
}

func (_c *OrderRepository_GetAllByUser_Call) Run(run func(ctx context.Context, user *entity.User)) *OrderRepository_GetAllByUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.User))
	})
	return _c
}

func (_c *OrderRepository_GetAllByUser_Call) Return(_a0 []*entity.Order, _a1 error) *OrderRepository_GetAllByUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// QueryContext provides a mock function with given fields: ctx, db, query, args
func (_m *OrderRepository) QueryContext(ctx context.Context, db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, db, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *sql.Rows
	if rf, ok := ret.Get(0).(func(context.Context, *sql.DB, string, ...interface{}) *sql.Rows); ok {
		r0 = rf(ctx, db, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Rows)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sql.DB, string, ...interface{}) error); ok {
		r1 = rf(ctx, db, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrderRepository_QueryContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryContext'
type OrderRepository_QueryContext_Call struct {
	*mock.Call
}

// QueryContext is a helper method to define mock.On call
//  - ctx context.Context
//  - db *sql.DB
//  - query string
//  - args ...interface{}
func (_e *OrderRepository_Expecter) QueryContext(ctx interface{}, db interface{}, query interface{}, args ...interface{}) *OrderRepository_QueryContext_Call {
	return &OrderRepository_QueryContext_Call{Call: _e.mock.On("QueryContext",
		append([]interface{}{ctx, db, query}, args...)...)}
}

func (_c *OrderRepository_QueryContext_Call) Run(run func(ctx context.Context, db *sql.DB, query string, args ...interface{})) *OrderRepository_QueryContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(*sql.DB), args[2].(string), variadicArgs...)
	})
	return _c
}

func (_c *OrderRepository_QueryContext_Call) Return(_a0 *sql.Rows, _a1 error) *OrderRepository_QueryContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// QueryRowContext provides a mock function with given fields: ctx, db, query, args
func (_m *OrderRepository) QueryRowContext(ctx context.Context, db *sql.DB, query string, args ...interface{}) *sql.Row {
	var _ca []interface{}
	_ca = append(_ca, ctx, db, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *sql.Row
	if rf, ok := ret.Get(0).(func(context.Context, *sql.DB, string, ...interface{}) *sql.Row); ok {
		r0 = rf(ctx, db, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Row)
		}
	}

	return r0
}

// OrderRepository_QueryRowContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryRowContext'
type OrderRepository_QueryRowContext_Call struct {
	*mock.Call
}

// QueryRowContext is a helper method to define mock.On call
//  - ctx context.Context
//  - db *sql.DB
//  - query string
//  - args ...interface{}
func (_e *OrderRepository_Expecter) QueryRowContext(ctx interface{}, db interface{}, query interface{}, args ...interface{}) *OrderRepository_QueryRowContext_Call {
	return &OrderRepository_QueryRowContext_Call{Call: _e.mock.On("QueryRowContext",
		append([]interface{}{ctx, db, query}, args...)...)}
}

func (_c *OrderRepository_QueryRowContext_Call) Run(run func(ctx context.Context, db *sql.DB, query string, args ...interface{})) *OrderRepository_QueryRowContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(*sql.DB), args[2].(string), variadicArgs...)
	})
	return _c
}

func (_c *OrderRepository_QueryRowContext_Call) Return(_a0 *sql.Row) *OrderRepository_QueryRowContext_Call {
	_c.Call.Return(_a0)
	return _c
}

// Remove provides a mock function with given fields: ctx, number
func (_m *OrderRepository) Remove(ctx context.Context, number string) error {
	ret := _m.Called(ctx, number)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, number)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrderRepository_Remove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Remove'
type OrderRepository_Remove_Call struct {
	*mock.Call
}

// Remove is a helper method to define mock.On call
//  - ctx context.Context
//  - number string
func (_e *OrderRepository_Expecter) Remove(ctx interface{}, number interface{}) *OrderRepository_Remove_Call {
	return &OrderRepository_Remove_Call{Call: _e.mock.On("Remove", ctx, number)}
}

func (_c *OrderRepository_Remove_Call) Run(run func(ctx context.Context, number string)) *OrderRepository_Remove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *OrderRepository_Remove_Call) Return(_a0 error) *OrderRepository_Remove_Call {
	_c.Call.Return(_a0)
	return _c
}

// WithinTransaction provides a mock function with given fields: ctx, tFunc
func (_m *OrderRepository) WithinTransaction(ctx context.Context, tFunc func(context.Context) error) error {
	ret := _m.Called(ctx, tFunc)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(context.Context) error) error); ok {
		r0 = rf(ctx, tFunc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrderRepository_WithinTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithinTransaction'
type OrderRepository_WithinTransaction_Call struct {
	*mock.Call
}

// WithinTransaction is a helper method to define mock.On call
//  - ctx context.Context
//  - tFunc func(context.Context) error
func (_e *OrderRepository_Expecter) WithinTransaction(ctx interface{}, tFunc interface{}) *OrderRepository_WithinTransaction_Call {
	return &OrderRepository_WithinTransaction_Call{Call: _e.mock.On("WithinTransaction", ctx, tFunc)}
}

func (_c *OrderRepository_WithinTransaction_Call) Run(run func(ctx context.Context, tFunc func(context.Context) error)) *OrderRepository_WithinTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(func(context.Context) error))
	})
	return _c
}

func (_c *OrderRepository_WithinTransaction_Call) Return(_a0 error) *OrderRepository_WithinTransaction_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewOrderRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewOrderRepository creates a new instance of OrderRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOrderRepository(t mockConstructorTestingTNewOrderRepository) *OrderRepository {
	mock := &OrderRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}