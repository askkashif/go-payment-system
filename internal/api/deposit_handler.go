package api

import (
	"github.com/gin-gonic/gin"
	"payment-system-four/internal/models"
	"payment-system-four/internal/util"
)

// DepositHandler handles the deposit of funds into a user's account
func (u *HTTPHandler) DepositHandler(c *gin.Context) {
	var depositRequest struct {
		Amount float64 `json:"amount"`
	}

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
		UserID:          user.ID,
		Amount:          depositRequest.Amount,
		TransactionType: "deposit",
	}

	// Save the transaction to the database
	if err := u.Repository.CreateTransaction(transaction); err != nil {
		util.Response(c, "Error creating transaction", 500, "Internal server error", nil)
		return
	}

	// Update the user's balance
	user.InitialBalance += depositRequest.Amount

	// Save the updated user to the database
	if err := u.Repository.UpdateUser(user); err != nil {
		util.Response(c, "Error updating balance", 500, "Internal server error", nil)
		return
	}

	util.Response(c, "Deposit successful", 200, gin.H{
		"new_balance": user.InitialBalance,
	}, nil)
}