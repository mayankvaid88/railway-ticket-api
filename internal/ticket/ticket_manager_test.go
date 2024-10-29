package ticket

import (
	"railwai-ticket-api/internal/config"
	"railwai-ticket-api/internal/user"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type ManagerTestSuite struct {
	suite.Suite
}

func TestManagerTestSuite(t *testing.T) {
	suite.Run(t, new(ManagerTestSuite))
}

func (m *ManagerTestSuite) SetupSuite() {
	UUIDFunc = func() string {
		return "abcd"
	}
	Now = func() time.Time {
		return time.Date(2023, 01, 01, 10, 10, 0, 0, time.Local)
	}
}

func (m *ManagerTestSuite) TearDownSuite() {
	UUIDFunc = uuid.NewString
	Now = time.Now
}

func (m *ManagerTestSuite) TestBookTicket_ShouldReturnTicket() {

	manager := NewManager([]config.Section{
		{
			Name:          "A",
			NumberOfSeats: 30,
		},
	},20)

	ticket, err := manager.BookTicket(user.User{Name: "TEST", EmalAddress: "abcd"}, "A", "B")
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

func (m *ManagerTestSuite) TestBookTicket_ShouldReturnErrorWhenSeatIsNotAvailable() {

	manager := NewManager([]config.Section{
		{
			Name:          "A",
			NumberOfSeats: 0,
		},
	},20)

	ticket, err1 := manager.BookTicket(user.User{Name: "TEST", EmalAddress: "def"}, "A", "B")
	m.EqualError(err1, "seat not available")
	m.Empty(ticket)
}

func (m *ManagerTestSuite) TestGetTicket_ShouldReturnTicketForGivenUser() {

	manager := NewManager([]config.Section{
		{
			Name:          "A",
			NumberOfSeats: 1,
		},
	},20)

	_, err1 := manager.BookTicket(user.User{Name: "TEST", EmalAddress: "def"}, "A", "B")
	m.Nil(err1)
	ticket, err := manager.GetTicket("def")
	m.Nil(err)
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
}

func (m *ManagerTestSuite) TestGetTicket_ShouldReturnErrorWhenTicketNotFound() {

	manager := NewManager([]config.Section{
		{
			Name:          "A",
			NumberOfSeats: 1,
		},
	},20)

	ticket, err := manager.GetTicket("def")
	m.EqualError(err, "ticket not found")
	m.Empty(ticket)
}

func (m *ManagerTestSuite) TestGetSeats_ShouldReturnAllocatedSeatsForGivenSection() {

	manager := NewManager([]config.Section{
		{
			Name:          "A",
			NumberOfSeats: 1,
		},
	},20)

	_, err1 := manager.BookTicket(user.User{Name: "TEST", EmalAddress: "def"}, "A", "B")
	m.Nil(err1)
	seats, err := manager.GetSeatsBySection("A")
	m.Nil(err)
	m.Len(seats, 1)
	m.Equal(seats[0], Seat{
		Number: 1,
		Section: "A",
	})
}


func (m *ManagerTestSuite) TestGetSeats_ShouldReturnErrorWhenSeatsForGivenSectionIsNotAllocated() {

	manager := NewManager([]config.Section{
		{
			Name:          "A",
			NumberOfSeats: 1,
		},
	},20)

	seats, err := manager.GetSeatsBySection("D")
	m.EqualError(err,"seats are not allocated for given section")
	m.Nil(seats)
}

func (m *ManagerTestSuite) TestCancelTicket_ShouldDeallocateTheSeat() {

	manager := NewManager([]config.Section{
		{
			Name:          "A",
			NumberOfSeats: 1,
		},
	},20)

	_, err1 := manager.BookTicket(user.User{Name: "TEST", EmalAddress: "def"}, "A", "B")
	m.Nil(err1)
	err := manager.CancelTicket("def")
	m.Nil(err)
	_, err1 = manager.BookTicket(user.User{Name: "TEST", EmalAddress: "def"}, "A", "B")
	m.Nil(err1)
}

func (m *ManagerTestSuite) TestCancelTicket_ShouldReturnErrorWhenTicketForGivenUserNotFound() {

	manager := NewManager([]config.Section{
		{
			Name:          "A",
			NumberOfSeats: 1,
		},
	},20)

	err := manager.CancelTicket("def")
	m.EqualError(err,"ticket not found")
}
