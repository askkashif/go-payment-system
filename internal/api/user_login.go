package api

import (
	"github.com/gin-gonic/gin"
	"payment-system-one/internal/models"
	"payment-system-one/internal/util"
	"regexp"
)

// LoginUser is a handler function that handles user login
func (u *HTTPHandler) LoginUser(c *gin.Context) {
	var loginRequest *models.LoginRequest

	// Bind the request body to the loginRequest struct
	if err := c.ShouldBind(&loginRequest); err != nil {
		util.Response(c, "Invalid request", 400, "Bad request body", nil)
		return
	}

	// Validate the email format
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(loginRequest.Email) {
		util.Response(c, "Invalid email format", 400, "Invalid email format", nil)
		return
	}

	// Validate the password strength
	passwordRegex := regexp.MustCompile(`^(?=.*[!@#$%^&*(),.?":{}|<>])(?=.*\d).{8,}$`)
	if !passwordRegex.MatchString(loginRequest.Password) {
		util.Response(c, "Password must be at least 8 characters long and contain at least 2 special characters", 400, "Invalid password format", nil)
		return
	}

	// Find the user by email
	user, err := u.Repository.FindUserByEmail(loginRequest.Email)
	if err != nil {
		util.Response(c, "User not found", 404, "User not found", nil)
		return
	}

	// Check if the password matches
	if user.Password != loginRequest.Password {
		util.Response(c, "Invalid credentials", 401, "Invalid credentials", nil)
		return
	}

	// User login successful
	util.Response(c, "Login successful", 200, "Success", nil)
}