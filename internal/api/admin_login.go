// payment-system-four/internal/api/admin_login.go

package api

import (
	"github.com/gin-gonic/gin"
	"payment-system-one/internal/models"
	"payment-system-one/internal/util"
)

// AdminLogin is a handler function that handles administrator login
func (u *HTTPHandler) AdminLogin(c *gin.Context) {
	var loginRequest *models.LoginRequest

	// Bind the request body to the loginRequest struct
	if err := c.ShouldBind(&loginRequest); err != nil {
		util.Response(c, "Invalid request", 400, "Bad request body", nil)
		return
	}

	// Find the administrator by email
	admin, err := u.Repository.FindUserByEmail(loginRequest.Email)
	if err != nil {
		util.Response(c, "Administrator not found", 404, "Administrator not found", nil)
		return
	}

	// Check if the password matches
	if admin.Password != loginRequest.Password {
		util.Response(c, "Invalid credentials", 401, "Invalid credentials", nil)
		return
	}

	// Administrator login successful
	util.Response(c, "Login successful", 200, "Success", nil)
}
