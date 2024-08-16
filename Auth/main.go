package main

import (
	"auth/service"
	"auth/storage/postgres"
	"log"
	"net"

	pb "auth/genproto/users"

	"google.golang.org/grpc"
)

func main() {

	stg, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatalln("Error while connecting to database", err)
	}
	log.Println("Database connected successfully! ")

	lis, err := net.Listen("tcp", "auth:8081")
	if err != nil {
		log.Fatal("Error while creating TCP listener", err)
	}
	defer lis.Close()

	server := grpc.NewServer()
	service := service.NewUserService(stg)

	pb.RegisterUsersServiceServer(server, service)

	log.Println("Server listening at", "auth:8081")
	if server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)

	}

}
