// Code generated by MockGen. DO NOT EDIT.
// Source: x/sync/client.go
//
// Generated by this command:
//
//	mockgen -source=x/sync/client.go -destination=x/sync/mock_client.go -package=sync -exclude_interfaces= -mock_names=MockClient=MockClient
//

// Package sync is a generated GoMock package.
package sync

import (
	context "context"
	reflect "reflect"

	sync "github.com/ava-labs/avalanchego/proto/pb/sync"
	merkledb "github.com/ava-labs/avalanchego/x/merkledb"
	gomock "go.uber.org/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// GetChangeProof mocks base method.
func (m *MockClient) GetChangeProof(ctx context.Context, request *sync.SyncGetChangeProofRequest, verificationDB DB) (*merkledb.ChangeOrRangeProof, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChangeProof", ctx, request, verificationDB)
	ret0, _ := ret[0].(*merkledb.ChangeOrRangeProof)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChangeProof indicates an expected call of GetChangeProof.
func (mr *MockClientMockRecorder) GetChangeProof(ctx, request, verificationDB any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChangeProof", reflect.TypeOf((*MockClient)(nil).GetChangeProof), ctx, request, verificationDB)
}

// GetRangeProof mocks base method.
func (m *MockClient) GetRangeProof(ctx context.Context, request *sync.SyncGetRangeProofRequest) (*merkledb.RangeProof, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRangeProof", ctx, request)
	ret0, _ := ret[0].(*merkledb.RangeProof)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRangeProof indicates an expected call of GetRangeProof.
func (mr *MockClientMockRecorder) GetRangeProof(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRangeProof", reflect.TypeOf((*MockClient)(nil).GetRangeProof), ctx, request)
}
