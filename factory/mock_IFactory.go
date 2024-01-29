// Code generated by mockery v2.40.1. DO NOT EDIT.

package factory

import (
	mock "github.com/stretchr/testify/mock"
	broker "github.com/surendratiwari3/paota/broker"

	task "github.com/surendratiwari3/paota/task"
)

// MockIFactory is an autogenerated mock type for the IFactory type
type MockIFactory struct {
	mock.Mock
}

// CreateBroker provides a mock function with given fields:
func (_m *MockIFactory) CreateBroker() (broker.Broker, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CreateBroker")
	}

	var r0 broker.Broker
	var r1 error
	if rf, ok := ret.Get(0).(func() (broker.Broker, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() broker.Broker); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(broker.Broker)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateStore provides a mock function with given fields:
func (_m *MockIFactory) CreateStore() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CreateStore")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateTaskRegistrar provides a mock function with given fields:
func (_m *MockIFactory) CreateTaskRegistrar() task.TaskRegistrarInterface {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CreateTaskRegistrar")
	}

	var r0 task.TaskRegistrarInterface
	if rf, ok := ret.Get(0).(func() task.TaskRegistrarInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(task.TaskRegistrarInterface)
		}
	}

	return r0
}

// NewMockIFactory creates a new instance of MockIFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIFactory(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIFactory {
	mock := &MockIFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}