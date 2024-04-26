package repository

import (
	"fmt"

	"github.com/HouseCham/customerService/internal/config"
	"github.com/HouseCham/customerService/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBContext *gorm.DB

// SetDBConnection is a function to set the database connection with the given connection string
func SetUpDBConnection() error {
	db, err := gorm.Open(postgres.Open(config.ConfigFile.DB.ConnectionString), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database: %v", err)
	}
	// Migrate the schema
	db.AutoMigrate(&model.Customer{})
	DBContext = db
	return nil
}