package handler

import (
	analytics "api/genproto/analytics"
	pb "api/genproto/users"
)

type Handler struct {
	Auth      pb.UsersServiceClient
	Analytics analytics.AnalyticsServiceClient
}

func NewHandler(auth pb.UsersServiceClient, analytics analytics.AnalyticsServiceClient) *Handler {
	return &Handler{
		Auth:      auth,
		Analytics: analytics,
	}
}
