package api

import (
	"github.com/gin-gonic/gin"
	"payment-system-four/internal/models"
	"payment-system-four/internal/util"
	"regexp"
)

// CreateAdminAccount is a handler function that creates an administrator account
func (u *HTTPHandler) CreateAdminAccount(c *gin.Context) {
	var adminUser *models.User

	// Bind the request body to the adminUser struct
	if err := c.ShouldBind(&adminUser); err != nil {
		util.Response(c, "Invalid request", 400, "Bad request body", nil)
		return
	}

	// Validate the email format
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(adminUser.Email) {
		util.Response(c, "Invalid email format", 400, "Invalid email format", nil)
		return
	}

	// Validate the password strength
	passwordRegex := regexp.MustCompile(`^(?=.*[!@#$%^&*(),.?":{}|<>])(?=.*\d).{8,}$`)
	if !passwordRegex.MatchString(adminUser.Password) {
		util.Response(c, "Password must be at least 8 characters long and contain at least 2 special characters", 400, "Invalid password format", nil)
		return
	}

	// Hash the password before storing it
	hashedPassword, err := util.HashPassword(adminUser.Password)
	if err != nil {
		util.Response(c, "Error hashing password", 500, "Internal server error", nil)
		return
	}
	adminUser.Password = hashedPassword

	// Persist the administrator user in the database
	err = u.Repository.CreateUser(adminUser)
	if err != nil {
		util.Response(c, "Error creating administrator account", 500, "Internal server error", nil)
		return
	}

	util.Response(c, "Administrator account created successfully", 200, "Success", nil)
}
