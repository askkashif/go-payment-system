package repository

import "payment-system-four/internal/models"

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
