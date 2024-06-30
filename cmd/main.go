package main

import (
	"api-gateway/api"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	server := api.New(conn)

	if err := server.Run(":7070"); err != nil {
		log.Fatal(err)
	}
}
