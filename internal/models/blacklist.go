package models

import "time"

// Blacklist struct is used to represent a blacklisted token in the system.
type Blacklist struct {
	ID        uint      `gorm:"primaryKey"`      // Primary key for the blacklist entry
	Token     string    `gorm:"unique;not null"` // Blacklisted token string
	CreatedAt time.Time `gorm:"autoCreateTime"`  // Timestamp when the entry was created
	UpdatedAt time.Time `gorm:"autoUpdateTime"`  // Timestamp when the entry was last updated
}
