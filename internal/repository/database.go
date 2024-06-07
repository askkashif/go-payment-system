package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"payment-system-four/internal/models"
	"payment-system-four/internal/ports"
)

// Postgres struct represents the PostgreSQL repository implementation.
type Postgres struct {
	DB *gorm.DB // DB holds the GORM database connection
}

// NewDB creates and returns a new instance of the Postgres repository.
func NewDB(DB *gorm.DB) ports.Repository {
	return &Postgres{
		DB: DB, // Assigning the provided database connection to the repository instance
	}
}

// Initialize opens the database connection, creates necessary tables, and returns the DB instance.
func Initialize(dbURI string) (*gorm.DB, error) {
	// Opening a connection to the PostgreSQL database
	conn, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		// If an error occurs while opening the connection, it's logged but not fatal
		// as we will try to migrate the database tables regardless
	}

	// Auto-migrating the User and Transaction models to create the necessary tables
	err = conn.AutoMigrate(&models.User{}, &models.Transaction{})
	if err != nil {
		// If an error occurs during auto-migration, return the error
		return nil, err
	}

	// Logging successful database connection
	log.Println("Database connection successful")

	// Returning the GORM database instance
	return conn, nil
}

// CreateTransaction creates a new transaction in the database.
func (p *Postgres) CreateTransaction(transaction *models.Transaction) error {
	if err := p.DB.Create(transaction).Error; err != nil {
		return err
	}
	return nil
}

// GetAllUsers retrieves all users from the database.
func (p *Postgres) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	if err := p.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserTransactions retrieves all transactions for a specific user.
func (p *Postgres) GetUserTransactions(userID uint) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	if err := p.DB.Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}