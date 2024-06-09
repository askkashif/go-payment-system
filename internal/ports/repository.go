package ports

import "payment-system-four/internal/models"

// Repository defines the interface for interacting with data storage.
type Repository interface {
	// FindUserByEmail retrieves a user by their email address.
	FindUserByEmail(email string) (*models.User, error)

	// FindAdminByEmail retrieves an admin by their email address.
	FindAdminByEmail(email string) (*models.Admin, error)

	// TokenInBlacklist checks if a token is in the blacklist.
	TokenInBlacklist(token *string) bool

	// CreateUser creates a new user in the database.
	CreateUser(user *models.User) error

	// UpdateUser updates an existing user in the database.
	UpdateUser(user *models.User) error

	// UpdateUserBalance updates the user's available balance in the database.
	UpdateUserBalance(userID uint, amount float64) error

	// GetAllUsers retrieves all users from the database.
	GetAllUsers() ([]*models.User, error)

	// GetUserTransactions retrieves all transactions for a specific user.
	GetUserTransaction(account_no int) ([]models.Transaction, error)

	// CreateTransaction creates a new transaction in the database.
	CreateTransaction(transaction *models.Transaction) error

	FindUserByAccountNo(accountNo int) (*models.User, error)

	// GetAllUsersTransactions retrieves all transactions for all users.
	GetAllUsersTransactions() ([]models.Transaction, error)
}
