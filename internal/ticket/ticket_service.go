package ticket

import (
	"railwai-ticket-api/internal/user"
	"github.com/sirupsen/logrus"
)

type TicketService interface {
	BookTicket(user user.User, from, to string) (Ticket, error)
	GetTicket(emailAddress string) (Ticket, error)
	GetSeatsPerSection(section string) ([]Seat, error)
	CancelBooking(emailAddress string) error
}

func NewTicketService(manager Manager) TicketService {
	return &ticketService{
		manager: manager,
	}
}

type ticketService struct {
	manager Manager
}

func (t *ticketService) BookTicket(user user.User, from, to string) (Ticket, error) {
	log := logrus.WithFields(logrus.Fields{
		"method": "BookTicket",
		"struct": "ticketService",
	})

	log.Infof("booking ticket for user: %v from: %v to: %v", user.EmalAddress, from, to)
	ticket, err := t.manager.BookTicket(user, from, to)
	if err != nil {
			log.Errorf("error while booking ticket %v ",err.Error());
			return Ticket{},err
	}
	return ticket,nil
}

func (t *ticketService) GetTicket(emailAddress string) (Ticket, error) {
	log := logrus.WithFields(logrus.Fields{
		"method": "GetTicket",
		"struct": "ticketService",
	})

	log.Infof("getting ticket for email address: %v", emailAddress)
	ticket,err := t.manager.GetTicket(emailAddress)
	if err!=nil{
		log.Errorf("error while getting ticket %v ",err.Error());
		return Ticket{},err
	}
	return ticket,nil
}

func (t *ticketService) GetSeatsPerSection(section string) ([]Seat, error) {
	log := logrus.WithFields(logrus.Fields{
		"method": "GetSeatsPerSection",
		"struct": "ticketService",
	})

	log.Infof("getting seats for section: %v", section)
	seats,err := t.manager.GetSeatsBySection(section)
	if err!=nil{
		log.Errorf("error while getting seats for section. %v ",err.Error());
		return nil,err
	}
	return seats,nil
}

func (t *ticketService) CancelBooking(emailAddress string) error {
	log := logrus.WithFields(logrus.Fields{
		"method": "CancelBooking",
		"struct": "ticketService",
	})
	log.Infof("cancelling ticket for user : %v", emailAddress)
	err := t.manager.CancelTicket(emailAddress)
	if err!=nil{
		log.Errorf("error while cancelling ticket %v ",err.Error());
		return err
	}
	return nil
}