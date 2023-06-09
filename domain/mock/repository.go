// Code generated by MockGen. DO NOT EDIT.
// Source: playground/domain (interfaces: EntryRepository,SessionRepository,TransferRepository,UserRepository,WalletRepository)

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	context "context"
	domain "playground/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockEntryRepository is a mock of EntryRepository interface.
type MockEntryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockEntryRepositoryMockRecorder
}

// MockEntryRepositoryMockRecorder is the mock recorder for MockEntryRepository.
type MockEntryRepositoryMockRecorder struct {
	mock *MockEntryRepository
}

// NewMockEntryRepository creates a new mock instance.
func NewMockEntryRepository(ctrl *gomock.Controller) *MockEntryRepository {
	mock := &MockEntryRepository{ctrl: ctrl}
	mock.recorder = &MockEntryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEntryRepository) EXPECT() *MockEntryRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockEntryRepository) Create(arg0 context.Context, arg1 *domain.Entry) (*domain.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*domain.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockEntryRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockEntryRepository)(nil).Create), arg0, arg1)
}

// Get mocks base method.
func (m *MockEntryRepository) Get(arg0 context.Context, arg1 int64) (*domain.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*domain.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockEntryRepositoryMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockEntryRepository)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockEntryRepository) List(arg0 context.Context, arg1 domain.ListEntriesInputParams) ([]domain.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]domain.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockEntryRepositoryMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockEntryRepository)(nil).List), arg0, arg1)
}

// WithCtx mocks base method.
func (m *MockEntryRepository) WithCtx(arg0 context.Context) domain.EntryRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithCtx", arg0)
	ret0, _ := ret[0].(domain.EntryRepository)
	return ret0
}

// WithCtx indicates an expected call of WithCtx.
func (mr *MockEntryRepositoryMockRecorder) WithCtx(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithCtx", reflect.TypeOf((*MockEntryRepository)(nil).WithCtx), arg0)
}

// MockSessionRepository is a mock of SessionRepository interface.
type MockSessionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSessionRepositoryMockRecorder
}

// MockSessionRepositoryMockRecorder is the mock recorder for MockSessionRepository.
type MockSessionRepositoryMockRecorder struct {
	mock *MockSessionRepository
}

// NewMockSessionRepository creates a new mock instance.
func NewMockSessionRepository(ctrl *gomock.Controller) *MockSessionRepository {
	mock := &MockSessionRepository{ctrl: ctrl}
	mock.recorder = &MockSessionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionRepository) EXPECT() *MockSessionRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSessionRepository) Create(arg0 context.Context, arg1 *domain.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockSessionRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSessionRepository)(nil).Create), arg0, arg1)
}

// Get mocks base method.
func (m *MockSessionRepository) Get(arg0 context.Context, arg1 uuid.UUID) (*domain.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*domain.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSessionRepositoryMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSessionRepository)(nil).Get), arg0, arg1)
}

// MockTransferRepository is a mock of TransferRepository interface.
type MockTransferRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTransferRepositoryMockRecorder
}

// MockTransferRepositoryMockRecorder is the mock recorder for MockTransferRepository.
type MockTransferRepositoryMockRecorder struct {
	mock *MockTransferRepository
}

// NewMockTransferRepository creates a new mock instance.
func NewMockTransferRepository(ctrl *gomock.Controller) *MockTransferRepository {
	mock := &MockTransferRepository{ctrl: ctrl}
	mock.recorder = &MockTransferRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransferRepository) EXPECT() *MockTransferRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTransferRepository) Create(arg0 context.Context, arg1 *domain.Transfer) (*domain.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*domain.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTransferRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTransferRepository)(nil).Create), arg0, arg1)
}

// WithCtx mocks base method.
func (m *MockTransferRepository) WithCtx(arg0 context.Context) domain.TransferRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithCtx", arg0)
	ret0, _ := ret[0].(domain.TransferRepository)
	return ret0
}

// WithCtx indicates an expected call of WithCtx.
func (mr *MockTransferRepositoryMockRecorder) WithCtx(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithCtx", reflect.TypeOf((*MockTransferRepository)(nil).WithCtx), arg0)
}

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

// Create mocks base method.
func (m *MockUserRepository) Create(arg0 context.Context, arg1 *domain.User) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepository)(nil).Create), arg0, arg1)
}

// GetByUsername mocks base method.
func (m *MockUserRepository) GetByUsername(arg0 context.Context, arg1 string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUsername", arg0, arg1)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUsername indicates an expected call of GetByUsername.
func (mr *MockUserRepositoryMockRecorder) GetByUsername(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUsername", reflect.TypeOf((*MockUserRepository)(nil).GetByUsername), arg0, arg1)
}

// MockWalletRepository is a mock of WalletRepository interface.
type MockWalletRepository struct {
	ctrl     *gomock.Controller
	recorder *MockWalletRepositoryMockRecorder
}

// MockWalletRepositoryMockRecorder is the mock recorder for MockWalletRepository.
type MockWalletRepositoryMockRecorder struct {
	mock *MockWalletRepository
}

// NewMockWalletRepository creates a new mock instance.
func NewMockWalletRepository(ctrl *gomock.Controller) *MockWalletRepository {
	mock := &MockWalletRepository{ctrl: ctrl}
	mock.recorder = &MockWalletRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWalletRepository) EXPECT() *MockWalletRepositoryMockRecorder {
	return m.recorder
}

// AddWalletBalance mocks base method.
func (m *MockWalletRepository) AddWalletBalance(arg0 context.Context, arg1 domain.AddWalletBalanceParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWalletBalance", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddWalletBalance indicates an expected call of AddWalletBalance.
func (mr *MockWalletRepositoryMockRecorder) AddWalletBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWalletBalance", reflect.TypeOf((*MockWalletRepository)(nil).AddWalletBalance), arg0, arg1)
}

// Create mocks base method.
func (m *MockWalletRepository) Create(arg0 context.Context, arg1 *domain.Wallet) (*domain.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*domain.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockWalletRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWalletRepository)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockWalletRepository) Delete(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockWalletRepositoryMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWalletRepository)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockWalletRepository) Get(arg0 context.Context, arg1 int64) (*domain.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*domain.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockWalletRepositoryMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockWalletRepository)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockWalletRepository) List(arg0 context.Context, arg1 domain.ListWalletsInputParams) ([]domain.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]domain.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockWalletRepositoryMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockWalletRepository)(nil).List), arg0, arg1)
}

// WithCtx mocks base method.
func (m *MockWalletRepository) WithCtx(arg0 context.Context) domain.WalletRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithCtx", arg0)
	ret0, _ := ret[0].(domain.WalletRepository)
	return ret0
}

// WithCtx indicates an expected call of WithCtx.
func (mr *MockWalletRepositoryMockRecorder) WithCtx(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithCtx", reflect.TypeOf((*MockWalletRepository)(nil).WithCtx), arg0)
}
