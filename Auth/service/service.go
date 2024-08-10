package service

import (
	"context"
	"log"

	pb "auth/genproto/users"
	stg "auth/storage"
)

type UserService struct {
	stg stg.InitRoot
	pb.UnimplementedUsersServiceServer
}

func NewUserService(stg stg.InitRoot) *UserService {
	return &UserService{
		stg: stg,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, req *pb.Users) (*pb.RegisterResponse, error) {
	resp, err := s.stg.Users().RegisterUser(req)
	if err != nil {
		log.Printf("Error registering user: %v", err)
		return nil, err
	}
	return resp, nil
}

func (s *UserService) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginResponse, error) {
	resp, err := s.stg.Users().LoginUser(req)
	if err != nil {
		log.Printf("Error logging in user: %v", err)
		return nil, err
	}
	return resp, nil
}

func (s *UserService) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.TokenValidationResponse, error) {
	resp, err := s.stg.Users().ValidateToken(req)
	if err != nil {
		log.Printf("Error validating token: %v", err)
		return nil, err
	}
	return resp, nil
}

func (s *UserService) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.TokenResponse, error) {
	resp, err := s.stg.Users().RefreshToken(req)
	if err != nil {
		log.Printf("Error refreshing token: %v", err)
		return nil, err
	}
	return resp, nil
}

func (s *UserService) ValidateEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*pb.Empty, error) {
	resp, err := s.stg.Users().ValidateEmail(req)
	if err != nil {
		log.Printf("Error validating email: %v", err)
		return nil, err
	}
	return resp, nil
}

func (s *UserService) GetUserProfile(ctx context.Context, req *pb.GetUserProfileRequest) (*pb.UserProfileResponse, error) {
	resp, err := s.stg.Users().GetUserProfile(req)
	if err != nil {
		log.Printf("Error getting user profile: %v", err)
		return nil, err
	}
	return resp, nil
}

func (s *UserService) UpdateUserProfile(ctx context.Context, req *pb.UpdateUserProfileRequest) (*pb.UserProfileResponse, error) {
	resp, err := s.stg.Users().UpdateUserProfile(req)
	if err != nil {
		log.Printf("Error updating user profile: %v", err)
		return nil, err
	}
	return resp, nil
}

func (s *UserService) DeleteUserProfile(ctx context.Context, req *pb.DeleteUserProfileRequest) (*pb.Empty, error) {
	resp, err := s.stg.Users().DeleteUserProfile(req)
	if err != nil {
		log.Printf("Error deleting user profile: %v", err)
		return nil, err
	}
	return resp, nil
}
