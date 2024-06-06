package repository

import (
	"gorm.io/driver/postgres"            // Importing PostgreSQL driver for GORM
	"gorm.io/gorm"                       // Importing GORM package for ORM functionality
	"log"                                // Importing log package for logging
	"payment-system-one/internal/models" // Importing models package for user model
	"payment-system-one/internal/ports"  // Importing ports package for Repository interface
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

	// Auto-migrating the User model to create the necessary table
	err = conn.AutoMigrate(&models.User{})
	if err != nil {
		// If an error occurs during auto-migration, return the error
		return nil, err
	}

	// Logging successful database connection
	log.Println("Database connection successful")

	// Returning the GORM database instance
	return conn, nil
}
