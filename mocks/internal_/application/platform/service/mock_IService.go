// Code generated by mockery v2.52.4. DO NOT EDIT.

package service

import (
	context "context"

	service "github.com/radityacandra/ecommerce-connector-cli/internal/application/platform/service"
	mock "github.com/stretchr/testify/mock"
)

// MockIService is an autogenerated mock type for the IService type
type MockIService struct {
	mock.Mock
}

type MockIService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIService) EXPECT() *MockIService_Expecter {
	return &MockIService_Expecter{mock: &_m.Mock}
}

// Set provides a mock function with given fields: _a0, _a1
func (_m *MockIService) Set(_a0 context.Context, _a1 service.SetPlatformInput) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Set")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, service.SetPlatformInput) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIService_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type MockIService_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 service.SetPlatformInput
func (_e *MockIService_Expecter) Set(_a0 interface{}, _a1 interface{}) *MockIService_Set_Call {
	return &MockIService_Set_Call{Call: _e.mock.On("Set", _a0, _a1)}
}

func (_c *MockIService_Set_Call) Run(run func(_a0 context.Context, _a1 service.SetPlatformInput)) *MockIService_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(service.SetPlatformInput))
	})
	return _c
}

func (_c *MockIService_Set_Call) Return(_a0 error) *MockIService_Set_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIService_Set_Call) RunAndReturn(run func(context.Context, service.SetPlatformInput) error) *MockIService_Set_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockIService creates a new instance of MockIService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIService {
	mock := &MockIService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
