package postgres

import (
	"auth/genproto/users"
	"database/sql"
	"errors"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db: db}
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (u *UserStorage) RegisterUser(req *users.Users) (*users.RegisterResponse, error) {
	query := `
		INSERT INTO users (user_id, email, password_hash, first_name, last_name, date_of_birth, gender, role)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING user_id
	`

	password, err := hashPassword(req.PasswordHash)
	if err != nil {
		log.Fatalln("Error while hashing the password")
	}

	_, err = u.db.Exec(query, req.UserId, req.Email, password, req.FirstName, req.LastName, req.DateOfBirth, req.Gender, req.Role)
	if err != nil {
		return nil, err
	}

	response := &users.RegisterResponse{
		UserId:       req.UserId,
		AccessToken:  req.AccessToken,
		RefreshToken: req.RefereshToken,
		ExpiresAt:    req.ExpiresAt,
	}

	return response, nil
}

func (u *UserStorage) LoginUser(req *users.LoginUserRequest) (*users.LoginResponse, error) {
	var storedPasswordHash string
	var userID string

	query := `
		SELECT user_id, password_hash
		FROM users
		WHERE email = $1
	`
	err := u.db.QueryRow(query, req.Email).Scan(&userID, &storedPasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	tokenQuery := `
		SELECT access_token, refresh_token, expires_at
		FROM tokens
		WHERE user_id = $1
	`
	var accessToken, refreshToken string
	var expiresAt time.Time

	err = u.db.QueryRow(tokenQuery, userID).Scan(&accessToken, &refreshToken, &expiresAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no tokens found for the user")
		}
		return nil, err
	}

	response := &users.LoginResponse{
		UserId:       userID,
		Token:        accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt.Format(time.RFC3339),
	}

	return response, nil
}

func (u *UserStorage) ValidateToken(req *users.ValidateTokenRequest) (*users.Empty, error) {
	var expiresAt time.Time

	query := `
		SELECT expires_at
		FROM tokens
		WHERE access_token = $1
	`

	err := u.db.QueryRow(query, req.Token).Scan(&expiresAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return &users.Empty{
				Text:   "Token does not exsists in database! ",
				IsDone: false,
			}, nil
		}
		return nil, err
	}

	if time.Now().After(expiresAt) {
		return &users.Empty{
			Text:   "Token is already expired",
			IsDone: false,
		}, nil
	}

	return &users.Empty{
		Text:   "Token is valid",
		IsDone: true,
	}, nil
}

func (u *UserStorage) RefreshToken(req *users.RefreshTokenRequest) (*users.TokenResponse, error) {
	// Implement the logic to refresh a token
	return nil, nil
}

func (u *UserStorage) ValidateEmail(req *users.VerifyEmailRequest) (*users.Empty, error) {
	// Implement the logic to validate an email
	return nil, nil
}

func (u *UserStorage) GetUserProfile(req *users.GetUserProfileRequest) (*users.UserProfileResponse, error) {
	var user users.UserProfileResponse

	query := `
		SELECT user_id, email, first_name, last_name, date_of_birth, gender, role
		FROM users
		WHERE user_id = $1 AND deleted_at = 0
	`

	err := u.db.QueryRow(query, req.UserId).Scan(
		&user.UserId,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.DateOfBirth,
		&user.Gender,
		&user.Role,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (u *UserStorage) UpdateUserProfile(req *users.UpdateUserProfileRequest) (*users.UserProfileResponse, error) {
	query := `
		UPDATE users
		SET first_name = $1, last_name = $2, date_of_birth = $3, gender = $4, updated_at = CURRENT_TIMESTAMP
		WHERE user_id = $5 AND deleted_at = 0
		RETURNING user_id, email, first_name, last_name, date_of_birth, gender, role
	`

	var updatedUser users.UserProfileResponse
	err := u.db.QueryRow(query, req.FirstName, req.LastName, req.DateOfBirth, req.Gender, req.UserId).Scan(
		&updatedUser.UserId,
		&updatedUser.Email,
		&updatedUser.FirstName,
		&updatedUser.LastName,
		&updatedUser.DateOfBirth,
		&updatedUser.Gender,
		&updatedUser.Role,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found or already deleted")
		}
		return nil, err
	}

	return &updatedUser, nil
}

func (u *UserStorage) DeleteUserProfile(req *users.DeleteUserProfileRequest) (*users.Empty, error) {
	query := `
		UPDATE users
		SET deleted_at = $1, updated_at = CURRENT_TIMESTAMP
		WHERE user_id = $2 AND deleted_at = 0
		RETURNING user_id
	`

	currentTime := time.Now().Unix()

	var userID string
	err := u.db.QueryRow(query, currentTime, req.UserId).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found or already deleted")
		}
		return nil, err
	}

	return &users.Empty{}, nil
}