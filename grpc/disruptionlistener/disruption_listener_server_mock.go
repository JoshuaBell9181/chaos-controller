// Code generated by mockery. DO NOT EDIT.

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023 Datadog, Inc.
package disruptionlistener

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// DisruptionListenerServerMock is an autogenerated mock type for the DisruptionListenerServer type
type DisruptionListenerServerMock struct {
	mock.Mock
}

type DisruptionListenerServerMock_Expecter struct {
	mock *mock.Mock
}

func (_m *DisruptionListenerServerMock) EXPECT() *DisruptionListenerServerMock_Expecter {
	return &DisruptionListenerServerMock_Expecter{mock: &_m.Mock}
}

// Disrupt provides a mock function with given fields: _a0, _a1
func (_m *DisruptionListenerServerMock) Disrupt(_a0 context.Context, _a1 *DisruptionSpec) (*emptypb.Empty, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *emptypb.Empty
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *DisruptionSpec) (*emptypb.Empty, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *DisruptionSpec) *emptypb.Empty); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *DisruptionSpec) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisruptionListenerServerMock_Disrupt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Disrupt'
type DisruptionListenerServerMock_Disrupt_Call struct {
	*mock.Call
}

// Disrupt is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *DisruptionSpec
func (_e *DisruptionListenerServerMock_Expecter) Disrupt(_a0 interface{}, _a1 interface{}) *DisruptionListenerServerMock_Disrupt_Call {
	return &DisruptionListenerServerMock_Disrupt_Call{Call: _e.mock.On("Disrupt", _a0, _a1)}
}

func (_c *DisruptionListenerServerMock_Disrupt_Call) Run(run func(_a0 context.Context, _a1 *DisruptionSpec)) *DisruptionListenerServerMock_Disrupt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*DisruptionSpec))
	})
	return _c
}

func (_c *DisruptionListenerServerMock_Disrupt_Call) Return(_a0 *emptypb.Empty, _a1 error) *DisruptionListenerServerMock_Disrupt_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DisruptionListenerServerMock_Disrupt_Call) RunAndReturn(run func(context.Context, *DisruptionSpec) (*emptypb.Empty, error)) *DisruptionListenerServerMock_Disrupt_Call {
	_c.Call.Return(run)
	return _c
}

// ResetDisruptions provides a mock function with given fields: _a0, _a1
func (_m *DisruptionListenerServerMock) ResetDisruptions(_a0 context.Context, _a1 *emptypb.Empty) (*emptypb.Empty, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *emptypb.Empty
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emptypb.Empty) (*emptypb.Empty, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emptypb.Empty) *emptypb.Empty); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emptypb.Empty) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisruptionListenerServerMock_ResetDisruptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ResetDisruptions'
type DisruptionListenerServerMock_ResetDisruptions_Call struct {
	*mock.Call
}

// ResetDisruptions is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *emptypb.Empty
func (_e *DisruptionListenerServerMock_Expecter) ResetDisruptions(_a0 interface{}, _a1 interface{}) *DisruptionListenerServerMock_ResetDisruptions_Call {
	return &DisruptionListenerServerMock_ResetDisruptions_Call{Call: _e.mock.On("ResetDisruptions", _a0, _a1)}
}

func (_c *DisruptionListenerServerMock_ResetDisruptions_Call) Run(run func(_a0 context.Context, _a1 *emptypb.Empty)) *DisruptionListenerServerMock_ResetDisruptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*emptypb.Empty))
	})
	return _c
}

func (_c *DisruptionListenerServerMock_ResetDisruptions_Call) Return(_a0 *emptypb.Empty, _a1 error) *DisruptionListenerServerMock_ResetDisruptions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DisruptionListenerServerMock_ResetDisruptions_Call) RunAndReturn(run func(context.Context, *emptypb.Empty) (*emptypb.Empty, error)) *DisruptionListenerServerMock_ResetDisruptions_Call {
	_c.Call.Return(run)
	return _c
}

// mustEmbedUnimplementedDisruptionListenerServer provides a mock function with given fields:
func (_m *DisruptionListenerServerMock) mustEmbedUnimplementedDisruptionListenerServer() {
	_m.Called()
}

// DisruptionListenerServerMock_mustEmbedUnimplementedDisruptionListenerServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedDisruptionListenerServer'
type DisruptionListenerServerMock_mustEmbedUnimplementedDisruptionListenerServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedDisruptionListenerServer is a helper method to define mock.On call
func (_e *DisruptionListenerServerMock_Expecter) mustEmbedUnimplementedDisruptionListenerServer() *DisruptionListenerServerMock_mustEmbedUnimplementedDisruptionListenerServer_Call {
	return &DisruptionListenerServerMock_mustEmbedUnimplementedDisruptionListenerServer_Call{Call: _e.mock.On("mustEmbedUnimplementedDisruptionListenerServer")}
}

func (_c *DisruptionListenerServerMock_mustEmbedUnimplementedDisruptionListenerServer_Call) Run(run func()) *DisruptionListenerServerMock_mustEmbedUnimplementedDisruptionListenerServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DisruptionListenerServerMock_mustEmbedUnimplementedDisruptionListenerServer_Call) Return() *DisruptionListenerServerMock_mustEmbedUnimplementedDisruptionListenerServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *DisruptionListenerServerMock_mustEmbedUnimplementedDisruptionListenerServer_Call) RunAndReturn(run func()) *DisruptionListenerServerMock_mustEmbedUnimplementedDisruptionListenerServer_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewDisruptionListenerServerMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewDisruptionListenerServerMock creates a new instance of DisruptionListenerServerMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDisruptionListenerServerMock(t mockConstructorTestingTNewDisruptionListenerServerMock) *DisruptionListenerServerMock {
	mock := &DisruptionListenerServerMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}