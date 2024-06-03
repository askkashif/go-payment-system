package middleware

import (
	"fmt"                      // Importing fmt package for formatted I/O
	"github.com/dgrijalva/jwt-go"  // Importing jwt-go package for working with JWT tokens
	"github.com/gin-gonic/gin" // Importing gin package for HTTP web framework
	"log"                      // Importing log package for logging
	"payment-system-one/internal/util" // Importing util package for utility functions
	"time"                     // Importing time package for time-related functions
)

// Constants defining the validity duration for access and refresh tokens
const AccessTokenValidity = time.Hour * 24
const RefreshTokenValidity = time.Hour * 24

// Claims struct defines the structure of JWT claims including standard claims
type Claims struct {
	UserEmail string `json:"email"`
	jwt.StandardClaims
}

// GenerateClaims generates JWT claims for access and refresh tokens
func GenerateClaims(email string) (jwt.MapClaims, jwt.MapClaims) {
	log.Println("generate  claim function", email)
	accessClaims := jwt.MapClaims{
		"user_email": email,
		"exp":        time.Now().Add(AccessTokenValidity).Unix(),
	}

	refreshClaims := jwt.MapClaims{
		"exp": time.Now().Add(RefreshTokenValidity).Unix(),
		"sub": 1,
	}

	return accessClaims, refreshClaims
}

// GenerateToken generates a JWT token with given signing method, claims, and secret
func GenerateToken(signMethod *jwt.SigningMethodHMAC, claims jwt.MapClaims, secret *string) (*string, error) {
	token := jwt.NewWithClaims(signMethod, claims) // Create a new token with the specified claims
	tokenString, err := token.SignedString([]byte(*secret)) // Sign the token with the secret
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

// GetTokenFromHeader extracts the token string from the Authorization header
func GetTokenFromHeader(c *gin.Context) string {
	authHeader := c.Request.Header.Get("Authorization") // Get the Authorization header
	if len(authHeader) > 8 {
		return authHeader[7:] // Return the token part, skipping the "Bearer " prefix
	}
	return ""
}

// verifyToken verifies the token with the given claims and secret
func verifyToken(tokenString *string, claims jwt.MapClaims, secret *string) (*jwt.Token, error) {
	parser := &jwt.Parser{SkipClaimsValidation: true}
	return parser.ParseWithClaims(*tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(*secret), nil
	})
}

// AuthorizeToken checks if a token is valid and returns the token and its claims
func AuthorizeToken(token *string, secret *string) (*jwt.Token, jwt.MapClaims, error) {
	if token != nil && *token != "" && secret != nil && *secret != "" {
		claims := jwt.MapClaims{}
		token, err := verifyToken(token, claims, secret)
		if err != nil {
			return nil, nil, err
		}
		return token, claims, nil
	}
	return nil, nil, fmt.Errorf("empty token or secret")
}

// IsTokenExpired checks if a token has expired based on its claims
func IsTokenExpired(claims jwt.MapClaims) bool {
	if exp, ok := claims["exp"].(float64); ok {
		return float64(time.Now().Unix()) > exp // Check if current time is greater than expiration time
	}
	return true
}

// RespondAndAbort sends a response and aborts the request
func RespondAndAbort(c *gin.Context, message string, status int, data interface{}, errs []string) {
	util.Response(c, message, status, data, errs) // Send a response using util.Response
	c.Abort() // Abort the request
}
