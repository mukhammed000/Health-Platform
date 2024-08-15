package handler

import (
	"api/api/token"
	"api/genproto/users"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RegisterUser godoc
// @Summary Register a new user
// @Description Register a new user with email, password, and personal details
// @Security BearerAuth
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body users.RegisterUserRequest true "User registration details"
// @Success 200 {object} users.RegisterResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /auth/register [post]
func (h *Handler) Register(c *gin.Context) {
	var req users.RegisterUserRequest
	var user users.Users
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tkn := token.GenereteJWTToken(&user)

	res, err := h.Auth.RegisterUser(context.Background(), &users.Users{
		UserId:        uuid.NewString(),
		Email:         req.Email,
		PasswordHash:  req.Password,
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		DateOfBirth:   req.DateOfBirth,
		Gender:        req.Gender,
		Role:          "user",
		AccessToken:   tkn.AccessToken,
		RefereshToken: tkn.RefreshToken,
		ExpiresAt:     time.Now().Add(60 * time.Minute).String(),
		CreatedAt:     time.Now().String(),
		UpdatedAt:     time.Now().String(),
		DeletedAt:     0,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// LoginUser godoc
// @Summary Login a user
// @Description Login a user with email and password
// @Security BearerAuth
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body users.LoginUserRequest true "User login details"
// @Success 200 {object} users.LoginResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /auth/login [post]
func (h *Handler) Login(c *gin.Context) {
	var req users.LoginUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Auth.LoginUser(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// ValidateToken godoc
// @Summary Validate an access token
// @Description Validate a user's access token
// @Security BearerAuth
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body users.ValidateTokenRequest true "Token validation details"
// @Success 200 {object} users.Empty
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /auth/validateToken [post]
func (h *Handler) ValidateToken(c *gin.Context) {
	var req users.ValidateTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Auth.ValidateToken(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// RefreshToken godoc
// @Summary Refresh an access token
// @Description Refresh a user's access token using the refresh token
// @Security BearerAuth
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body users.RefreshTokenRequest true "Token refresh details"
// @Success 200 {object} users.TokenResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /auth/refreshToken [post]
func (h *Handler) RefreshToken(c *gin.Context) {
	var req users.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Auth.RefreshToken(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// ValidateEmail godoc
// @Summary Validate email
// @Description Validate a user's email
// @Security BearerAuth
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body users.VerifyEmailRequest true "Email validation details"
// @Success 200 {object} users.Empty
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /auth/validateEmail [post]
func (h *Handler) ValidateEmail(c *gin.Context) {
	var req users.VerifyEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Auth.ValidateEmail(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetProfileInfo godoc
// @Summary Get user profile information
// @Description Get user profile information by user ID
// @Security BearerAuth
// @Tags Users
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} users.UserProfileResponse
// @Failure 500 {string} string "Internal Server Error"
// @Router /users/profile/{user_id} [get]
func (h *Handler) GetProfileInfo(c *gin.Context) {
	userID := c.Param("user_id")
	req := &users.GetUserProfileRequest{
		UserId: userID,
	}

	res, err := h.Auth.GetUserProfile(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateProfile godoc
// @Summary Update user profile
// @Description Update user profile details
// @Security BearerAuth
// @Tags Users
// @Accept json
// @Produce json
// @Param body body users.UpdateUserProfileRequest true "User profile update details"
// @Success 200 {object} users.UserProfileResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /users/updateProfile [put]
func (h *Handler) UpdateProfile(c *gin.Context) {
	var req users.UpdateUserProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Auth.UpdateUserProfile(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteProfile godoc
// @Summary Delete user profile
// @Description Delete user profile by user ID
// @Security BearerAuth
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} users.Empty
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /users/deleteProfile/{user_id} [delete]
func (h *Handler) DeleteProfile(c *gin.Context) {
	userID := c.Param("user_id")
	var req users.DeleteUserProfileRequest

	req.UserId = userID

	res, err := h.Auth.DeleteUserProfile(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// ChangePassword godoc
// @Summary Change user password
// @Description Change user password by providing the current and new password\
// @Security BearerAuth
// @Tags Users
// @Accept json
// @Produce json
// @Param body body users.ChangePasswordReq true "Change password details"
// @Success 200 {object} users.Empty
// @Failure 400 {string} string "Error while changing the password"
// @Failure 500 {string} string "Error while changing the password"
// @Router /users/change-password [post]
func (h *Handler) ChangePassword(c *gin.Context) {
	var req users.ChangePasswordReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Auth.ChangePassword(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// VerificationCode godoc
// @Summary Verify a verification code
// @Description Verify the provided verification code and print the associated email if valid
// @Security BearerAuth
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body users.VerificationCode true "Verification code details"
// @Success 200 {object} users.Empty
// @Failure 400 {string} string "Invalid or expired verification code"
// @Failure 500 {string} string "Internal Server Error"
// @Router /auth/verifyCode [post]
func (h *Handler) VerificationCode(c *gin.Context) {
	var req users.VerificationCode

	// Bind the JSON request body to the req object
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Log the verification code for debugging
	fmt.Println("Verification code received:", req.VerificationCode)

	// Call the service method to validate the code
	res, err := h.Auth.EnterTheValidationCode(context.Background(), &req)
	if err != nil {
		if err.Error() == "verification code is invalid or expired" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, res)
}
