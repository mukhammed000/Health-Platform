package handler

import (
	pb "api/genproto/users"
)

type Handler struct {
	Auth pb.UsersServiceClient
}

func NewHandler(auth pb.UsersServiceClient) *Handler {
	return &Handler{
		Auth: auth,
	}
}
