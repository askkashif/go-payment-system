package util

import (
	"github.com/gin-gonic/gin"   // Importing Gin framework for HTTP handling
	"golang.org/x/crypto/bcrypt" // Importing bcrypt for password hashing
	"net/http"                   // Importing net/http for HTTP status codes
	"time"                       // Importing time for timestamp
)

// Response is a utility function for customized HTTP responses.
func Response(c *gin.Context, message string, status int, data interface{}, errs []string) {
	// Constructing the response data with message, data, errors, status, and timestamp
	responsedata := gin.H{
		"message":   message,
		"data":      data,
		"errors":    errs,
		"status":    http.StatusText(status),                  // Converting HTTP status code to text
		"timestamp": time.Now().Format("2006-01-02 15:04:05"), // Formatting current time as timestamp
	}

	// Sending the response data as JSON with an indentation corresponding to the provided status code
	c.IndentedJSON(status, responsedata)
}

// HashPassword hashes a password using bcrypt.
func HashPassword(password string) (string, error) {
	// Generating a hashed password using bcrypt with a cost of 14
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	// Converting the hashed password bytes to a string and returning it along with any error
	return string(bytes), err
}
