package main

import (
	"api-gateway/api"
	"api-gateway/api/handler"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Connect(port string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func main() {
	conn1, err := Connect("localhost:50051")
	if err != nil {
		log.Fatal(err)
	}
	defer conn1.Close()

	conn2, err := Connect("localhost:50052")
	if err != nil {
		log.Fatal(err)
	}
	defer conn2.Close()

	conn3, err := Connect("localhost:50053")
	if err != nil {
		log.Fatal(err)
	}
	defer conn3.Close()

	conn4, err := Connect("localhost:50054")
	if err != nil {
		log.Fatal(err)
	}
	defer conn4.Close()

	conn5, err := Connect("localhost:50055")
	if err != nil {
		log.Fatal(err)
	}
	defer conn5.Close()

	server := api.New(&handler.Server{
		Usermanagement:      conn2,
		Gargardenmanagement: conn3,
		Sustainability:      conn4,
		Community:           conn5,
	})

	if err := server.Run(":7070"); err != nil {
		log.Fatal(err)
	}
}
