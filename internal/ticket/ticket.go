package ticket

type Ticket struct{
	Id 		string
	From 	string
	To 		string
	Price float32
	Seat	Seat
}

func NewTicket(id, from, to string, price float32,seat Seat) Ticket{
	return Ticket{
		Id: id,
		From: from,
		To: to,
		Price: price,
		Seat: seat,
	}
}