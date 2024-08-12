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
		log.Println("Error registering user: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *UserService) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginResponse, error) {
	resp, err := s.stg.Users().LoginUser(req)
	if err != nil {
		log.Println("Error logging in user: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *UserService) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.Empty, error) {
	resp, err := s.stg.Users().ValidateToken(req)
	if err != nil {
		log.Println("Error validating token: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *UserService) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.TokenResponse, error) {
	resp, err := s.stg.Users().RefreshToken(req)
	if err != nil {
		log.Println("Error refreshing token: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *UserService) ValidateEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*pb.Empty, error) {
	resp, err := s.stg.Users().ValidateEmail(req)
	if err != nil {
		log.Println("Error validating email: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *UserService) GetUserProfile(ctx context.Context, req *pb.GetUserProfileRequest) (*pb.UserProfileResponse, error) {
	resp, err := s.stg.Users().GetUserProfile(req)
	if err != nil {
		log.Println("Error getting user profile: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *UserService) UpdateUserProfile(ctx context.Context, req *pb.UpdateUserProfileRequest) (*pb.UserProfileResponse, error) {
	resp, err := s.stg.Users().UpdateUserProfile(req)
	if err != nil {
		log.Println("Error updating user profile: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *UserService) DeleteUserProfile(ctx context.Context, req *pb.DeleteUserProfileRequest) (*pb.Empty, error) {
	resp, err := s.stg.Users().DeleteUserProfile(req)
	if err != nil {
		log.Println("Error deleting user profile: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *UserService) ChangePassword(ctx context.Context, req *pb.ChangePasswordReq) (*pb.Empty, error) {
	resp, err := s.stg.Users().ChangePassword(req)
	if err != nil {
		log.Println("Error changing password: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *UserService) EnterTheValidationCode(ctx context.Context, req *pb.VerificationCode) (*pb.Empty, error) {
	resp, err := s.stg.Users().EnterTheValidationCode(req)
	if err != nil {
		log.Println("Error while verificating the code: ", err)
		return nil, err
	}
	return resp, nil

}
