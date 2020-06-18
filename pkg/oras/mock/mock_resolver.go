// Code generated by MockGen. DO NOT EDIT.
// Source: vendor/github.com/containerd/containerd/remotes/resolver.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	content "github.com/containerd/containerd/content"
	remotes "github.com/containerd/containerd/remotes"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	io "io"
	reflect "reflect"
)

// MockResolver is a mock of Resolver interface
type MockResolver struct {
	ctrl     *gomock.Controller
	recorder *MockResolverMockRecorder
}

// MockResolverMockRecorder is the mock recorder for MockResolver
type MockResolverMockRecorder struct {
	mock *MockResolver
}

// NewMockResolver creates a new mock instance
func NewMockResolver(ctrl *gomock.Controller) *MockResolver {
	mock := &MockResolver{ctrl: ctrl}
	mock.recorder = &MockResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResolver) EXPECT() *MockResolverMockRecorder {
	return m.recorder
}

// Resolve mocks base method
func (m *MockResolver) Resolve(ctx context.Context, ref string) (string, v1.Descriptor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resolve", ctx, ref)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(v1.Descriptor)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Resolve indicates an expected call of Resolve
func (mr *MockResolverMockRecorder) Resolve(ctx, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resolve", reflect.TypeOf((*MockResolver)(nil).Resolve), ctx, ref)
}

// Fetcher mocks base method
func (m *MockResolver) Fetcher(ctx context.Context, ref string) (remotes.Fetcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetcher", ctx, ref)
	ret0, _ := ret[0].(remotes.Fetcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetcher indicates an expected call of Fetcher
func (mr *MockResolverMockRecorder) Fetcher(ctx, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetcher", reflect.TypeOf((*MockResolver)(nil).Fetcher), ctx, ref)
}

// Pusher mocks base method
func (m *MockResolver) Pusher(ctx context.Context, ref string) (remotes.Pusher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pusher", ctx, ref)
	ret0, _ := ret[0].(remotes.Pusher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pusher indicates an expected call of Pusher
func (mr *MockResolverMockRecorder) Pusher(ctx, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pusher", reflect.TypeOf((*MockResolver)(nil).Pusher), ctx, ref)
}

// MockFetcher is a mock of Fetcher interface
type MockFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockFetcherMockRecorder
}

// MockFetcherMockRecorder is the mock recorder for MockFetcher
type MockFetcherMockRecorder struct {
	mock *MockFetcher
}

// NewMockFetcher creates a new mock instance
func NewMockFetcher(ctrl *gomock.Controller) *MockFetcher {
	mock := &MockFetcher{ctrl: ctrl}
	mock.recorder = &MockFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFetcher) EXPECT() *MockFetcherMockRecorder {
	return m.recorder
}

// Fetch mocks base method
func (m *MockFetcher) Fetch(ctx context.Context, desc v1.Descriptor) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", ctx, desc)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch
func (mr *MockFetcherMockRecorder) Fetch(ctx, desc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockFetcher)(nil).Fetch), ctx, desc)
}

// MockPusher is a mock of Pusher interface
type MockPusher struct {
	ctrl     *gomock.Controller
	recorder *MockPusherMockRecorder
}

// MockPusherMockRecorder is the mock recorder for MockPusher
type MockPusherMockRecorder struct {
	mock *MockPusher
}

// NewMockPusher creates a new mock instance
func NewMockPusher(ctrl *gomock.Controller) *MockPusher {
	mock := &MockPusher{ctrl: ctrl}
	mock.recorder = &MockPusherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPusher) EXPECT() *MockPusherMockRecorder {
	return m.recorder
}

// Push mocks base method
func (m *MockPusher) Push(ctx context.Context, d v1.Descriptor) (content.Writer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Push", ctx, d)
	ret0, _ := ret[0].(content.Writer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Push indicates an expected call of Push
func (mr *MockPusherMockRecorder) Push(ctx, d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockPusher)(nil).Push), ctx, d)
}