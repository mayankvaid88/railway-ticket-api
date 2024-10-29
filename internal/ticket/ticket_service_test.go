package ticket

import (
	"errors"
	"railwai-ticket-api/internal/user"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type TicketServiceTestSuite struct {
	suite.Suite
	ctrl          *gomock.Controller
	mockManager   *MockManager
	ticketService TicketService
}

func TestTicketServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TicketServiceTestSuite))
}

func (m *TicketServiceTestSuite) SetupSuite() {
	m.ctrl = gomock.NewController(m.T())
	m.mockManager = NewMockManager(m.ctrl)
	m.ticketService = NewTicketService(m.mockManager)
}

func (m *TicketServiceTestSuite) TearDownTest() {
	m.ctrl.Finish()
}

func (m *TicketServiceTestSuite) TestBookTicket_ShouldBeSuccessful() {
	u := user.User{Name: "TEST", EmalAddress: "abcd"}
	m.mockManager.EXPECT().BookTicket(u, "A", "B").Return(Ticket{
		Id:    "abcd",
		From:  "A",
		To:    "B",
		Price: 20,
		Seat: Seat{
			Number:  1,
			Section: "A",
		},
	}, nil)
	ticket, err := m.ticketService.BookTicket(u, "A", "B")
	m.Equal(Ticket{
		Id:    "abcd",
		From:  "A",
		To:    "B",
		Price: 20,
		Seat: Seat{
			Number:  1,
			Section: "A",
		},
	}, ticket)
	m.Nil(err)
}

func (m *TicketServiceTestSuite) TestBookTicket_ShouldReturnErrorWhenManagerThrowsError() {
	u := user.User{Name: "TEST", EmalAddress: "abcd"}
	m.mockManager.EXPECT().BookTicket(u, "A", "B").Return(Ticket{}, errors.New("unable to book"))
	ticket, err := m.ticketService.BookTicket(u, "A", "B")
	m.Empty(ticket)
	m.EqualError(err, "unable to book")
}

func (m *TicketServiceTestSuite) TestGetTicket_ShouldBeSuccessful() {
	m.mockManager.EXPECT().GetTicket("abcd").Return(Ticket{
		Id:    "abcd",
		From:  "A",
		To:    "B",
		Price: 20,
		Seat: Seat{
			Number:  1,
			Section: "A",
		},
	}, nil)
	ticket, err := m.ticketService.GetTicket("abcd")
	m.Equal(Ticket{
		Id:    "abcd",
		From:  "A",
		To:    "B",
		Price: 20,
		Seat: Seat{
			Number:  1,
			Section: "A",
		},
	}, ticket)
	m.Nil(err)
}

func (m *TicketServiceTestSuite) TestGetTicket_ShouldReturnErrorWhenManagerThrowsError() {
	m.mockManager.EXPECT().GetTicket("abcd").Return(Ticket{}, errors.New("unable to fetch"))
	ticket, err := m.ticketService.GetTicket("abcd")
	m.Empty(ticket)
	m.EqualError(err, "unable to fetch")
}

func (m *TicketServiceTestSuite) TestGetSeatsForGivenSection_ShouldBeSuccessful() {
	m.mockManager.EXPECT().GetSeatsBySection("A").Return([]Seat{
		{
			Number:  1,
			Section: "A",
		},
	}, nil)
	seats, err := m.ticketService.GetSeatsPerSection("A")
	m.Len(seats, 1)
	m.Equal(seats[0], Seat{
		Number:  1,
		Section: "A",
	})
	m.Nil(err)
}

func (m *TicketServiceTestSuite) TestGetSeats_ShouldReturnErrorWhenManagerThrowsError() {
	m.mockManager.EXPECT().GetSeatsBySection("A").Return(nil, errors.New("unable to fetch"))
	seats, err := m.ticketService.GetSeatsPerSection("A")
	m.Nil(seats)
	m.EqualError(err, "unable to fetch")
}

func (m *TicketServiceTestSuite) TestCancelTicket_ShouldBeSuccessful() {
	m.mockManager.EXPECT().CancelTicket("ABC").Return(nil)
	err := m.ticketService.CancelBooking("ABC")
	m.Nil(err)
}

func (m *TicketServiceTestSuite) TestCancelBooking_ShouldReturnErrorWhenManagerThrowsError() {
	m.mockManager.EXPECT().CancelTicket("ABC").Return(errors.New("unable to cancel"))
	err := m.ticketService.CancelBooking("ABC")
	m.EqualError(err, "unable to cancel")
}
