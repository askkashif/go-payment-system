package api

import (
	"net/http"
	"os"
	"payment-system-four/internal/middleware"
	"payment-system-four/internal/models"
	"payment-system-four/internal/util"
	"regexp"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

	// Check if the user already exists
	_, err := u.Repository.FindAdminByEmail(adminUser.Email)
	if err == nil {
		util.Response(c, "user does exist", 404, "user already exists", nil)
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

// AdminLogin is a handler function that handles administrator login
func (u *HTTPHandler) AdminLogin(c *gin.Context) {
	var loginRequest *models.LoginRequest
	if err := c.ShouldBind(&loginRequest); err != nil {
		util.Response(c, "invalid request", 400, "bad request body", nil)
		return
	}

	if loginRequest.Email == "" || loginRequest.Password == "" {
		util.Response(c, "Please enter your email or password", 400, "bad request body", nil)
		return
	}

	// check if user already exists
	admin, err := u.Repository.FindAdminByEmail(loginRequest.Email)
	if err != nil {
		util.Response(c, "admin does not exist", 404, "admin not found", nil)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(loginRequest.Password)); err != nil {
		util.Response(c, "invalid password", 400, "invalid request", nil)
	}

	//Generate token
	accessClaims, refreshClaims := middleware.GenerateClaims(admin.Email)

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
	c.Header("access_token", *accessToken)
	c.Header("refresh_token", *refreshToken)

	util.Response(c, "login successful", http.StatusOK, gin.H{
		"admin":         admin,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil)
}

// AdminTransactionsHandler retrieves all transactions and available balances for all customers
func (u *HTTPHandler) AdminTransactionsHandler(c *gin.Context) {
	// get admin from context
	_, err := u.GetAdminFromContext(c)
	if err != nil {
		util.Response(c, "admin not found", 500, "admin not found", nil)
		return
	}
	// Add/ create a function to the repository to retrieve all transactions
	transactions, err := u.Repository.GetAllUsersTransactions()
	if err != nil {
		util.Response(c, "could not retrieve transactions", 500, "not retrieved", nil)
		return
	}
	util.Response(c, "transactions successfully retrieved", 200, "successful", nil)
	c.IndentedJSON(200, gin.H{"Transactions": transactions})
}
