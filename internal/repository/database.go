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
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		config.ConfigFile.DB.User, 
		config.ConfigFile.DB.Password, 
		config.ConfigFile.DB.Host, 
		config.ConfigFile.DB.Port,
		config.ConfigFile.DB.DbName, 
	)
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database: %v", err)
	}
	// Migrate the schema
	db.AutoMigrate(&model.Customer{})
	DBContext = db
	return nil
}