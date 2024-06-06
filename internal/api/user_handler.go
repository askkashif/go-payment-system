package api

import (
	"github.com/gin-gonic/gin"
	"payment-system-one/internal/models"
	"payment-system-one/internal/util"
	"regexp"
	"math/rand"
	"time"
)

// RegisterUser is a handler function that registers a new user
func (u *HTTPHandler) RegisterUser(c *gin.Context) {
	var newUser *models.User

	// Bind the request body to the newUser struct
	if err := c.ShouldBind(&newUser); err != nil {
		util.Response(c, "Invalid request", 400, "Bad request body", nil)
		return
	}

	// Validate the email format
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(newUser.Email) {
		util.Response(c, "Invalid email format", 400, "Invalid email format", nil)
		return
	}

	// Validate the password strength
	passwordRegex := regexp.MustCompile(`^(?=.*[!@#$%^&*(),.?":{}|<>])(?=.*\d).{8,}$`)
	if !passwordRegex.MatchString(newUser.Password) {
		util.Response(c, "Password must be at least 8 characters long and contain at least 2 special characters", 400, "Invalid password format", nil)
		return
	}

	// Hash the password before storing it
	hashedPassword, err := util.HashPassword(newUser.Password)
	if err != nil {
		util.Response(c, "Error hashing password", 500, "Internal server error", nil)
		return
	}
	newUser.Password = hashedPassword

	// Generate an account number
	rand.Seed(time.Now().UnixNano())
	accountNumber := rand.Intn(999999999-100000000) + 100000000 // Generate a random 9-digit account number

	newUser.AccountNumber = accountNumber

	// Persist the new user in the database
	err = u.Repository.CreateUser(newUser)
	if err != nil {
		util.Response(c, "Error creating user account", 500, "Internal server error", nil)
		return
	}

	// Send an email to the administrator for approval
	// ... (code to send email to administrator)

	util.Response(c, "User registration successful. Your account is pending approval.", 200, gin.H{
		"message":       "Success",
		"account_number": accountNumber,
		"initial_balance": newUser.InitialBalance, // Include the initial balance in the response
	}, nil)
}
