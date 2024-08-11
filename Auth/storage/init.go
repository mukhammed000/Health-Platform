package storage

import (
	"auth/genproto/users"
)

type InitRoot interface {
	Users() UsersService
}

type UsersService interface {
	// auth
	RegisterUser(req *users.Users) (*users.RegisterResponse, error)
	LoginUser(req *users.LoginUserRequest) (*users.LoginResponse, error)
	ValidateToken(req *users.ValidateTokenRequest) (*users.Empty, error)
	RefreshToken(req *users.RefreshTokenRequest) (*users.TokenResponse, error)
	ValidateEmail(req *users.VerifyEmailRequest) (*users.Empty, error)

	// users
	GetUserProfile(req *users.GetUserProfileRequest) (*users.UserProfileResponse, error)
	UpdateUserProfile(req *users.UpdateUserProfileRequest) (*users.UserProfileResponse, error)
	DeleteUserProfile(req *users.DeleteUserProfileRequest) (*users.Empty, error)
}
