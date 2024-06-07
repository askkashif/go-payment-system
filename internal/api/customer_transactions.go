package api

import (
	"github.com/gin-gonic/gin"
	"payment-system-four/internal/util"
)

// CustomerTransactionsHandler retrieves all transactions and the available balance for a customer
func (u *HTTPHandler) CustomerTransactionsHandler(c *gin.Context) {
	// Get the user from the context
	user, err := u.GetUserFromContext(c)
	if err != nil {
		util.Response(c, "User not found", 404, "User not found", nil)
		return
	}

	// Retrieve the user's transactions from the database
	transactions, err := u.Repository.GetUserTransactions(user.ID)
	if err != nil {
		util.Response(c, "Error retrieving transactions", 500, "Internal server error", nil)
		return
	}

	util.Response(c, "Transactions retrieved", 200, gin.H{
		"balance":      user.InitialBalance,
		"transactions": transactions,
	}, nil)
}