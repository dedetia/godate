// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/core/port/repository/user.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	domain "github.com/dedetia/godate/internal/core/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CountUserRecommendation mocks base method.
func (m *MockUserRepository) CountUserRecommendation(ctx context.Context, request *domain.UserRecommendation) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountUserRecommendation", ctx, request)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountUserRecommendation indicates an expected call of CountUserRecommendation.
func (mr *MockUserRepositoryMockRecorder) CountUserRecommendation(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountUserRecommendation", reflect.TypeOf((*MockUserRepository)(nil).CountUserRecommendation), ctx, request)
}

// Create mocks base method.
func (m *MockUserRepository) Create(ctx context.Context, user *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserRepositoryMockRecorder) Create(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepository)(nil).Create), ctx, user)
}

// GetByID mocks base method.
func (m *MockUserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockUserRepositoryMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUserRepository)(nil).GetByID), ctx, id)
}

// GetByPhoneNumber mocks base method.
func (m *MockUserRepository) GetByPhoneNumber(ctx context.Context, phoneNumber string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByPhoneNumber", ctx, phoneNumber)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByPhoneNumber indicates an expected call of GetByPhoneNumber.
func (mr *MockUserRepositoryMockRecorder) GetByPhoneNumber(ctx, phoneNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByPhoneNumber", reflect.TypeOf((*MockUserRepository)(nil).GetByPhoneNumber), ctx, phoneNumber)
}

// GetByPhoneNumberOrEmail mocks base method.
func (m *MockUserRepository) GetByPhoneNumberOrEmail(ctx context.Context, phoneNumber, email string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByPhoneNumberOrEmail", ctx, phoneNumber, email)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByPhoneNumberOrEmail indicates an expected call of GetByPhoneNumberOrEmail.
func (mr *MockUserRepositoryMockRecorder) GetByPhoneNumberOrEmail(ctx, phoneNumber, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByPhoneNumberOrEmail", reflect.TypeOf((*MockUserRepository)(nil).GetByPhoneNumberOrEmail), ctx, phoneNumber, email)
}

// GetUserRecommendation mocks base method.
func (m *MockUserRepository) GetUserRecommendation(ctx context.Context, request *domain.UserRecommendation) ([]*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserRecommendation", ctx, request)
	ret0, _ := ret[0].([]*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserRecommendation indicates an expected call of GetUserRecommendation.
func (mr *MockUserRepositoryMockRecorder) GetUserRecommendation(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRecommendation", reflect.TypeOf((*MockUserRepository)(nil).GetUserRecommendation), ctx, request)
}

// Update mocks base method.
func (m *MockUserRepository) Update(ctx context.Context, user *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserRepositoryMockRecorder) Update(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRepository)(nil).Update), ctx, user)
}
