// Code generated by mockery. DO NOT EDIT.

package wait_test

import (
	context "context"

	container "github.com/docker/docker/api/types/container"

	exec "github.com/testcontainers/testcontainers-go/exec"

	io "io"

	mock "github.com/stretchr/testify/mock"

	nat "github.com/docker/go-connections/nat"
)

// mockStrategyTarget is an autogenerated mock type for the StrategyTarget type
type mockStrategyTarget struct {
	mock.Mock
}

type mockStrategyTarget_Expecter struct {
	mock *mock.Mock
}

func (_m *mockStrategyTarget) EXPECT() *mockStrategyTarget_Expecter {
	return &mockStrategyTarget_Expecter{mock: &_m.Mock}
}

// CopyFileFromContainer provides a mock function with given fields: ctx, filePath
func (_m *mockStrategyTarget) CopyFileFromContainer(ctx context.Context, filePath string) (io.ReadCloser, error) {
	ret := _m.Called(ctx, filePath)

	if len(ret) == 0 {
		panic("no return value specified for CopyFileFromContainer")
	}

	var r0 io.ReadCloser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (io.ReadCloser, error)); ok {
		return rf(ctx, filePath)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) io.ReadCloser); ok {
		r0 = rf(ctx, filePath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, filePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockStrategyTarget_CopyFileFromContainer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CopyFileFromContainer'
type mockStrategyTarget_CopyFileFromContainer_Call struct {
	*mock.Call
}

// CopyFileFromContainer is a helper method to define mock.On call
//   - ctx context.Context
//   - filePath string
func (_e *mockStrategyTarget_Expecter) CopyFileFromContainer(ctx interface{}, filePath interface{}) *mockStrategyTarget_CopyFileFromContainer_Call {
	return &mockStrategyTarget_CopyFileFromContainer_Call{Call: _e.mock.On("CopyFileFromContainer", ctx, filePath)}
}

func (_c *mockStrategyTarget_CopyFileFromContainer_Call) Run(run func(ctx context.Context, filePath string)) *mockStrategyTarget_CopyFileFromContainer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *mockStrategyTarget_CopyFileFromContainer_Call) Return(_a0 io.ReadCloser, _a1 error) *mockStrategyTarget_CopyFileFromContainer_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockStrategyTarget_CopyFileFromContainer_Call) RunAndReturn(run func(context.Context, string) (io.ReadCloser, error)) *mockStrategyTarget_CopyFileFromContainer_Call {
	_c.Call.Return(run)
	return _c
}

// Exec provides a mock function with given fields: _a0, _a1, _a2
func (_m *mockStrategyTarget) Exec(_a0 context.Context, _a1 []string, _a2 ...exec.ProcessOption) (int, io.Reader, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 int
	var r1 io.Reader
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, []string, ...exec.ProcessOption) (int, io.Reader, error)); ok {
		return rf(_a0, _a1, _a2...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string, ...exec.ProcessOption) int); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string, ...exec.ProcessOption) io.Reader); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(io.Reader)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, []string, ...exec.ProcessOption) error); ok {
		r2 = rf(_a0, _a1, _a2...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// mockStrategyTarget_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type mockStrategyTarget_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 []string
//   - _a2 ...exec.ProcessOption
func (_e *mockStrategyTarget_Expecter) Exec(_a0 interface{}, _a1 interface{}, _a2 ...interface{}) *mockStrategyTarget_Exec_Call {
	return &mockStrategyTarget_Exec_Call{Call: _e.mock.On("Exec",
		append([]interface{}{_a0, _a1}, _a2...)...)}
}

func (_c *mockStrategyTarget_Exec_Call) Run(run func(_a0 context.Context, _a1 []string, _a2 ...exec.ProcessOption)) *mockStrategyTarget_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]exec.ProcessOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(exec.ProcessOption)
			}
		}
		run(args[0].(context.Context), args[1].([]string), variadicArgs...)
	})
	return _c
}

func (_c *mockStrategyTarget_Exec_Call) Return(_a0 int, _a1 io.Reader, _a2 error) *mockStrategyTarget_Exec_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *mockStrategyTarget_Exec_Call) RunAndReturn(run func(context.Context, []string, ...exec.ProcessOption) (int, io.Reader, error)) *mockStrategyTarget_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// Host provides a mock function with given fields: _a0
func (_m *mockStrategyTarget) Host(_a0 context.Context) (string, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Host")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (string, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockStrategyTarget_Host_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Host'
type mockStrategyTarget_Host_Call struct {
	*mock.Call
}

// Host is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *mockStrategyTarget_Expecter) Host(_a0 interface{}) *mockStrategyTarget_Host_Call {
	return &mockStrategyTarget_Host_Call{Call: _e.mock.On("Host", _a0)}
}

func (_c *mockStrategyTarget_Host_Call) Run(run func(_a0 context.Context)) *mockStrategyTarget_Host_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *mockStrategyTarget_Host_Call) Return(_a0 string, _a1 error) *mockStrategyTarget_Host_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockStrategyTarget_Host_Call) RunAndReturn(run func(context.Context) (string, error)) *mockStrategyTarget_Host_Call {
	_c.Call.Return(run)
	return _c
}

// Inspect provides a mock function with given fields: _a0
func (_m *mockStrategyTarget) Inspect(_a0 context.Context) (*container.InspectResponse, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Inspect")
	}

	var r0 *container.InspectResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*container.InspectResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *container.InspectResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*container.InspectResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockStrategyTarget_Inspect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Inspect'
type mockStrategyTarget_Inspect_Call struct {
	*mock.Call
}

