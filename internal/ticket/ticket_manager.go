package ticket

import (
	"railwai-ticket-api/internal/config"
	err "railwai-ticket-api/internal/error"
	"railwai-ticket-api/internal/user"
	"sync"
	"time"

	"github.com/google/uuid"
)

var UUIDFunc = uuid.NewString
var Now = time.Now

type Manager interface {
	BookTicket(u user.User,from,to string) (Ticket, error)
	GetTicket(emailAddress string) (Ticket,error)
	GetSeatsBySection(section string) ([]Seat,error)
	CancelTicket(emailAddress string) error
}

func NewManager(sections []config.Section,ticketCost float32) Manager {
	availableSeats := []*Seat{}
	for _,v := range sections{
		for i:=1;i<=v.NumberOfSeats;i++ {
			availableSeats = append(availableSeats, &Seat{
				Number: int32(i),
				Section: v.Name,
			})
		}
	}
	return &manager{
		availableSeats: availableSeats,
		allocatedSeatsPerSection: map[string][]Seat{},
		userToReceipt:            map[string]*Ticket{},
		ticketCost: ticketCost,
	}
}

type manager struct {
	availableSeats []*Seat
	allocatedSeatsPerSection 	map[string][]Seat
	userToReceipt            	map[string]*Ticket
	mu                       	sync.RWMutex
	ticketCost								float32
}

func (m *manager) BookTicket(u user.User,from,to string) (Ticket, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	seat := m.allocateSeat()
	if seat != nil {
		t := NewTicket(UUIDFunc(), from,to, m.ticketCost,*seat)
		m.userToReceipt[u.EmalAddress] = &t
		return t, nil
	}
	return Ticket{}, err.GetErrorResponseByCode(err.SeatNotAvailable)
}

func (m *manager) allocateSeat() (seat *Seat) {
		if len(m.availableSeats)!=0 {
			seat = m.availableSeats[0]
			m.availableSeats = append(m.availableSeats[:0], m.availableSeats[0+1:]...)
			m.allocatedSeatsPerSection[seat.Section] = append(m.allocatedSeatsPerSection[seat.Section], *seat)
			return
		}
	return
}

func (m *manager) GetTicket(emailAddress string) (Ticket,error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	ticket := m.userToReceipt[emailAddress]
	if ticket==nil{
		return Ticket{},err.GetErrorResponseByCode(err.TicketNotFound)
	}
	return *ticket,nil
}

func (m *manager) GetSeatsBySection(section string) ([]Seat,error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	seats := m.allocatedSeatsPerSection[section]
	if seats==nil{
		return nil,err.GetErrorResponseByCode(err.SeatNotAllocationForSection)
	}
	return seats,nil
}

func (m *manager) CancelTicket(emailAddress string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	ticket := m.userToReceipt[emailAddress]
	if ticket==nil{
		return err.GetErrorResponseByCode(err.TicketNotFound)
	}
	seat := ticket.Seat
	delete(m.userToReceipt,emailAddress)
	var isSeatFound bool
	for k,v := range m.allocatedSeatsPerSection[seat.Section]{
			if v==seat{
				m.allocatedSeatsPerSection[seat.Section] = append(m.allocatedSeatsPerSection[seat.Section][:k], 
					m.allocatedSeatsPerSection[seat.Section][k+1:]...)
					isSeatFound = true
					break
			}
	}
	if !isSeatFound{
		return err.GetErrorResponseByCode(err.InternalServerError)
	}
	m.availableSeats = append(m.availableSeats,&seat)
	return nil
}