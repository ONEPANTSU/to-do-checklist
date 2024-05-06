// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"
	domain "to-do-checklist/internal/domain"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthorization is a mock of Authorization interface.
type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization.
type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance.
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAuthorization) CreateUser(user domain.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthorizationMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorization)(nil).CreateUser), user)
}

// GenerateToken mocks base method.
func (m *MockAuthorization) GenerateToken(authInfo domain.SignIn) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", authInfo)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthorizationMockRecorder) GenerateToken(authInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthorization)(nil).GenerateToken), authInfo)
}

// ValidateToken mocks base method.
func (m *MockAuthorization) ValidateToken(token string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateToken", token)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateToken indicates an expected call of ValidateToken.
func (mr *MockAuthorizationMockRecorder) ValidateToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateToken", reflect.TypeOf((*MockAuthorization)(nil).ValidateToken), token)
}

// generatePasswordHash mocks base method.
func (m *MockAuthorization) generatePasswordHash(password string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "generatePasswordHash", password)
	ret0, _ := ret[0].(string)
	return ret0
}

// generatePasswordHash indicates an expected call of generatePasswordHash.
func (mr *MockAuthorizationMockRecorder) generatePasswordHash(password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "generatePasswordHash", reflect.TypeOf((*MockAuthorization)(nil).generatePasswordHash), password)
}

// MockTodoList is a mock of TodoList interface.
type MockTodoList struct {
	ctrl     *gomock.Controller
	recorder *MockTodoListMockRecorder
}

// MockTodoListMockRecorder is the mock recorder for MockTodoList.
type MockTodoListMockRecorder struct {
	mock *MockTodoList
}

// NewMockTodoList creates a new mock instance.
func NewMockTodoList(ctrl *gomock.Controller) *MockTodoList {
	mock := &MockTodoList{ctrl: ctrl}
	mock.recorder = &MockTodoListMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoList) EXPECT() *MockTodoListMockRecorder {
	return m.recorder
}

// CreateList mocks base method.
func (m *MockTodoList) CreateList(list *domain.TodoList, userID int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", list, userID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockTodoListMockRecorder) CreateList(list, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockTodoList)(nil).CreateList), list, userID)
}

// DeleteList mocks base method.
func (m *MockTodoList) DeleteList(listID, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteList", listID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteList indicates an expected call of DeleteList.
func (mr *MockTodoListMockRecorder) DeleteList(listID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteList", reflect.TypeOf((*MockTodoList)(nil).DeleteList), listID, userID)
}

// GetAllLists mocks base method.
func (m *MockTodoList) GetAllLists() *[]domain.TodoList {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllLists")
	ret0, _ := ret[0].(*[]domain.TodoList)
	return ret0
}

// GetAllLists indicates an expected call of GetAllLists.
func (mr *MockTodoListMockRecorder) GetAllLists() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllLists", reflect.TypeOf((*MockTodoList)(nil).GetAllLists))
}

// GetListById mocks base method.
func (m *MockTodoList) GetListById(listID, userID int) (*domain.TodoList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListById", listID, userID)
	ret0, _ := ret[0].(*domain.TodoList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListById indicates an expected call of GetListById.
func (mr *MockTodoListMockRecorder) GetListById(listID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListById", reflect.TypeOf((*MockTodoList)(nil).GetListById), listID, userID)
}

// GetUsersLists mocks base method.
func (m *MockTodoList) GetUsersLists(userID int) (*[]domain.TodoList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersLists", userID)
	ret0, _ := ret[0].(*[]domain.TodoList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersLists indicates an expected call of GetUsersLists.
func (mr *MockTodoListMockRecorder) GetUsersLists(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersLists", reflect.TypeOf((*MockTodoList)(nil).GetUsersLists), userID)
}

// UpdateList mocks base method.
func (m *MockTodoList) UpdateList(list *domain.UpdateTodoList, listID, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateList", list, listID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateList indicates an expected call of UpdateList.
func (mr *MockTodoListMockRecorder) UpdateList(list, listID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateList", reflect.TypeOf((*MockTodoList)(nil).UpdateList), list, listID, userID)
}

// MockTodoItem is a mock of TodoItem interface.
type MockTodoItem struct {
	ctrl     *gomock.Controller
	recorder *MockTodoItemMockRecorder
}

// MockTodoItemMockRecorder is the mock recorder for MockTodoItem.
type MockTodoItemMockRecorder struct {
	mock *MockTodoItem
}

// NewMockTodoItem creates a new mock instance.
func NewMockTodoItem(ctrl *gomock.Controller) *MockTodoItem {
	mock := &MockTodoItem{ctrl: ctrl}
	mock.recorder = &MockTodoItemMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoItem) EXPECT() *MockTodoItemMockRecorder {
	return m.recorder
}

// CreateItem mocks base method.
func (m *MockTodoItem) CreateItem(item *domain.TodoItem, userID int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateItem", item, userID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateItem indicates an expected call of CreateItem.
func (mr *MockTodoItemMockRecorder) CreateItem(item, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateItem", reflect.TypeOf((*MockTodoItem)(nil).CreateItem), item, userID)
}

// DeleteItem mocks base method.
func (m *MockTodoItem) DeleteItem(itemID, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteItem", itemID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteItem indicates an expected call of DeleteItem.
func (mr *MockTodoItemMockRecorder) DeleteItem(itemID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteItem", reflect.TypeOf((*MockTodoItem)(nil).DeleteItem), itemID, userID)
}

// GetItemById mocks base method.
func (m *MockTodoItem) GetItemById(itemID, userID int) (*domain.TodoItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemById", itemID, userID)
	ret0, _ := ret[0].(*domain.TodoItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemById indicates an expected call of GetItemById.
func (mr *MockTodoItemMockRecorder) GetItemById(itemID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemById", reflect.TypeOf((*MockTodoItem)(nil).GetItemById), itemID, userID)
}

// GetItems mocks base method.
func (m *MockTodoItem) GetItems(listID, userID int) (*[]domain.TodoItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItems", listID, userID)
	ret0, _ := ret[0].(*[]domain.TodoItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItems indicates an expected call of GetItems.
func (mr *MockTodoItemMockRecorder) GetItems(listID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItems", reflect.TypeOf((*MockTodoItem)(nil).GetItems), listID, userID)
}

// UpdateItem mocks base method.
func (m *MockTodoItem) UpdateItem(item *domain.UpdateTodoItem, itemID, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateItem", item, itemID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateItem indicates an expected call of UpdateItem.
func (mr *MockTodoItemMockRecorder) UpdateItem(item, itemID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateItem", reflect.TypeOf((*MockTodoItem)(nil).UpdateItem), item, itemID, userID)
}
