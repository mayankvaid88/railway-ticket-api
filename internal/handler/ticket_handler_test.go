package handler

import (
	"context"
	"errors"
	"railwai-ticket-api/internal/proto"
	"railwai-ticket-api/internal/ticket"
	"railwai-ticket-api/internal/user"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type TicketHandlerTestSuite struct {
	suite.Suite
	ctrl              *gomock.Controller
	mockTicketService *ticket.MockTicketService
	ticketHandler     TicketHandler
	ctx               context.Context
}

func TestTicketHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(TicketHandlerTestSuite))
}

func (m *TicketHandlerTestSuite) SetupSuite() {
	m.ctrl = gomock.NewController(m.T())
	m.mockTicketService = ticket.NewMockTicketService(m.ctrl)
	m.ticketHandler = NewTicketHandler(m.mockTicketService)
	m.ctx = context.Background()
}

func (m *TicketHandlerTestSuite) TearDownTest() {
	m.ctrl.Finish()
}

func (m *TicketHandlerTestSuite) TestBookTicket_ShouldBeSuccessful() {
	u := user.User{Name: "TEST", EmalAddress: "abcd"}
	m.mockTicketService.EXPECT().BookTicket(u, "X", "Y").Return(ticket.Ticket{
		Id:    "23312",
		From:  "X",
		To:    "Y",
		Price: 20,
		Seat: ticket.Seat{
			Number:  1,
			Section: "A",
		},
	}, nil)
	ticket, err := m.ticketHandler.BookTicket(m.ctx, &proto.BookTicketRequest{
		Name:         "TEST",
		From:         "X",
		To:           "Y",
		EmailAddress: "abcd",
	})
	m.Equal(&proto.BookTicketResponse{
		Id:    "23312",
		From:  "X",
		To:    "Y",
		Price: 20,
		Seat: &proto.Seat{
			Number:  1,
			Section: "A",
		},
	}, ticket)
	m.Nil(err)
}

func (m *TicketHandlerTestSuite) TestBookTicket_ShouldReturnErrorWhenManagerThrowsError() {
	u := user.User{Name: "TEST", EmalAddress: "abcd"}
	m.mockTicketService.EXPECT().BookTicket(u, "X", "Y").Return(ticket.Ticket{}, errors.New("unable to book"))
	ticket, err := m.ticketHandler.BookTicket(m.ctx, &proto.BookTicketRequest{
		Name:         "TEST",
		From:         "X",
		To:           "Y",
		EmailAddress: "abcd",
	})
	m.Empty(ticket)
	m.EqualError(err, "unable to book")
}

func (m *TicketHandlerTestSuite) TestBookTicket_ShouldReturnErrorWhenValidationFails() {
	ticket, err := m.ticketHandler.BookTicket(m.ctx, &proto.BookTicketRequest{
		Name:         "TEST",
		From:         "X",
		EmailAddress: "abcd",
	})
	m.Empty(ticket)
	m.EqualError(err, "field To value length must be at least 1 runes")
}

func (m *TicketHandlerTestSuite) TestGetTicket_ShouldBeSuccessful() {
	m.mockTicketService.EXPECT().GetTicket("abcd").Return(ticket.Ticket{
		Id:    "23312",
		From:  "X",
		To:    "Y",
		Price: 20,
		Seat: ticket.Seat{
			Number:  1,
			Section: "A",
		},
	}, nil)
	ticket, err := m.ticketHandler.GetTicket(m.ctx, &proto.GetTicketRequest{
		EmailAddress: "abcd",
	})
	m.Equal(&proto.BookTicketResponse{
		Id:    "23312",
		From:  "X",
		To:    "Y",
		Price: 20,
		Seat: &proto.Seat{
			Number:  1,
			Section: "A",
		},
	}, ticket)
	m.Nil(err)
}

func (m *TicketHandlerTestSuite) TestGetTicket_ShouldReturnErrorWhenManagerThrowsError() {
	m.mockTicketService.EXPECT().GetTicket("abcd").Return(ticket.Ticket{}, errors.New("unable to fetch"))
	ticket, err := m.ticketHandler.GetTicket(m.ctx, &proto.GetTicketRequest{
		EmailAddress: "abcd",
	})
	m.Empty(ticket)
	m.EqualError(err, "unable to fetch")
}

func (m *TicketHandlerTestSuite) TestGetTicket_ShouldReturnErrorWhenRequestValidationFails() {
	ticket, err := m.ticketHandler.GetTicket(m.ctx, &proto.GetTicketRequest{})
	m.Empty(ticket)
	m.EqualError(err, "field EmailAddress value length must be at least 1 runes")
}

func (m *TicketHandlerTestSuite) TestGetSeatsForGivenSection_ShouldBeSuccessful() {
	m.mockTicketService.EXPECT().GetSeatsPerSection("A").Return([]ticket.Seat{
		{
			Number:  1,
			Section: "A",
		},
	}, nil)
	seatResponse, err := m.ticketHandler.GetSeats(m.ctx, &proto.GetSeatsRequest{
		Section: "A",
	})
	m.Len(seatResponse.Seats, 1)
	m.Equal(seatResponse.Seats[0], &proto.Seat{
		Number:  1,
		Section: "A",
	})
	m.Nil(err)
}

func (m *TicketHandlerTestSuite) TestGetSeats_ShouldReturnErrorWhenManagerThrowsError() {
	m.mockTicketService.EXPECT().GetSeatsPerSection("A").Return(nil, errors.New("unable to fetch"))
	seatResponse, err := m.ticketHandler.GetSeats(m.ctx, &proto.GetSeatsRequest{
		Section: "A",
	})
	m.Empty(seatResponse)
	m.EqualError(err, "unable to fetch")
}

func (m *TicketHandlerTestSuite) TestGetSeats_ShouldReturnErrorWhenRequestValidationFails() {
	seatResponse, err := m.ticketHandler.GetSeats(m.ctx, &proto.GetSeatsRequest{})
	m.Empty(seatResponse)
	m.EqualError(err, "field Section value length must be at least 1 runes")
}

func (m *TicketHandlerTestSuite) TestCancelTicket_ShouldBeSuccessful() {
	m.mockTicketService.EXPECT().CancelBooking("ABC").Return(nil)
	empty, err := m.ticketHandler.CancelTicket(m.ctx, &proto.CancelTicketRequest{EmailAddress: "ABC"})
	m.Nil(err)
	m.Empty(empty)
}

func (m *TicketHandlerTestSuite) TestCancelBooking_ShouldReturnErrorWhenManagerThrowsError() {
	m.mockTicketService.EXPECT().CancelBooking("ABC").Return(errors.New("unable to cancel"))
	empty, err := m.ticketHandler.CancelTicket(m.ctx, &proto.CancelTicketRequest{EmailAddress: "ABC"})
	m.EqualError(err, "unable to cancel")
	m.Empty(empty)
}

func (m *TicketHandlerTestSuite) TestCancelBooking_ShouldReturnErrorWhenRequestValidationFails() {
	empty, err := m.ticketHandler.CancelTicket(m.ctx, &proto.CancelTicketRequest{})
	m.EqualError(err, "field EmailAddress value length must be at least 1 runes")
	m.Empty(empty)
}
