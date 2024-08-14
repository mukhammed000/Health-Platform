package mongo

import (
	"context"
	"log"

	u "analytics/storage"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	Db        *mongo.Database
	analytics u.AnalyticsService
}

func NewMongoStorage() (u.InitRoot, error) {
	uri := "mongodb+srv://mukhammed:1234@mycluster.xqxlv.mongodb.net/"

	clientOptions := options.Client().ApplyURI(uri).
		SetAuth(options.Credential{Username: "muhammad", Password: "1234"})

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error: Couldn't connect to the database.", err)
	}

	db := client.Database("lerning")

	return &Storage{Db: db}, err
}

func (s *Storage) Analytics() u.AnalyticsService {
	if s.analytics == nil {
		s.analytics = &AnalyticsStorage{s.Db}
	}
	return s.analytics
}
