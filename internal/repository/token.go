package repository

import "payment-system-one/internal/models"

// TokenInBlacklist checks if a token is already in the blacklist collection.
func (p *Postgres) TokenInBlacklist(token *string) bool {
	// Create a new instance of the Blacklist model
	tok := &models.Blacklist{}

	// Query the database to find a record with the provided token
	if err := p.DB.Where("token = ?", token).First(&tok).Error; err != nil {
		// If no record is found, return false
		return false
	}

	// If a record is found, return true
	return true
}
