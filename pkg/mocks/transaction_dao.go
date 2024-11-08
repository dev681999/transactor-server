// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"
	ent "transactor-server/pkg/db/ent"

	mock "github.com/stretchr/testify/mock"

	transaction "transactor-server/pkg/transaction"
)

// MockTransactionDAO is an autogenerated mock type for the DAO type
type MockTransactionDAO struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, req
func (_m *MockTransactionDAO) Create(ctx context.Context, req *transaction.CreateRequest) (*ent.Transaction, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *ent.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *transaction.CreateRequest) (*ent.Transaction, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *transaction.CreateRequest) *ent.Transaction); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *transaction.CreateRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockTransactionDAO creates a new instance of MockTransactionDAO. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTransactionDAO(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTransactionDAO {
	mock := &MockTransactionDAO{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
