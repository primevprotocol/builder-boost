// Code generated by mockery v2.20.0. DO NOT EDIT.

package rollup

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"
)

// MockRollup is an autogenerated mock type for the Rollup type
type MockRollup struct {
	mock.Mock
}

type MockRollup_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRollup) EXPECT() *MockRollup_Expecter {
	return &MockRollup_Expecter{mock: &_m.Mock}
}

// GetAggregaredStake provides a mock function with given fields: searcher
func (_m *MockRollup) GetAggregaredStake(searcher common.Address) *big.Int {
	ret := _m.Called(searcher)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(common.Address) *big.Int); ok {
		r0 = rf(searcher)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// MockRollup_GetAggregaredStake_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAggregaredStake'
type MockRollup_GetAggregaredStake_Call struct {
	*mock.Call
}

// GetAggregaredStake is a helper method to define mock.On call
//   - searcher common.Address
func (_e *MockRollup_Expecter) GetAggregaredStake(searcher interface{}) *MockRollup_GetAggregaredStake_Call {
	return &MockRollup_GetAggregaredStake_Call{Call: _e.mock.On("GetAggregaredStake", searcher)}
}

func (_c *MockRollup_GetAggregaredStake_Call) Run(run func(searcher common.Address)) *MockRollup_GetAggregaredStake_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address))
	})
	return _c
}

func (_c *MockRollup_GetAggregaredStake_Call) Return(_a0 *big.Int) *MockRollup_GetAggregaredStake_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRollup_GetAggregaredStake_Call) RunAndReturn(run func(common.Address) *big.Int) *MockRollup_GetAggregaredStake_Call {
	_c.Call.Return(run)
	return _c
}

// GetBuilderAddress provides a mock function with given fields:
func (_m *MockRollup) GetBuilderAddress() common.Address {
	ret := _m.Called()

	var r0 common.Address
	if rf, ok := ret.Get(0).(func() common.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	return r0
}

// MockRollup_GetBuilderAddress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBuilderAddress'
type MockRollup_GetBuilderAddress_Call struct {
	*mock.Call
}

// GetBuilderAddress is a helper method to define mock.On call
func (_e *MockRollup_Expecter) GetBuilderAddress() *MockRollup_GetBuilderAddress_Call {
	return &MockRollup_GetBuilderAddress_Call{Call: _e.mock.On("GetBuilderAddress")}
}

func (_c *MockRollup_GetBuilderAddress_Call) Run(run func()) *MockRollup_GetBuilderAddress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRollup_GetBuilderAddress_Call) Return(_a0 common.Address) *MockRollup_GetBuilderAddress_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRollup_GetBuilderAddress_Call) RunAndReturn(run func() common.Address) *MockRollup_GetBuilderAddress_Call {
	_c.Call.Return(run)
	return _c
}

// GetMinimalStake provides a mock function with given fields: builder
func (_m *MockRollup) GetMinimalStake(builder common.Address) *big.Int {
	ret := _m.Called(builder)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(common.Address) *big.Int); ok {
		r0 = rf(builder)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// MockRollup_GetMinimalStake_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMinimalStake'
type MockRollup_GetMinimalStake_Call struct {
	*mock.Call
}

// GetMinimalStake is a helper method to define mock.On call
//   - builder common.Address
func (_e *MockRollup_Expecter) GetMinimalStake(builder interface{}) *MockRollup_GetMinimalStake_Call {
	return &MockRollup_GetMinimalStake_Call{Call: _e.mock.On("GetMinimalStake", builder)}
}

func (_c *MockRollup_GetMinimalStake_Call) Run(run func(builder common.Address)) *MockRollup_GetMinimalStake_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address))
	})
	return _c
}

func (_c *MockRollup_GetMinimalStake_Call) Return(_a0 *big.Int) *MockRollup_GetMinimalStake_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRollup_GetMinimalStake_Call) RunAndReturn(run func(common.Address) *big.Int) *MockRollup_GetMinimalStake_Call {
	_c.Call.Return(run)
	return _c
}

