package repository

import (
	"time"

	"gorm.io/gorm"
)

// Customer Struct corresponds to the Customers table in the database
type Customer struct {
	ID             uint      `gorm:"primaryKey"`
	FirstName      string    `gorm:"size:50;not null"`
	SecondName     string    `gorm:"size:50"`
	LastNameP      string    `gorm:"size:150;not null"`
	LastNameM      string    `gorm:"size:150;not null"`
	Phone          string    `gorm:"size:20;not null"`
	Email          string    `gorm:"size:150;unique;not null"`
	HashedPassword string    `gorm:"not null"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

// Indexes
func (Customer) TableName() string {
	return "customers"
}

func (Customer) CustomerFullNameIndex() string {
	return "customer_fullname"
}

func (Customer) EmailIndex() string {
	return "email"
}
