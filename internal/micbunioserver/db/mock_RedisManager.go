// Code generated by mockery v2.42.1. DO NOT EDIT.

package db

import (
	context "context"
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// MockRedisManager is an autogenerated mock type for the RedisManager type
type MockRedisManager struct {
	mock.Mock
}

type MockRedisManager_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRedisManager) EXPECT() *MockRedisManager_Expecter {
	return &MockRedisManager_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: ctx, key
func (_m *MockRedisManager) Delete(ctx context.Context, key string) {
	_m.Called(ctx, key)
}

// MockRedisManager_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockRedisManager_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
func (_e *MockRedisManager_Expecter) Delete(ctx interface{}, key interface{}) *MockRedisManager_Delete_Call {
	return &MockRedisManager_Delete_Call{Call: _e.mock.On("Delete", ctx, key)}
}

func (_c *MockRedisManager_Delete_Call) Run(run func(ctx context.Context, key string)) *MockRedisManager_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRedisManager_Delete_Call) Return() *MockRedisManager_Delete_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockRedisManager_Delete_Call) RunAndReturn(run func(context.Context, string)) *MockRedisManager_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, key, value
func (_m *MockRedisManager) Get(ctx context.Context, key string, value interface{}) {
	_m.Called(ctx, key, value)
}

// MockRedisManager_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockRedisManager_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - value interface{}
func (_e *MockRedisManager_Expecter) Get(ctx interface{}, key interface{}, value interface{}) *MockRedisManager_Get_Call {
	return &MockRedisManager_Get_Call{Call: _e.mock.On("Get", ctx, key, value)}
}

func (_c *MockRedisManager_Get_Call) Run(run func(ctx context.Context, key string, value interface{})) *MockRedisManager_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(interface{}))
	})
	return _c
}

func (_c *MockRedisManager_Get_Call) Return() *MockRedisManager_Get_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockRedisManager_Get_Call) RunAndReturn(run func(context.Context, string, interface{})) *MockRedisManager_Get_Call {
	_c.Call.Return(run)
	return _c
}

// SetEx provides a mock function with given fields: ctx, key, value, duration
func (_m *MockRedisManager) SetEx(ctx context.Context, key string, value interface{}, duration time.Duration) {
	_m.Called(ctx, key, value, duration)
}

// MockRedisManager_SetEx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetEx'
type MockRedisManager_SetEx_Call struct {
	*mock.Call
}

// SetEx is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - value interface{}
//   - duration time.Duration
func (_e *MockRedisManager_Expecter) SetEx(ctx interface{}, key interface{}, value interface{}, duration interface{}) *MockRedisManager_SetEx_Call {
	return &MockRedisManager_SetEx_Call{Call: _e.mock.On("SetEx", ctx, key, value, duration)}
}

func (_c *MockRedisManager_SetEx_Call) Run(run func(ctx context.Context, key string, value interface{}, duration time.Duration)) *MockRedisManager_SetEx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(interface{}), args[3].(time.Duration))
	})
	return _c
}

func (_c *MockRedisManager_SetEx_Call) Return() *MockRedisManager_SetEx_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockRedisManager_SetEx_Call) RunAndReturn(run func(context.Context, string, interface{}, time.Duration)) *MockRedisManager_SetEx_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRedisManager creates a new instance of MockRedisManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRedisManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRedisManager {
	mock := &MockRedisManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