// GetStake provides a mock function with given fields: searcher, commitment
func (_m *MockRollup) GetStake(searcher common.Address, commitment common.Hash) *big.Int {
	ret := _m.Called(searcher, commitment)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(common.Address, common.Hash) *big.Int); ok {
		r0 = rf(searcher, commitment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// MockRollup_GetStake_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStake'
type MockRollup_GetStake_Call struct {
	*mock.Call
}

// GetStake is a helper method to define mock.On call
//   - searcher common.Address
//   - commitment common.Hash
func (_e *MockRollup_Expecter) GetStake(searcher interface{}, commitment interface{}) *MockRollup_GetStake_Call {
	return &MockRollup_GetStake_Call{Call: _e.mock.On("GetStake", searcher, commitment)}
}

func (_c *MockRollup_GetStake_Call) Run(run func(searcher common.Address, commitment common.Hash)) *MockRollup_GetStake_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address), args[1].(common.Hash))
	})
	return _c
}

func (_c *MockRollup_GetStake_Call) Return(_a0 *big.Int) *MockRollup_GetStake_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRollup_GetStake_Call) RunAndReturn(run func(common.Address, common.Hash) *big.Int) *MockRollup_GetStake_Call {
	_c.Call.Return(run)
	return _c
}

// GetStakeRemote provides a mock function with given fields: searcher, commitment
func (_m *MockRollup) GetStakeRemote(searcher common.Address, commitment common.Hash) (*big.Int, error) {
	ret := _m.Called(searcher, commitment)

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Address, common.Hash) (*big.Int, error)); ok {
		return rf(searcher, commitment)
	}
	if rf, ok := ret.Get(0).(func(common.Address, common.Hash) *big.Int); ok {
		r0 = rf(searcher, commitment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Address, common.Hash) error); ok {
		r1 = rf(searcher, commitment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRollup_GetStakeRemote_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStakeRemote'
type MockRollup_GetStakeRemote_Call struct {
	*mock.Call
}

// GetStakeRemote is a helper method to define mock.On call
//   - searcher common.Address
//   - commitment common.Hash
func (_e *MockRollup_Expecter) GetStakeRemote(searcher interface{}, commitment interface{}) *MockRollup_GetStakeRemote_Call {
	return &MockRollup_GetStakeRemote_Call{Call: _e.mock.On("GetStakeRemote", searcher, commitment)}
}

func (_c *MockRollup_GetStakeRemote_Call) Run(run func(searcher common.Address, commitment common.Hash)) *MockRollup_GetStakeRemote_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address), args[1].(common.Hash))
	})
	return _c
}

func (_c *MockRollup_GetStakeRemote_Call) Return(_a0 *big.Int, _a1 error) *MockRollup_GetStakeRemote_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRollup_GetStakeRemote_Call) RunAndReturn(run func(common.Address, common.Hash) (*big.Int, error)) *MockRollup_GetStakeRemote_Call {
	_c.Call.Return(run)
	return _c
}

// IsSyncing provides a mock function with given fields:
func (_m *MockRollup) IsSyncing() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockRollup_IsSyncing_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsSyncing'
type MockRollup_IsSyncing_Call struct {
	*mock.Call
}

// IsSyncing is a helper method to define mock.On call
func (_e *MockRollup_Expecter) IsSyncing() *MockRollup_IsSyncing_Call {
	return &MockRollup_IsSyncing_Call{Call: _e.mock.On("IsSyncing")}
}

func (_c *MockRollup_IsSyncing_Call) Run(run func()) *MockRollup_IsSyncing_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRollup_IsSyncing_Call) Return(_a0 bool) *MockRollup_IsSyncing_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRollup_IsSyncing_Call) RunAndReturn(run func() bool) *MockRollup_IsSyncing_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with given fields: ctx
func (_m *MockRollup) Run(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRollup_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type MockRollup_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockRollup_Expecter) Run(ctx interface{}) *MockRollup_Run_Call {
	return &MockRollup_Run_Call{Call: _e.mock.On("Run", ctx)}
}

func (_c *MockRollup_Run_Call) Run(run func(ctx context.Context)) *MockRollup_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockRollup_Run_Call) Return(_a0 error) *MockRollup_Run_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRollup_Run_Call) RunAndReturn(run func(context.Context) error) *MockRollup_Run_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockRollup interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRollup creates a new instance of MockRollup. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRollup(t mockConstructorTestingTNewMockRollup) *MockRollup {
	mock := &MockRollup{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
