// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/service.go

// Package servicemock is a generated GoMock package.
package servicemock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongostore "github.com/junkd0g/go-api-example/internal/persistence/mongo"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// DeleteMovie mocks base method.
func (m *MockStore) DeleteMovie(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMovie", ctx, id)
	ret0, _ := ret[0].(*mongo.DeleteResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMovie indicates an expected call of DeleteMovie.
func (mr *MockStoreMockRecorder) DeleteMovie(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMovie", reflect.TypeOf((*MockStore)(nil).DeleteMovie), ctx, id)
}

// FindMovie mocks base method.
func (m *MockStore) FindMovie(ctx context.Context, title string) (*mongostore.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindMovie", ctx, title)
	ret0, _ := ret[0].(*mongostore.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindMovie indicates an expected call of FindMovie.
func (mr *MockStoreMockRecorder) FindMovie(ctx, title interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMovie", reflect.TypeOf((*MockStore)(nil).FindMovie), ctx, title)
}

// InsertMovie mocks base method.
func (m *MockStore) InsertMovie(ctx context.Context, movie mongostore.Movie) (*mongo.InsertOneResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertMovie", ctx, movie)
	ret0, _ := ret[0].(*mongo.InsertOneResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertMovie indicates an expected call of InsertMovie.
func (mr *MockStoreMockRecorder) InsertMovie(ctx, movie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertMovie", reflect.TypeOf((*MockStore)(nil).InsertMovie), ctx, movie)
}