// Inspect is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *mockStrategyTarget_Expecter) Inspect(_a0 interface{}) *mockStrategyTarget_Inspect_Call {
	return &mockStrategyTarget_Inspect_Call{Call: _e.mock.On("Inspect", _a0)}
}

func (_c *mockStrategyTarget_Inspect_Call) Run(run func(_a0 context.Context)) *mockStrategyTarget_Inspect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *mockStrategyTarget_Inspect_Call) Return(_a0 *container.InspectResponse, _a1 error) *mockStrategyTarget_Inspect_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockStrategyTarget_Inspect_Call) RunAndReturn(run func(context.Context) (*container.InspectResponse, error)) *mockStrategyTarget_Inspect_Call {
	_c.Call.Return(run)
	return _c
}

// Logs provides a mock function with given fields: _a0
func (_m *mockStrategyTarget) Logs(_a0 context.Context) (io.ReadCloser, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Logs")
	}

	var r0 io.ReadCloser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (io.ReadCloser, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) io.ReadCloser); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockStrategyTarget_Logs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Logs'
type mockStrategyTarget_Logs_Call struct {
	*mock.Call
}

// Logs is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *mockStrategyTarget_Expecter) Logs(_a0 interface{}) *mockStrategyTarget_Logs_Call {
	return &mockStrategyTarget_Logs_Call{Call: _e.mock.On("Logs", _a0)}
}

func (_c *mockStrategyTarget_Logs_Call) Run(run func(_a0 context.Context)) *mockStrategyTarget_Logs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *mockStrategyTarget_Logs_Call) Return(_a0 io.ReadCloser, _a1 error) *mockStrategyTarget_Logs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockStrategyTarget_Logs_Call) RunAndReturn(run func(context.Context) (io.ReadCloser, error)) *mockStrategyTarget_Logs_Call {
	_c.Call.Return(run)
	return _c
}

// MappedPort provides a mock function with given fields: _a0, _a1
func (_m *mockStrategyTarget) MappedPort(_a0 context.Context, _a1 nat.Port) (nat.Port, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for MappedPort")
	}

	var r0 nat.Port
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, nat.Port) (nat.Port, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, nat.Port) nat.Port); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(nat.Port)
	}

	if rf, ok := ret.Get(1).(func(context.Context, nat.Port) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockStrategyTarget_MappedPort_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MappedPort'
type mockStrategyTarget_MappedPort_Call struct {
	*mock.Call
}

// MappedPort is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 nat.Port
func (_e *mockStrategyTarget_Expecter) MappedPort(_a0 interface{}, _a1 interface{}) *mockStrategyTarget_MappedPort_Call {
	return &mockStrategyTarget_MappedPort_Call{Call: _e.mock.On("MappedPort", _a0, _a1)}
}

func (_c *mockStrategyTarget_MappedPort_Call) Run(run func(_a0 context.Context, _a1 nat.Port)) *mockStrategyTarget_MappedPort_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(nat.Port))
	})
	return _c
}

func (_c *mockStrategyTarget_MappedPort_Call) Return(_a0 nat.Port, _a1 error) *mockStrategyTarget_MappedPort_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockStrategyTarget_MappedPort_Call) RunAndReturn(run func(context.Context, nat.Port) (nat.Port, error)) *mockStrategyTarget_MappedPort_Call {
	_c.Call.Return(run)
	return _c
}

// Ports provides a mock function with given fields: ctx
func (_m *mockStrategyTarget) Ports(ctx context.Context) (nat.PortMap, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Ports")
	}

	var r0 nat.PortMap
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (nat.PortMap, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) nat.PortMap); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(nat.PortMap)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockStrategyTarget_Ports_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ports'
type mockStrategyTarget_Ports_Call struct {
	*mock.Call
}

// Ports is a helper method to define mock.On call
//   - ctx context.Context
func (_e *mockStrategyTarget_Expecter) Ports(ctx interface{}) *mockStrategyTarget_Ports_Call {
	return &mockStrategyTarget_Ports_Call{Call: _e.mock.On("Ports", ctx)}
}

func (_c *mockStrategyTarget_Ports_Call) Run(run func(ctx context.Context)) *mockStrategyTarget_Ports_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *mockStrategyTarget_Ports_Call) Return(_a0 nat.PortMap, _a1 error) *mockStrategyTarget_Ports_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockStrategyTarget_Ports_Call) RunAndReturn(run func(context.Context) (nat.PortMap, error)) *mockStrategyTarget_Ports_Call {
	_c.Call.Return(run)
	return _c
}

// State provides a mock function with given fields: _a0
func (_m *mockStrategyTarget) State(_a0 context.Context) (*container.State, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for State")
	}

	var r0 *container.State
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*container.State, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *container.State); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*container.State)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockStrategyTarget_State_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'State'
type mockStrategyTarget_State_Call struct {
	*mock.Call
}

// State is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *mockStrategyTarget_Expecter) State(_a0 interface{}) *mockStrategyTarget_State_Call {
	return &mockStrategyTarget_State_Call{Call: _e.mock.On("State", _a0)}
}

func (_c *mockStrategyTarget_State_Call) Run(run func(_a0 context.Context)) *mockStrategyTarget_State_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *mockStrategyTarget_State_Call) Return(_a0 *container.State, _a1 error) *mockStrategyTarget_State_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockStrategyTarget_State_Call) RunAndReturn(run func(context.Context) (*container.State, error)) *mockStrategyTarget_State_Call {
	_c.Call.Return(run)
	return _c
}

// newMockStrategyTarget creates a new instance of mockStrategyTarget. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockStrategyTarget(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockStrategyTarget {
	mock := &mockStrategyTarget{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
