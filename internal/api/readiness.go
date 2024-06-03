package api

import (
	"github.com/dgrijalva/jwt-go"        // Importing JWT package for token generation and validation
	"github.com/gin-gonic/gin"           // Importing the Gin framework for creating the HTTP server
	"net/http"                          // Importing net/http package for HTTP constants and functions
	"os"                                // Importing os package for OS-related functionality
	"payment-system-one/internal/middleware" // Importing the package containing middleware functions
	"payment-system-one/internal/models"     // Importing the package containing data models
	"payment-system-one/internal/util"       // Importing the package containing utility functions
	"time"                              // Importing time package for handling durations
)

// Readiness is to check if server is up
func (u *HTTPHandler) Readiness(c *gin.Context) {
	data := "server is up and running"

	// Send a readiness response using the utility function
	util.Response(c, "Ready to go", 200, data, nil)
}

// CreateUser handles the creation of a new user
func (u *HTTPHandler) CreateUser(c *gin.Context) {
	var user *models.User
	// Bind the request body to the user model
	if err := c.ShouldBind(&user); err != nil {
		util.Response(c, "invalid request", 400, "bad request body", nil)
		return
	}

	// Validate user email (add your validation logic here)

	// Validate user password (add your validation logic here)

	// Persist user information in the database
	err := u.Repository.CreateUser(user)
	if err != nil {
		util.Response(c, "user not created", 400, err.Error(), nil)
		return
	}
	util.Response(c, "user created", 200, "success", nil)
}

// LoginUser handles user login and token generation
func (u *HTTPHandler) LoginUer(c *gin.Context) {
	var loginRequest *models.LoginRequest
	// Bind the request body to the login request model
	if err := c.ShouldBind(&loginRequest); err != nil {
		util.Response(c, "invalid request", 400, "bad request body", nil)
		return
	}
	// Check if email or password is empty
	if loginRequest.Email == "" || loginRequest.Password == "" {
		util.Response(c, "Please enter your email or password", 400, "bad request body", nil)
		return
	}

	// Find the user by email
	user, err := u.Repository.FindUserByEmail(loginRequest.Email)
	if err != nil {
		util.Response(c, "user does not exist", 404, "user not found", nil)
		return
	}
	// Check if the user account is locked after 3 failed attempts
	if user.LoginCounter >= 3 {
		user.IsLocked = true
		user.UpdatedAt = time.Now()
		err = u.Repository.UpdateUser(user)
		if err != nil {
			return
		}
		util.Response(c, "Your account has been locked after 3 failed attempts, contact customer care for assistance", 200, "success", nil)
		return
	}

	// Check if the password is correct
	if user.Password != loginRequest.Password {
		user.LoginCounter++
		err := u.Repository.UpdateUser(user)
		if err != nil {
			util.Response(c, "internal server error", 500, "user not found", nil)
			return
		}
		util.Response(c, "password mismatch", 404, "user not found", nil)
		return
	}

	// Generate access and refresh tokens
	accessClaims, refreshClaims := middleware.GenerateClaims(user.Email)
	secret := os.Getenv("JWT_SECRET")

	accessToken, err := middleware.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)
	if err != nil {
		util.Response(c, "error generating access token", 500, "error generating access token", nil)
		return
	}
	refreshToken, err := middleware.GenerateToken(jwt.SigningMethodHS256, refreshClaims, &secret)
	if err != nil {
		util.Response(c, "error generating refresh token", 500, "error generating refresh token", nil)
		return
	}
	// Set tokens in the response headers
	c.Header("access_token", *accessToken)
	c.Header("refresh_token", *refreshToken)

	// Send a successful login response
	util.Response(c, "login successful", http.StatusOK, gin.H{
		"user":          user,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil)
}

// GetUserByEmail handles the retrieval of a user by their email
func (u *HTTPHandler) GetUserByEmail(c *gin.Context) {
	// Extract the user from the context
	_, err := u.GetUserFromContext(c)
	if err != nil {
		util.Response(c, "User not logged in", 500, "user not found", nil)
		return
	}

	// Get the email query parameter
	email := c.Query("email")
	if email == "" {
		util.Response(c, "email is required", 400, "email is required", nil)
		return
	}

	// Find the user by email
	user, err := u.Repository.FindUserByEmail(email)
	if err != nil {
		util.Response(c, "user not found", 500, "user not found", nil)
		return
	}

	// Send a response with the user information
	util.Response(c, "user found", 200, user, nil)
}

// HTTP Status Code Explanation:
// 100 ---- Informational responses
// 200 ---- Success (200: OK, 201: Created, 202: Accepted)
// 300 ---- Redirection
// 400 ---- Client errors (e.g., 400: Bad Request)
// 500 ---- Server errors (e.g., 500: Internal Server Error)

// Note: Add validation logic and appropriate error handling as needed.
