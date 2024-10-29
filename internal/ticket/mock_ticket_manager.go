package ticket

import (
        user "railwai-ticket-api/internal/user"
        reflect "reflect"

        gomock "github.com/golang/mock/gomock"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
        ctrl     *gomock.Controller
        recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
        mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
        mock := &MockManager{ctrl: ctrl}
        mock.recorder = &MockManagerMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
        return m.recorder
}

// BookTicket mocks base method.
func (m *MockManager) BookTicket(u user.User, from, to string) (Ticket, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "BookTicket", u, from, to)
        ret0, _ := ret[0].(Ticket)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// BookTicket indicates an expected call of BookTicket.
func (mr *MockManagerMockRecorder) BookTicket(u, from, to interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BookTicket", reflect.TypeOf((*MockManager)(nil).BookTicket), u, from, to)
}

// CancelTicket mocks base method.
func (m *MockManager) CancelTicket(emailAddress string) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CancelTicket", emailAddress)
        ret0, _ := ret[0].(error)
        return ret0
}

// CancelTicket indicates an expected call of CancelTicket.
func (mr *MockManagerMockRecorder) CancelTicket(emailAddress interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelTicket", reflect.TypeOf((*MockManager)(nil).CancelTicket), emailAddress)
}

// GetSeatsBySection mocks base method.
func (m *MockManager) GetSeatsBySection(section string) ([]Seat, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetSeatsBySection", section)
        ret0, _ := ret[0].([]Seat)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetSeatsBySection indicates an expected call of GetSeatsBySection.
func (mr *MockManagerMockRecorder) GetSeatsBySection(section interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSeatsBySection", reflect.TypeOf((*MockManager)(nil).GetSeatsBySection), section)
}

// GetTicket mocks base method.
func (m *MockManager) GetTicket(emailAddress string) (Ticket, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetTicket", emailAddress)
        ret0, _ := ret[0].(Ticket)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetTicket indicates an expected call of GetTicket.
func (mr *MockManagerMockRecorder) GetTicket(emailAddress interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTicket", reflect.TypeOf((*MockManager)(nil).GetTicket), emailAddress)
}