package middleware

import (
	"github.com/gin-gonic/gin"            // Importing the Gin framework for creating the HTTP server
	"log"                                 // Importing log package for logging
	"net/http"                            // Importing net/http package for HTTP constants and functions
	"os"                                  // Importing os package for OS-related functionality
	"payment-system-four/internal/models" // Importing the package containing data models
)

// AuthorizeAdmin is a middleware function that authorizes admin users
func AuthorizeAdmin(findUserByEmail func(string) (*models.User, error), tokenInBlacklist func(*string) bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *models.User
		var errors error
		secret := os.Getenv("JWT_SECRET") // Retrieve the JWT secret from environment variables
		accToken := GetTokenFromHeader(c) // Extract the token from the request header

		// Authorize the token and extract claims
		accessToken, accessClaims, err := AuthorizeToken(&accToken, &secret)
		if err != nil {
			log.Printf("authorize access token errors: %s\n", err.Error())
			RespondAndAbort(c, "", http.StatusUnauthorized, nil, []string{"unauthorized"})
			return
		}

		// Check if the token is blacklisted or expired (code commented out for now)
		// if tokenInBlacklist(&accessToken.Raw) || IsTokenExpired(accessClaims) {
		// 	c.AbortWithStatusJSON(http.StatusBadRequest, "unauthorized route ")
		// }

		// Extract user email from token claims
		if email, ok := accessClaims["user_email"].(string); ok {
			// Find the user by email
			if user, errors = findUserByEmail(email); errors != nil {
				log.Printf("find user by email errors: %v\n", err)
				RespondAndAbort(c, "", http.StatusNotFound, nil, []string{"user not found"})
				return
			}
		} else {
			log.Printf("user email is not string\n")
			RespondAndAbort(c, "", http.StatusInternalServerError, nil, []string{"internal server errors"})
			return
		}

		// Set the user and token as context parameters
		c.Set("user", user)
		c.Set("access_token", accessToken.Raw)

		// Call the next handler
		c.Next()
	}
}
