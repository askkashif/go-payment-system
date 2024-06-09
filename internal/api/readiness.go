package api

import (
	"github.com/gin-gonic/gin"          // Importing the Gin framework for creating the HTTP server
	"payment-system-four/internal/util" // Importing the package containing utility functions
)

// Readiness is to check if server is up
func (u *HTTPHandler) Readiness(c *gin.Context) {
	data := "server is up and running"

	// Send a readiness response using the utility function
	util.Response(c, "Ready to go", 200, data, nil)
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
