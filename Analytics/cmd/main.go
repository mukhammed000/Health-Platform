package main

import (
	"analytics/config"
	"analytics/service"
	"analytics/storage/mongo"
	"log"
	"net"

	pb "analytics/genproto/analytics"

	"google.golang.org/grpc"
)

func main() {

	cfg := config.Load()

	stg, err := mongo.NewMongoStorage()
	if err != nil {
		log.Fatalln("Error while connecting to database", err)
	}
	log.Println("Database connected successfully! ")

	lis, err := net.Listen("tcp", cfg.HTTPPort)
	if err != nil {
		log.Fatal("Error while creating TCP listener", err)
	}
	defer lis.Close()

	server := grpc.NewServer()
	service := service.NewAnalyticsService(stg)

	pb.RegisterAnalyticsServiceServer(server, service)

	log.Println("Server listening at", lis.Addr())
	if server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)

	}

}
