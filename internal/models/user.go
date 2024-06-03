package models

import "gorm.io/gorm"

// User struct represents a user in the system.
type User struct {
	gorm.Model                  // Embedding gorm.Model to include ID, CreatedAt, UpdatedAt, and DeletedAt fields
	FirstName    string `json:"first_name"`    // User's first name
	LastName     string `json:"last_name"`     // User's last name
	Password     string `json:"password"`      // User's password (should be hashed in a real application)
	DateOfBirth  string `json:"date_of_birth"` // User's date of birth
	Email        string `json:"email"`         // User's email address
	Phone        string `json:"phone"`         // User's phone number
	Address      string `json:"address"`       // User's physical address
	BankAccNum   string `json:"bank_acc_num"`  // User's Bank Account number
	SortCode     string `json:"sort_code"`     // User's physical address
	LoginCounter int    `json:"login_counter"` // Counter for login attempts (used for security measures)
	IsLocked     bool   `json:"is_locked"`     // Boolean to indicate if the user's account is locked
}

// Uncomment and expand the UserProfile struct if additional user profile details are needed.
// type UserProfile struct {
// 	gorm.Model
// 	ValidIdentity string `json:"valid_identity"` // Valid identity proof for the user
// 	PassPort      string `json:"passport"`       // User's passport information
// }

// Transaction struct represents a financial transaction made by a user.
type Transaction struct {
	gorm.Model                         // Embedding gorm.Model to include ID, CreatedAt, UpdatedAt, and DeletedAt fields
	UserID          uint    `json:"user_id"`          // Foreign key referencing the user who made the transaction
	Amount          float64 `json:"amount"`           // Amount of the transaction
	Reference       string  `json:"reference"`        // Reference ID for the transaction
	TransactionType string  `json:"transaction_type"` // Type of the transaction (e.g., credit, debit)
}

// LoginRequest struct is used for handling login requests.
type LoginRequest struct {
	Email    string `json:"email"`     // Email address provided for login
	Password string `json:"password"`  // Password provided for login
}
