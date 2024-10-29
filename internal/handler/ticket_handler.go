package handler

import (
	"context"
	"fmt"
	errorPkg "railwai-ticket-api/internal/error"
	"railwai-ticket-api/internal/proto"
	"railwai-ticket-api/internal/ticket"
	"railwai-ticket-api/internal/user"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
)

type TicketHandler struct {
	proto.UnimplementedTicketServiceServer
	ticketService ticket.TicketService
}

func NewTicketHandler(ticketService ticket.TicketService) TicketHandler {
	return TicketHandler{
		ticketService: ticketService,
	}
}

func (t TicketHandler) BookTicket(ctx context.Context, bookTicketRequest *proto.BookTicketRequest) (*proto.BookTicketResponse, error) {
	log := logrus.WithFields(logrus.Fields{
		"method": "BookTicket",
		"struct": "TicketHandler",
	})

	if err := bookTicketRequest.Validate(); err != nil {
		bookTickerErr,ok := err.(proto.BookTicketRequestValidationError)
		if !ok{
			return nil, errorPkg.GetErrorResponseByCode(errorPkg.BadRequest)
		}
		log.Errorf("request is invalid.Error: %v",bookTickerErr.Error())
		msg := fmt.Sprintf("field %v %v",bookTickerErr.Field(),bookTickerErr.Reason())
		return nil, errorPkg.GetErrorResponse(codes.InvalidArgument,errorPkg.BadRequest,msg)
	}
	ticket, err := t.ticketService.BookTicket(user.User{
		Name:        bookTicketRequest.Name,
		EmalAddress: bookTicketRequest.EmailAddress,
	}, bookTicketRequest.From, bookTicketRequest.To)
	if err != nil {
		return nil, err
	}
	return mapTicketToResponse(ticket), nil
}

func (t TicketHandler) GetTicket(ctx context.Context, getTicketRequest *proto.GetTicketRequest) (*proto.BookTicketResponse, error) {
	log := logrus.WithFields(logrus.Fields{
		"method": "GetTicket",
		"struct": "TicketHandler",
	})
	
	if err := getTicketRequest.Validate(); err != nil {
		getTickerErr,ok := err.(proto.GetTicketRequestValidationError)
		if !ok{
			return nil, errorPkg.GetErrorResponseByCode(errorPkg.BadRequest)
		}
		log.Errorf("request is invalid.Error: %v",getTickerErr.Error())
		msg := fmt.Sprintf("field %v %v",getTickerErr.Field(),getTickerErr.Reason())
		return nil, errorPkg.GetErrorResponse(codes.InvalidArgument,errorPkg.BadRequest,msg)
	}
	
	ticket, err := t.ticketService.GetTicket(getTicketRequest.EmailAddress)
	if err != nil {
		return nil, err
	}
	return mapTicketToResponse(ticket), nil
}
func (t TicketHandler) GetSeats(ctx context.Context, getSeatsRequest *proto.GetSeatsRequest) (*proto.GetSeatsResponse, error) {
	log := logrus.WithFields(logrus.Fields{
		"method": "GetSeats",
		"struct": "TicketHandler",
	})
	
	if err := getSeatsRequest.Validate(); err != nil {
		getSeatsErr,ok := err.(proto.GetSeatsRequestValidationError)
		if !ok{
			return nil, errorPkg.GetErrorResponseByCode(errorPkg.BadRequest)
		}
		log.Errorf("request is invalid.Error: %v",getSeatsErr.Error())
		msg := fmt.Sprintf("field %v %v",getSeatsErr.Field(),getSeatsErr.Reason())
		return nil, errorPkg.GetErrorResponse(codes.InvalidArgument,errorPkg.BadRequest,msg)
	}
	seats, err := t.ticketService.GetSeatsPerSection(getSeatsRequest.Section)
	if err != nil {
		return nil, err
	}
	return mapSeatsToResponse(seats), nil
}
func (t TicketHandler) CancelTicket(ctx context.Context, cancelTicketRequest *proto.CancelTicketRequest) (*empty.Empty, error) {
	log := logrus.WithFields(logrus.Fields{
		"method": "cancelTicketRequest",
		"struct": "TicketHandler",
	})
	
	if err := cancelTicketRequest.Validate(); err != nil {
		cancelTicketErr,ok := err.(proto.CancelTicketRequestValidationError)
		if !ok{
			return nil, errorPkg.GetErrorResponseByCode(errorPkg.BadRequest)
		}
		log.Errorf("request is invalid.Error: %v",cancelTicketErr.Error())
		msg := fmt.Sprintf("field %v %v",cancelTicketErr.Field(),cancelTicketErr.Reason())
		return nil, errorPkg.GetErrorResponse(codes.InvalidArgument,errorPkg.BadRequest,msg)
	}
	
	err := t.ticketService.CancelBooking(cancelTicketRequest.EmailAddress)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func mapTicketToResponse(t ticket.Ticket) *proto.BookTicketResponse {
	return &proto.BookTicketResponse{
		Id:    t.Id,
		From:  t.From,
		To:    t.To,
		Price: t.Price,
		Seat: &proto.Seat{
			Number:  t.Seat.Number,
			Section: t.Seat.Section,
		},
	}
}

func mapSeatsToResponse(seats []ticket.Seat) *proto.GetSeatsResponse {
	seatsResponse := []*proto.Seat{}
	for _, v := range seats {
		seatsResponse = append(seatsResponse, &proto.Seat{
			Number:  v.Number,
			Section: v.Section,
		})
	}
	return &proto.GetSeatsResponse{
		Seats: seatsResponse,
	}
}
