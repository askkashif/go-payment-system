package models

import "gorm.io/gorm"

// User struct represents a user in the system.
type User struct {
	gorm.Model
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	Password       string  `json:"password"`
	DateOfBirth    string  `json:"date_of_birth"`
	Email          string  `json:"email"`
	Phone          string  `json:"phone"`
	AccountNo      int     `json:"account_no"`
	InitialBalance float64 `json:"initial_balance"`
	Address        string  `json:"address"`
}

type Admin struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type AdminRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Transaction struct represents a financial transaction made by a user.
type Transaction struct {
	gorm.Model              // Embedding gorm.Model to include ID, CreatedAt, UpdatedAt, and DeletedAt fields
	Amount          float64 `json:"amount"`           // Amount of the transaction
	FromAccountNo   int     `json:"from_account_no"`  // Account number of the sender
	ToAccountNo     int     `json:"to_account_no"`    // Account number of the receiver
	Reference       string  `json:"reference"`        // Reference ID for the transaction
	TransactionType string  `json:"transaction_type"` // Type of the transaction (e.g., credit, debit)
}

// LoginRequest struct is used for handling login requests.
type LoginRequest struct {
	Email    string `json:"email"`    // Email address provided for login
	Password string `json:"password"` // Password provided for login
}

// DepositRequest struct is used for handling deposit requests.
type DepositRequest struct {
	Amount float64 `json:"amount"` // Amount to be deposited
}

// TransferRequest struct is used for handling transfer requests.
type TransferRequest struct {
	Amount    float64 `json:"amount"`     // Amount to be transferred
	AccountNo int     `json:"account_no"` // Account number of the receiver
}
