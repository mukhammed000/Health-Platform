package main

import (
	"log"

	"api/api"
	"api/api/handler"
	"api/config"
	pb "api/genproto/users"

	_ "api/docs"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	cfg := config.Load()

	userConn, err := grpc.NewClient(cfg.UsersPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while creating new client of users service: ", err.Error())
	}
	defer userConn.Close()

	usc := pb.NewUsersServiceClient(userConn)
	h := handler.NewHandler(usc)
	r := api.NewGin(h)

	err = r.Run(cfg.ApiPort)
	if err != nil {
		log.Fatalln("Error while running server: ", err.Error())
	}
}
