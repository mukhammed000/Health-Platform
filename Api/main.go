package main

import (
	"log"

	"api/api"
	"api/api/handler"
	anlytics "api/genproto/analytics"
	pb "api/genproto/users"

	_ "api/docs"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {


	userConn, err := grpc.Dial("auth:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while creating new client of USERS service: ", err.Error())
	}
	defer userConn.Close()

	analyticsConn, err := grpc.Dial("analytics:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while creating new client of ANALYTICS service: ", err.Error())
	}
	defer analyticsConn.Close()

	usc := pb.NewUsersServiceClient(userConn)
	asc := anlytics.NewAnalyticsServiceClient(analyticsConn)

	// Create a new handler with the clients
	h := handler.NewHandler(usc, asc)
	r := api.NewGin(h)

	// Start the Gin server
	err = r.Run("api-gateway:8080")
	if err != nil {
		log.Fatalln("Error while running server: ", err.Error())
	}
}
