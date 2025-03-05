// Code generated by mockery v2.52.4. DO NOT EDIT.

package repository

import (
	context "context"

	model "github.com/radityacandra/ecommerce-connector-cli/internal/application/product/model"
	mock "github.com/stretchr/testify/mock"
)

// MockIRepository is an autogenerated mock type for the IRepository type
type MockIRepository struct {
	mock.Mock
}

type MockIRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIRepository) EXPECT() *MockIRepository_Expecter {
	return &MockIRepository_Expecter{mock: &_m.Mock}
}

// FindProductByEanCode provides a mock function with given fields: ctx, eanCode, userId
func (_m *MockIRepository) FindProductByEanCode(ctx context.Context, eanCode string, userId string) (*model.Product, error) {
	ret := _m.Called(ctx, eanCode, userId)

	if len(ret) == 0 {
		panic("no return value specified for FindProductByEanCode")
	}

	var r0 *model.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*model.Product, error)); ok {
		return rf(ctx, eanCode, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *model.Product); ok {
		r0 = rf(ctx, eanCode, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, eanCode, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIRepository_FindProductByEanCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindProductByEanCode'
type MockIRepository_FindProductByEanCode_Call struct {
	*mock.Call
}

// FindProductByEanCode is a helper method to define mock.On call
//   - ctx context.Context
//   - eanCode string
//   - userId string
func (_e *MockIRepository_Expecter) FindProductByEanCode(ctx interface{}, eanCode interface{}, userId interface{}) *MockIRepository_FindProductByEanCode_Call {
	return &MockIRepository_FindProductByEanCode_Call{Call: _e.mock.On("FindProductByEanCode", ctx, eanCode, userId)}
}

func (_c *MockIRepository_FindProductByEanCode_Call) Run(run func(ctx context.Context, eanCode string, userId string)) *MockIRepository_FindProductByEanCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockIRepository_FindProductByEanCode_Call) Return(_a0 *model.Product, _a1 error) *MockIRepository_FindProductByEanCode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIRepository_FindProductByEanCode_Call) RunAndReturn(run func(context.Context, string, string) (*model.Product, error)) *MockIRepository_FindProductByEanCode_Call {
	_c.Call.Return(run)
	return _c
}

// Insert provides a mock function with given fields: ctx, product
func (_m *MockIRepository) Insert(ctx context.Context, product *model.Product) error {
	ret := _m.Called(ctx, product)

	if len(ret) == 0 {
		panic("no return value specified for Insert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Product) error); ok {
		r0 = rf(ctx, product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIRepository_Insert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Insert'
type MockIRepository_Insert_Call struct {
	*mock.Call
}

// Insert is a helper method to define mock.On call
//   - ctx context.Context
//   - product *model.Product
func (_e *MockIRepository_Expecter) Insert(ctx interface{}, product interface{}) *MockIRepository_Insert_Call {
	return &MockIRepository_Insert_Call{Call: _e.mock.On("Insert", ctx, product)}
}

func (_c *MockIRepository_Insert_Call) Run(run func(ctx context.Context, product *model.Product)) *MockIRepository_Insert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.Product))
	})
	return _c
}

func (_c *MockIRepository_Insert_Call) Return(_a0 error) *MockIRepository_Insert_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIRepository_Insert_Call) RunAndReturn(run func(context.Context, *model.Product) error) *MockIRepository_Insert_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockIRepository creates a new instance of MockIRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIRepository {
	mock := &MockIRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
