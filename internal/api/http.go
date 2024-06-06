package api

import (
	"fmt"                                // Importing fmt package for formatted I/O
	"github.com/gin-gonic/gin"           // Importing the Gin framework for creating the HTTP server
	"payment-system-one/internal/models" // Importing the package containing data models
	"payment-system-one/internal/ports"  // Importing the package containing port interfaces for dependency injection
)

// HTTPHandler struct holds the repository to interact with the database
type HTTPHandler struct {
	Repository ports.Repository
}

// NewHTTPHandler is a constructor function that returns a new HTTPHandler
func NewHTTPHandler(repository ports.Repository) *HTTPHandler {
	return &HTTPHandler{
		Repository: repository,
	}
}

// GetUserFromContext extracts the user from the request context
func (u *HTTPHandler) GetUserFromContext(c *gin.Context) (*models.User, error) {
	// Retrieve the user object stored in the context
	contextUser, exists := c.Get("user")
	if !exists {
		// Return an error if the user object is not found in the context
		return nil, fmt.Errorf("error getting user from context")
	}

	// Type assert the user object to the *models.User type
	user, ok := contextUser.(*models.User)
	if !ok {
		// Return an error if type assertion fails
		return nil, fmt.Errorf("an error occurred")
	}

	// Return the user object
	return user, nil
}

// GetTokenFromContext extracts the access token from the request context
func (u *HTTPHandler) GetTokenFromContext(c *gin.Context) (string, error) {
	// Retrieve the access token stored in the context
	tokenI, exists := c.Get("access_token")
	if !exists {
		// Return an error if the access token is not found in the context
		return "", fmt.Errorf("error getting access token")
	}

	// Type assert the token to the string type
	tokenstr := tokenI.(string)

	// Return the access token string
	return tokenstr, nil
}
