package repository

import (
	"log"
	"payment-system-four/internal/ports"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

// NewDB create/returns a new instance of our Database
func NewDB(DB *gorm.DB) ports.Repository {
	return &Postgres{
		DB: DB,
	}
}

// Initialize opens the database, create tables if not created and populate it if its empty and returns a DB
func Initialize(dbURI string) (*gorm.DB, error) {

	conn, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		//	log.Fatal(err)
	}
	err = conn.AutoMigrate()
	if err != nil {
		return nil, err
	}
	log.Println("Database connection successful")
	return conn, nil
}
