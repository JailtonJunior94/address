// Code generated by mockery v2.24.0. DO NOT EDIT.

package serviceMocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// IHttpClient is an autogenerated mock type for the IHttpClient type
type IHttpClient struct {
	mock.Mock
}

// Do provides a mock function with given fields: req
func (_m *IHttpClient) Do(req *http.Request) (*http.Response, error) {
	ret := _m.Called(req)

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(*http.Request) (*http.Response, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*http.Request) *http.Response); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(*http.Request) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIHttpClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewIHttpClient creates a new instance of IHttpClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIHttpClient(t mockConstructorTestingTNewIHttpClient) *IHttpClient {
	mock := &IHttpClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
