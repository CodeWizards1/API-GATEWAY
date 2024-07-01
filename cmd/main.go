package main

import (
	"api-gateway/api"
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

	server := api.New(conn1, conn2)

	if err := server.Run(":7070"); err != nil {
		log.Fatal(err)
	}
}
