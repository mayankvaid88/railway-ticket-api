package main

import (
	"net"
	"railwai-ticket-api/internal/config"
	"railwai-ticket-api/internal/handler"
	pb "railwai-ticket-api/internal/proto"
	"railwai-ticket-api/internal/ticket"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const(
    ConfigFilePath = "internal/config/config.json"
)

func main(){
    log := logrus.WithFields(logrus.Fields{
		"method": "main",
	})
	lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    grpcServer := grpc.NewServer()
    config := config.NewConfig(ConfigFilePath)
    tickettManager := ticket.NewManager(config.GetSections(),config.GetTicketCost())
    ticketHandler := handler.NewTicketHandler(ticket.NewTicketService(tickettManager))
    pb.RegisterTicketServiceServer(grpcServer, ticketHandler)
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}