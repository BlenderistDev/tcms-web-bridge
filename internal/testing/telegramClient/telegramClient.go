// Code generated by MockGen. DO NOT EDIT.
// Source: internal/telegramClient/telegramClient.go

// Package mock_telegramClient is a generated GoMock package.
package mock_telegramClient

import (
	context "context"
	reflect "reflect"
	telegram "tcms-web-bridge/pkg/telegram"

	gomock "github.com/golang/mock/gomock"
)

// MockTelegramClient is a mock of TelegramClient interface.
type MockTelegramClient struct {
	ctrl     *gomock.Controller
	recorder *MockTelegramClientMockRecorder
}

// MockTelegramClientMockRecorder is the mock recorder for MockTelegramClient.
type MockTelegramClientMockRecorder struct {
	mock *MockTelegramClient
}

// NewMockTelegramClient creates a new mock instance.
func NewMockTelegramClient(ctrl *gomock.Controller) *MockTelegramClient {
	mock := &MockTelegramClient{ctrl: ctrl}
	mock.recorder = &MockTelegramClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTelegramClient) EXPECT() *MockTelegramClientMockRecorder {
	return m.recorder
}

// AuthSignIn mocks base method.
func (m *MockTelegramClient) AuthSignIn(ctx context.Context, code string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthSignIn", ctx, code)
	ret0, _ := ret[0].(error)
	return ret0
}

// AuthSignIn indicates an expected call of AuthSignIn.
func (mr *MockTelegramClientMockRecorder) AuthSignIn(ctx, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthSignIn", reflect.TypeOf((*MockTelegramClient)(nil).AuthSignIn), ctx, code)
}

// Authorization mocks base method.
func (m *MockTelegramClient) Authorization(ctx context.Context, phone string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorization", ctx, phone)
	ret0, _ := ret[0].(error)
	return ret0
}

// Authorization indicates an expected call of Authorization.
func (mr *MockTelegramClientMockRecorder) Authorization(ctx, phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorization", reflect.TypeOf((*MockTelegramClient)(nil).Authorization), ctx, phone)
}

// Dialogs mocks base method.
func (m *MockTelegramClient) Dialogs(ctx context.Context) (*telegram.DialogsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dialogs", ctx)
	ret0, _ := ret[0].(*telegram.DialogsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Dialogs indicates an expected call of Dialogs.
func (mr *MockTelegramClientMockRecorder) Dialogs(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dialogs", reflect.TypeOf((*MockTelegramClient)(nil).Dialogs), ctx)
}

// GetCurrentUser mocks base method.
func (m *MockTelegramClient) GetCurrentUser(ctx context.Context) (*telegram.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentUser", ctx)
	ret0, _ := ret[0].(*telegram.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentUser indicates an expected call of GetCurrentUser.
func (mr *MockTelegramClientMockRecorder) GetCurrentUser(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentUser", reflect.TypeOf((*MockTelegramClient)(nil).GetCurrentUser), ctx)
}

// SendMessage mocks base method.
func (m *MockTelegramClient) SendMessage(ctx context.Context, peer, message string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", ctx, peer, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockTelegramClientMockRecorder) SendMessage(ctx, peer, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockTelegramClient)(nil).SendMessage), ctx, peer, message)
}
