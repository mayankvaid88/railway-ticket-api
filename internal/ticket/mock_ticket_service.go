package ticket

import (
        user "railwai-ticket-api/internal/user"
        reflect "reflect"

        gomock "github.com/golang/mock/gomock"
)

// MockTicketService is a mock of TicketService interface.
type MockTicketService struct {
        ctrl     *gomock.Controller
        recorder *MockTicketServiceMockRecorder
}

// MockTicketServiceMockRecorder is the mock recorder for MockTicketService.
type MockTicketServiceMockRecorder struct {
        mock *MockTicketService
}

// NewMockTicketService creates a new mock instance.
func NewMockTicketService(ctrl *gomock.Controller) *MockTicketService {
        mock := &MockTicketService{ctrl: ctrl}
        mock.recorder = &MockTicketServiceMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTicketService) EXPECT() *MockTicketServiceMockRecorder {
        return m.recorder
}

// BookTicket mocks base method.
func (m *MockTicketService) BookTicket(user user.User, from, to string) (Ticket, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "BookTicket", user, from, to)
        ret0, _ := ret[0].(Ticket)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// BookTicket indicates an expected call of BookTicket.
func (mr *MockTicketServiceMockRecorder) BookTicket(user, from, to interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BookTicket", reflect.TypeOf((*MockTicketService)(nil).BookTicket), user, from, to)
}

// CancelBooking mocks base method.
func (m *MockTicketService) CancelBooking(emailAddress string) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CancelBooking", emailAddress)
        ret0, _ := ret[0].(error)
        return ret0
}

// CancelBooking indicates an expected call of CancelBooking.
func (mr *MockTicketServiceMockRecorder) CancelBooking(emailAddress interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelBooking", reflect.TypeOf((*MockTicketService)(nil).CancelBooking), emailAddress)
}

// GetSeatsPerSection mocks base method.
func (m *MockTicketService) GetSeatsPerSection(section string) ([]Seat, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetSeatsPerSection", section)
        ret0, _ := ret[0].([]Seat)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetSeatsPerSection indicates an expected call of GetSeatsPerSection.
func (mr *MockTicketServiceMockRecorder) GetSeatsPerSection(section interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSeatsPerSection", reflect.TypeOf((*MockTicketService)(nil).GetSeatsPerSection), section)
}

// GetTicket mocks base method.
func (m *MockTicketService) GetTicket(emailAddress string) (Ticket, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetTicket", emailAddress)
        ret0, _ := ret[0].(Ticket)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetTicket indicates an expected call of GetTicket.
func (mr *MockTicketServiceMockRecorder) GetTicket(emailAddress interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTicket", reflect.TypeOf((*MockTicketService)(nil).GetTicket), emailAddress)
}