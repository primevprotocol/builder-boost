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

// GetBlockNumber provides a mock function with given fields:
func (_m *MockRollup) GetBlockNumber() (*big.Int, error) {
	ret := _m.Called()

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func() (*big.Int, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *big.Int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRollup_GetBlockNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBlockNumber'
type MockRollup_GetBlockNumber_Call struct {
	*mock.Call
}

// GetBlockNumber is a helper method to define mock.On call
func (_e *MockRollup_Expecter) GetBlockNumber() *MockRollup_GetBlockNumber_Call {
	return &MockRollup_GetBlockNumber_Call{Call: _e.mock.On("GetBlockNumber")}
}

func (_c *MockRollup_GetBlockNumber_Call) Run(run func()) *MockRollup_GetBlockNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRollup_GetBlockNumber_Call) Return(_a0 *big.Int, _a1 error) *MockRollup_GetBlockNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRollup_GetBlockNumber_Call) RunAndReturn(run func() (*big.Int, error)) *MockRollup_GetBlockNumber_Call {
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

// GetCommitment provides a mock function with given fields: searcher
func (_m *MockRollup) GetCommitment(searcher common.Address) common.Hash {
	ret := _m.Called(searcher)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(common.Address) common.Hash); ok {
		r0 = rf(searcher)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	return r0
}

// MockRollup_GetCommitment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCommitment'
type MockRollup_GetCommitment_Call struct {
	*mock.Call
}

// GetCommitment is a helper method to define mock.On call
//   - searcher common.Address
func (_e *MockRollup_Expecter) GetCommitment(searcher interface{}) *MockRollup_GetCommitment_Call {
	return &MockRollup_GetCommitment_Call{Call: _e.mock.On("GetCommitment", searcher)}
}

func (_c *MockRollup_GetCommitment_Call) Run(run func(searcher common.Address)) *MockRollup_GetCommitment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address))
	})
	return _c
}

func (_c *MockRollup_GetCommitment_Call) Return(_a0 common.Hash) *MockRollup_GetCommitment_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRollup_GetCommitment_Call) RunAndReturn(run func(common.Address) common.Hash) *MockRollup_GetCommitment_Call {
	_c.Call.Return(run)
	return _c
}

// GetMinimalStake provides a mock function with given fields: builder
func (_m *MockRollup) GetMinimalStake(builder common.Address) (*big.Int, error) {
	ret := _m.Called(builder)

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Address) (*big.Int, error)); ok {
		return rf(builder)
	}
	if rf, ok := ret.Get(0).(func(common.Address) *big.Int); ok {
		r0 = rf(builder)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(builder)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
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

func (_c *MockRollup_GetMinimalStake_Call) Return(_a0 *big.Int, _a1 error) *MockRollup_GetMinimalStake_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRollup_GetMinimalStake_Call) RunAndReturn(run func(common.Address) (*big.Int, error)) *MockRollup_GetMinimalStake_Call {
	_c.Call.Return(run)
	return _c
}

// GetSubscriptionEnd provides a mock function with given fields: commitment
func (_m *MockRollup) GetSubscriptionEnd(commitment common.Hash) (*big.Int, error) {
	ret := _m.Called(commitment)

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Hash) (*big.Int, error)); ok {
		return rf(commitment)
	}
	if rf, ok := ret.Get(0).(func(common.Hash) *big.Int); ok {
		r0 = rf(commitment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(commitment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRollup_GetSubscriptionEnd_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSubscriptionEnd'
type MockRollup_GetSubscriptionEnd_Call struct {
	*mock.Call
}

// GetSubscriptionEnd is a helper method to define mock.On call
//   - commitment common.Hash
func (_e *MockRollup_Expecter) GetSubscriptionEnd(commitment interface{}) *MockRollup_GetSubscriptionEnd_Call {
	return &MockRollup_GetSubscriptionEnd_Call{Call: _e.mock.On("GetSubscriptionEnd", commitment)}
}

func (_c *MockRollup_GetSubscriptionEnd_Call) Run(run func(commitment common.Hash)) *MockRollup_GetSubscriptionEnd_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Hash))
	})
	return _c
}

func (_c *MockRollup_GetSubscriptionEnd_Call) Return(_a0 *big.Int, _a1 error) *MockRollup_GetSubscriptionEnd_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRollup_GetSubscriptionEnd_Call) RunAndReturn(run func(common.Hash) (*big.Int, error)) *MockRollup_GetSubscriptionEnd_Call {
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
