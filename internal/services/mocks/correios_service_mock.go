// Code generated by mockery v2.24.0. DO NOT EDIT.

package serviceMocks

import (
	dtos "github.com/jailtonjunior94/address/internal/dtos"

	mock "github.com/stretchr/testify/mock"
)

// CorreiosService is an autogenerated mock type for the CorreiosService type
type CorreiosService struct {
	mock.Mock
}

// AddressByCEP provides a mock function with given fields: cep
func (_m *CorreiosService) AddressByCEP(cep string) (*dtos.AddressResponse, error) {
	ret := _m.Called(cep)

	var r0 *dtos.AddressResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*dtos.AddressResponse, error)); ok {
		return rf(cep)
	}
	if rf, ok := ret.Get(0).(func(string) *dtos.AddressResponse); ok {
		r0 = rf(cep)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dtos.AddressResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(cep)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCorreiosService interface {
	mock.TestingT
	Cleanup(func())
}

// NewCorreiosService creates a new instance of CorreiosService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCorreiosService(t mockConstructorTestingTNewCorreiosService) *CorreiosService {
	mock := &CorreiosService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
