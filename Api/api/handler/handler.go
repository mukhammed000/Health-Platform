package handler

import (
	analytics "api/genproto/analytics"
	pb "api/genproto/users"
	"api/kafka"
)

type Handler struct {
	Auth      pb.UsersServiceClient
	Analytics analytics.AnalyticsServiceClient
	Kafka     kafka.KafkaProducer
}

func NewHandler(auth pb.UsersServiceClient, analytics analytics.AnalyticsServiceClient) *Handler {
	return &Handler{
		Auth:      auth,
		Analytics: analytics,
	}
}
