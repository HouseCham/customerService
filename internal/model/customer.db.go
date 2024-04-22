package model

import (
	"time"

	"gorm.io/gorm"
)

// Customer Struct corresponds to the Customers table in the database
type Customer struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	FirstName      string    `gorm:"size:50;not null" json:"firstName" validate:"required,lettersAccents,max=50"`
	SecondName     string    `gorm:"size:50" json:"secondName" validate:"lettersAccents,max=50"`
	LastNameP      string    `gorm:"size:150;not null" json:"lastNameP" validate:"required,lettersAccents,max=150"`
	LastNameM      string    `gorm:"size:150;not null" json:"lastNameM" validate:"required,lettersAccents,max=150"`
	PhoneNumber    string    `gorm:"size:20;not null" json:"phoneNumber" validate:"required,phone,max=20"`
	Email          string    `gorm:"size:150;unique;not null" json:"email" validate:"required,email,max=150"`
	Password       string    `gorm:"-" json:"passwordHash" validate:"required"`
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
