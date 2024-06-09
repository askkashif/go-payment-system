package api

import (
	"math/rand"
	"net/http"
	"os"
	"payment-system-four/internal/middleware"
	"payment-system-four/internal/models"
	"payment-system-four/internal/util"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

	// check if user already exists
	_, err := u.Repository.FindUserByEmail(newUser.Email)
	if err == nil {
		util.Response(c, "user does exist", 404, "user already exists", nil)
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

	newUser.AccountNo = accountNumber

	// Persist the new user in the database
	err = u.Repository.CreateUser(newUser)
	if err != nil {
		util.Response(c, "Error creating user account", 500, "Internal server error", nil)
		return
	}

	util.Response(c, "User registration successful.", 200, gin.H{
		"message":         "Success",
		"account_number":  accountNumber,
		"initial_balance": newUser.InitialBalance, // Include the initial balance in the response
	}, nil)
}

// LoginUser is a handler function that handles user login
func (u *HTTPHandler) LoginUser(c *gin.Context) {
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
	user, err := u.Repository.FindUserByEmail(loginRequest.Email)
	if err != nil {
		util.Response(c, "user does not exist", 404, "user not found", nil)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		util.Response(c, "invalid password", 400, "invalid request", nil)
		return
	}

	//Generate token
	accessClaims, refreshClaims := middleware.GenerateClaims(user.Email)

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
		"user":          user,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil)
}

// Current Balance
func (u *HTTPHandler) BalanceCheck(c *gin.Context) {

	// get user from context
	user, err := u.GetUserFromContext(c)
	if err != nil {
		util.Response(c, "user not fount", 500, "user not found", nil)
		return
	}
	// checking balance
	util.Response(c, "Balance retrieved successfully", 200, "sucess", nil)
	c.IndentedJSON(200, gin.H{"balance": user.InitialBalance})

}

// TransferHandler handles the transfer of funds from one user to another
func (u *HTTPHandler) UserTransferHandler(c *gin.Context) {
	var transferRequest *models.TransferRequest

	// Bind the request body to the transferRequest struct
	if err := c.ShouldBindJSON(&transferRequest); err != nil {
		util.Response(c, "Invalid request", 400, "Bad request body", nil)
		return
	}

	// Get the sender user from the context
	senderUser, err := u.GetUserFromContext(c)
	if err != nil {
		util.Response(c, "Sender user not found", 404, "Sender user not found", nil)
		return
	}

	// Find the receiver user by email
	receiverUser, err := u.Repository.FindUserByAccountNo(transferRequest.AccountNo)
	if err != nil {
		util.Response(c, "Receiver user not found", 404, "Receiver user not found", nil)
		return
	}

	// Check if the sender has sufficient balance
	if senderUser.InitialBalance < transferRequest.Amount {
		util.Response(c, "Insufficient balance", 400, "Insufficient balance", nil)
		return
	}

	// Update the sender's balance
	if err := u.Repository.UpdateUserBalance(senderUser.ID, -transferRequest.Amount); err != nil {
		util.Response(c, "Error updating sender balance", 500, "Internal server error", nil)
		return
	}

	// Update the receiver's balance
	if err := u.Repository.UpdateUserBalance(receiverUser.ID, transferRequest.Amount); err != nil {
		util.Response(c, "Error updating receiver balance", 500, "Internal server error", nil)
		return
	}

	// Create a new transaction for the transfer
	transaction := &models.Transaction{
		FromAccountNo:   senderUser.AccountNo,
		ToAccountNo:     receiverUser.AccountNo,
		Amount:          transferRequest.Amount,
		TransactionType: "transfer",
	}

	// Save the transaction to the database
	if err := u.Repository.CreateTransaction(transaction); err != nil {
		util.Response(c, "Error creating transaction", 500, "Internal server error", nil)
		return
	}

	util.Response(c, "Funds transfered successfully", 200, nil, nil)
}

// DepositHandler handles the deposit of funds into a user's account
func (u *HTTPHandler) DepositHandler(c *gin.Context) {
	var depositRequest *models.DepositRequest

	// Bind the request body to the depositRequest struct
	if err := c.ShouldBindJSON(&depositRequest); err != nil {
		util.Response(c, "Invalid request", 400, "Bad request body", nil)
		return
	}

	// Get the user from the context
	user, err := u.GetUserFromContext(c)
	if err != nil {
		util.Response(c, "User not found", 404, "User not found", nil)
		return
	}

	// Create a new transaction for the deposit
	transaction := &models.Transaction{
		FromAccountNo:   user.AccountNo,
		ToAccountNo:     user.AccountNo,
		Amount:          depositRequest.Amount,
		TransactionType: "deposit",
	}

	// Save the transaction to the database
	if err := u.Repository.CreateTransaction(transaction); err != nil {
		util.Response(c, "Error creating transaction", 500, "Internal server error", nil)
		return
	}

	// Update the user's balance
	if err := u.Repository.UpdateUserBalance(user.ID, depositRequest.Amount); err != nil {
		util.Response(c, "Error updating balance", 500, "Internal server error", nil)
		return
	}

	util.Response(c, "Deposit successful", 200, gin.H{
		"new_balance": user.InitialBalance,
	}, nil)
}

// Transaction history
func (u *HTTPHandler) UserTransactionHistory(c *gin.Context) {
	// get user from context
	user, err := u.GetUserFromContext(c)
	if err != nil {
		util.Response(c, "user not fount", 500, "user not found", nil)
		return
	}
	// Add/ create a  function to the repository  to retrive the transacton history
	transaction, err := u.Repository.GetUserTransaction(user.AccountNo)
	if err != nil {
		util.Response(c, "count not retrieve trasaction", 500, "not retrieved", nil)
		return
	}
	util.Response(c, "transaction successfully retrieved", 200, "successful", nil)
	c.IndentedJSON(200, gin.H{"Transaction History": transaction})
}
