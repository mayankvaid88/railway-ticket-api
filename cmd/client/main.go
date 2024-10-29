package main

import (
	"context"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "railwai-ticket-api/internal/proto"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
			log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTicketServiceClient(conn)

	bookTicket(client, "Alice", "New York", "Los Angeles", "alice@example.com")

	getTicket(client, "alice@example.com")

	getSeats(client, "Economy")

	cancelTicket(client, "alice@example.com")
}

func bookTicket(client pb.TicketServiceClient, name, from, to, email string) {
	req := &pb.BookTicketRequest{
			Name:         name,
			From:         from,
			To:           to,
			EmailAddress: email,
	}

	resp, err := client.BookTicket(context.Background(), req)
	if err != nil {
			log.Fatalf("could not book ticket: %v", err)
	}

	log.Printf("Booked Ticket ID: %s, From: %s, To: %s, Price: %f, Seat: %+v", resp.Id, resp.From, resp.To, resp.Price, resp.Seat)
}

func getTicket(client pb.TicketServiceClient, email string) {
	req := &pb.GetTicketRequest{
			EmailAddress: email,
	}

	resp, err := client.GetTicket(context.Background(), req)
	if err != nil {
			log.Fatalf("could not get ticket: %v", err)
	}

	log.Printf("Ticket Details: ID: %s, From: %s, To: %s, Price: %f, Seat: %+v", resp.Id, resp.From, resp.To, resp.Price, resp.Seat)
}

func getSeats(client pb.TicketServiceClient, section string) {
	req := &pb.GetSeatsRequest{
			Section: section,
	}

	resp, err := client.GetSeats(context.Background(), req)
	if err != nil {
			log.Fatalf("could not get seats: %v", err)
	}

	log.Printf("Available Seats in %s: %+v", section, resp.Seats)
}

func cancelTicket(client pb.TicketServiceClient, email string) {
	req := &pb.CancelTicketRequest{
			EmailAddress: email,
	}

	_, err := client.CancelTicket(context.Background(), req)
	if err != nil {
			log.Fatalf("could not cancel ticket: %v", err)
	}

	log.Printf("Successfully canceled ticket for email: %s", email)
}