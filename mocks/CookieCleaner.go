// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// CookieCleaner is an autogenerated mock type for the CookieCleaner type
type CookieCleaner struct {
	mock.Mock
}

type CookieCleaner_Expecter struct {
	mock *mock.Mock
}

func (_m *CookieCleaner) EXPECT() *CookieCleaner_Expecter {
	return &CookieCleaner_Expecter{mock: &_m.Mock}
}

// ClearCookie provides a mock function with given fields: name, w
func (_m *CookieCleaner) ClearCookie(name string, w http.ResponseWriter) {
	_m.Called(name, w)
}

// CookieCleaner_ClearCookie_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClearCookie'
type CookieCleaner_ClearCookie_Call struct {
	*mock.Call
}

// ClearCookie is a helper method to define mock.On call
//  - name string
//  - w http.ResponseWriter
func (_e *CookieCleaner_Expecter) ClearCookie(name interface{}, w interface{}) *CookieCleaner_ClearCookie_Call {
	return &CookieCleaner_ClearCookie_Call{Call: _e.mock.On("ClearCookie", name, w)}
}

func (_c *CookieCleaner_ClearCookie_Call) Run(run func(name string, w http.ResponseWriter)) *CookieCleaner_ClearCookie_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(http.ResponseWriter))
	})
	return _c
}

func (_c *CookieCleaner_ClearCookie_Call) Return() *CookieCleaner_ClearCookie_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewCookieCleaner interface {
	mock.TestingT
	Cleanup(func())
}

// NewCookieCleaner creates a new instance of CookieCleaner. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCookieCleaner(t mockConstructorTestingTNewCookieCleaner) *CookieCleaner {
	mock := &CookieCleaner{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
