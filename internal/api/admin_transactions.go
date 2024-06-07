package api

import (
	"github.com/gin-gonic/gin"
	"payment-system-four/internal/util"
)

// AdminTransactionsHandler retrieves all transactions and available balances for all customers
func (u *HTTPHandler) AdminTransactionsHandler(c *gin.Context) {
	// Retrieve all users from the database
	users, err := u.Repository.GetAllUsers()
	if err != nil {
		util.Response(c, "Error retrieving users", 500, "Internal server error", nil)
		return
	}

	var userTransactions []gin.H
	for _, user := range users {
		// Retrieve the user's transactions from the database
		transactions, err := u.Repository.GetUserTransactions(user.ID)
		if err != nil {
			util.Response(c, "Error retrieving transactions", 500, "Internal server error", nil)
			return
		}

		userTransactions = append(userTransactions, gin.H{
			"user_id":      user.ID,
			"balance":      user.InitialBalance,
			"transactions": transactions,
		})
	}

	util.Response(c, "Transactions retrieved", 200, gin.H{
		"user_transactions": userTransactions,
	}, nil)
}