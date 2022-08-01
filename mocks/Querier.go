// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	sql "database/sql"
)

// Querier is an autogenerated mock type for the Querier type
type Querier struct {
	mock.Mock
}

type Querier_Expecter struct {
	mock *mock.Mock
}

func (_m *Querier) EXPECT() *Querier_Expecter {
	return &Querier_Expecter{mock: &_m.Mock}
}

// QueryContext provides a mock function with given fields: ctx, db, query, args
func (_m *Querier) QueryContext(ctx context.Context, db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
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

// Querier_QueryContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryContext'
type Querier_QueryContext_Call struct {
	*mock.Call
}

// QueryContext is a helper method to define mock.On call
//  - ctx context.Context
//  - db *sql.DB
//  - query string
//  - args ...interface{}
func (_e *Querier_Expecter) QueryContext(ctx interface{}, db interface{}, query interface{}, args ...interface{}) *Querier_QueryContext_Call {
	return &Querier_QueryContext_Call{Call: _e.mock.On("QueryContext",
		append([]interface{}{ctx, db, query}, args...)...)}
}

func (_c *Querier_QueryContext_Call) Run(run func(ctx context.Context, db *sql.DB, query string, args ...interface{})) *Querier_QueryContext_Call {
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

func (_c *Querier_QueryContext_Call) Return(_a0 *sql.Rows, _a1 error) *Querier_QueryContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewQuerier interface {
	mock.TestingT
	Cleanup(func())
}

// NewQuerier creates a new instance of Querier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewQuerier(t mockConstructorTestingTNewQuerier) *Querier {
	mock := &Querier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
