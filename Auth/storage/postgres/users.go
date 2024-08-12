package postgres

import (
	"auth/genproto/users"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	h "auth/helper"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserStorage struct {
	db  *sql.DB
	rdb *redis.Client
}

func NewUserStorage(db *sql.DB, rdb *redis.Client) *UserStorage {
	return &UserStorage{
		db:  db,
		rdb: rdb,
	}
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
		INSERT INTO users (user_id, email, password_hash, first_name, last_name, date_of_birth, gender, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING user_id
	`

	password, err := hashPassword(req.PasswordHash)
	if err != nil {
		log.Fatalln("Error while hashing the password")
	}

	_, err = u.db.Exec(query, req.UserId, req.Email, password, req.FirstName, req.LastName, req.DateOfBirth, req.Gender, req.Role, req.CreatedAt, req.UpdatedAt)
	if err != nil {
		return nil, err
	}

	response := &users.RegisterResponse{
		UserId:       req.UserId,
		AccessToken:  req.AccessToken,
		RefreshToken: req.RefereshToken,
		ExpiresAt:    req.ExpiresAt,
	}

	token_query := `
	INSERT INTO tokens (user_id, token_id, access_token, refresh_token, expires_at, created_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err = u.db.Exec(token_query, req.UserId, uuid.NewString(), req.AccessToken, req.RefereshToken, req.ExpiresAt, req.CreatedAt)
	if err != nil {
		return nil, err
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
	var expiresAt string

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
		ExpiresAt:    expiresAt,
	}

	return response, nil
}

func (u *UserStorage) ValidateToken(req *users.ValidateTokenRequest) (*users.Empty, error) {
	var expiresAt string

	query := `
		SELECT expires_at
		FROM tokens
		WHERE access_token = $1
	`

	err := u.db.QueryRow(query, req.Token).Scan(&expiresAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return &users.Empty{
				Text:   "Token does not exist in the database!",
				IsDone: false,
			}, nil
		}
		return nil, err
	}

	expiresAt = strings.Split(expiresAt, " m=")[0]

	const layout = "2006-01-02 15:04:05.999999999 -0700 MST"

	expiredAt, err := time.Parse(layout, expiresAt)
	if err != nil {
		log.Fatalln("Error while parsing the time", err)
	}

	if time.Now().After(expiredAt) {
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
	code := rand.Intn(899999) + 10000

	from := "muhammadjonxudaynazarov226@gmail.com"
	password := "rgkt oivo pyst zplt"
	err := h.SendVerificationCode(h.Params{
		From:     from,
		Password: password,
		To:       req.Email,
		Message:  fmt.Sprintf("Hi, here is your verification code: %d", code),
		Code:     fmt.Sprint(code),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to send verification email: %v", err)
	}

	err = u.rdb.Set(context.Background(), fmt.Sprint(code), req.Email, time.Minute*10).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to store verification code in Redis: %v", err)
	}

	return &users.Empty{
		Text:   "A verification code has been sent to your email address.",
		IsDone: true,
	}, nil
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
		SET deleted_at = $1
		WHERE user_id = $2 AND deleted_at = 0

	`

	currentTime := time.Now().Unix()

	res, err := u.db.Exec(query, currentTime, req.UserId)
	if err != nil {
		return nil, errors.New("error executing the query: " + err.Error())
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, errors.New("error retrieving rows affected: " + err.Error())
	}
	if rowsAffected == 0 {
		return nil, errors.New("user not found or already deleted")
	}

	return &users.Empty{Text: "User deleted successfully", IsDone: true}, nil
}

func (u *UserStorage) ChangePassword(req *users.ChangePasswordReq) (*users.Empty, error) {
	var currentPasswordHash string
	query := `
		SELECT password_hash
		FROM users
		WHERE user_id = $1 AND deleted_at = 0
	`
	err := u.db.QueryRow(query, req.UserId).Scan(&currentPasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return &users.Empty{IsDone: false, Text: "User not found or already deleted"}, nil
		}
		return nil, fmt.Errorf("could not fetch user: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(currentPasswordHash), []byte(req.CurrentPassword))
	if err != nil {
		return &users.Empty{IsDone: false, Text: "Current password is incorrect"}, nil
	}

	hashedNewPassword, err := hashPassword(req.NewPassword)
	if err != nil {
		return nil, fmt.Errorf("could not hash new password: %v", err)
	}

	updateQuery := `
		UPDATE users
		SET password_hash = $1, updated_at = CURRENT_TIMESTAMP
		WHERE user_id = $2 AND deleted_at = 0
	`
	result, err := u.db.Exec(updateQuery, hashedNewPassword, req.UserId)
	if err != nil {
		return &users.Empty{IsDone: false, Text: "Failed to update password"}, fmt.Errorf("could not change password: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &users.Empty{IsDone: false, Text: "Failed to update password"}, fmt.Errorf("could not determine rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return &users.Empty{IsDone: false, Text: "Failed to update password"}, fmt.Errorf("no rows were updated")
	}

	return &users.Empty{IsDone: true, Text: "Password successfully changed"}, nil
}

func (u *UserStorage) EnterTheValidationCode(req *users.VerificationCode) (*users.Empty, error) {

	storedEmail, err := u.rdb.Get(context.Background(), req.VerificationCode).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("verification code is invalid or expired")
	} else if err != nil {
		return nil, fmt.Errorf("failed to retrieve verification code from Redis: %v", err)
	}

	err = u.rdb.Del(context.Background(), req.VerificationCode).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to delete verification code from Redis: %v", err)
	}

	_, err = u.db.Exec(
		"UPDATE users SET email_verified = true WHERE email = $1",
		storedEmail,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update email verification status in database: %v", err)
	}

	return &users.Empty{
		Text:   "Verification code is valid",
		IsDone: true,
	}, nil
}
