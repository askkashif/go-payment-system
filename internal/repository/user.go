package repository

import (
	"payment-system-four/internal/models"
)

// FindUserByEmail retrieves a user from the database based on their email address.
func (p *Postgres) FindUserByEmail(email string) (*models.User, error) {
	// Create a new instance of the User model to hold the result of the query
	user := &models.User{}

	// Query the database to find a user with the provided email address
	if err := p.DB.Where("email = ?", email).First(&user).Error; err != nil {
		// If no user is found, return nil for the user and the error
		return nil, err
	}

	// If a user is found, return the user and nil for the error
	return user, nil
}

// FindAdminByEmail retrieves an admin from the database based on their email address.
func (p *Postgres) FindAdminByEmail(email string) (*models.Admin, error) {
	admin := &models.Admin{}

	if err := p.DB.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}

	return admin, nil
}

// FindUserByID retrieves a user from the database based on their ID.
func (p *Postgres) FindUserByID(userID uint) (*models.User, error) {
	// Create a new instance of the User model to hold the result of the query
	user := &models.User{}

	// Query the database to find a user with the provided ID
	if err := p.DB.First(&user, userID).Error; err != nil {
		// If no user is found, return nil for the user and the error
		return nil, err
	}

	// If a user is found, return the user and nil for the error
	return user, nil
}

// CreateUser creates a new user in the database.
func (p *Postgres) CreateUser(user *models.User) error {
	// Create a new user record in the database
	if err := p.DB.Create(user).Error; err != nil {
		// If an error occurs during creation, return the error
		return err
	}

	// Return nil if the creation is successful
	return nil
}

// UpdateUser updates an existing user in the database.
func (p *Postgres) UpdateUser(user *models.User) error {
	// Update the user record in the database
	if err := p.DB.Save(user).Error; err != nil {
		// If an error occurs during update, return the error
		return err
	}

	// Return nil if the update is successful
	return nil
}

// UpdateUserBalance updates the user's available balance in the database.
func (p *Postgres) UpdateUserBalance(userID uint, amount float64) error {
	// Find the user by ID
	user, err := p.FindUserByID(userID)
	if err != nil {
		return err
	}

	// Update the user's balance
	user.InitialBalance += amount

	// Save the updated user to the database
	if err := p.DB.Save(user).Error; err != nil {
		return err
	}

	return nil
}

func (p *Postgres) FindUserByAccountNo(accountNo int) (*models.User, error) {
	user := &models.User{}
	if err := p.DB.Where("account_no = ?", accountNo).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (p *Postgres) GetUserTransaction(account_no int) ([]models.Transaction, error) {
	var transaction []models.Transaction
	if err := p.DB.Where("from_account_no = ? OR to_account_no = ? ", account_no, account_no).Find(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func (p *Postgres) GetAllUsersTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := p.DB.Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
