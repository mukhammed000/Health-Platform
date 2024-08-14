package main

import (
	"analytics/storage/mongo"
	"fmt"
	"log"
)

func main() {
	fmt.Println("rewrew")
	stg, err := mongo.NewMongoStorage()
	if err != nil {
		log.Fatalln("Error while connecting to database", err)
	}
	log.Println("Database connected successfully! ", stg)
	fmt.Println("rewrwe")

}
