package repository

import (
	"github.com/HouseCham/customerService/internal/log"
	"github.com/HouseCham/customerService/internal/model"
	"gorm.io/gorm"
)

// InsertNewCostumerDB is a repository function to insert a new costumer in the DB
func InsertNewCostumerDB(entry *model.Customer) (uint, error) {
	log.Logger.Infoln("Repository function invoked -> InsertNewCostumerDB")
	// goroutine for inserting new costumer in DB
	dbChanResponse := make(chan *gorm.DB)
	go func() {
		dbChanResponse <- DBContext.Create(&entry)
	}()
	// checking if there was an error inserting new costumer in DB
	if dbResult := <-dbChanResponse; dbResult.Error != nil {
		return 0, dbResult.Error
	}
	log.Logger.Infoln("Repository function finished -> InsertNewCostumerDB")
	// returning the ID of the new costumer
	return entry.ID, nil
}
// UpdateCostumerDB is a repository function to update a costumer in the DB
func UpdateCostumerDB(entry *model.Customer) error {
	log.Logger.Infoln("Repository function invoked -> UpdateCostumerDB")
	// goroutine for updating costumer in DB
	dbChanResponse := make(chan *gorm.DB)
	go func() {
		dbChanResponse <- DBContext.Save(&entry)
	}()
	// checking if there was an error updating costumer in DB
	if dbResult := <-dbChanResponse; dbResult.Error != nil {
		return dbResult.Error
	}
	log.Logger.Infoln("Repository function finished -> UpdateCostumerDB")
	return nil
}
// DeleteCostumerDB is a repository function to delete a costumer in the DB
func DeleteCostumerDB(entry *model.Customer) error {
	log.Logger.Infoln("Repository function invoked -> DeleteCostumerDB")
	// goroutine for deleting costumer in DB
	dbChanResponse := make(chan *gorm.DB)
	go func() {
		dbChanResponse <- DBContext.Delete(&entry)
	}()
	// checking if there was an error deleting costumer in DB
	if dbResult := <-dbChanResponse; dbResult.Error != nil {
		return dbResult.Error
	}
	log.Logger.Infoln("Repository function finished -> DeleteCostumerDB")
	return nil
}
// GetCostumerByID is a repository function to get a costumer by ID in the DB
func GetCostumerByID(id uint) (*model.Customer, error) {
	log.Logger.Infoln("Repository function invoked -> GetCostumerByID")
	// creating a new costumer model
	entry := &model.Customer{}
	// goroutine for getting costumer by ID in DB
	dbChanResponse := make(chan *gorm.DB)
	go func() {
		dbChanResponse <- DBContext.First(&entry, id)
	}()
	// checking if there was an error getting costumer by ID in DB
	if dbResult := <-dbChanResponse; dbResult.Error != nil {
		return nil, dbResult.Error
	}
	log.Logger.Infoln("Repository function finished -> GetCostumerByID")
	return entry, nil
}
// GetAllCostumers is a repository function to get all costumers in the DB
func GetAllCostumers() ([]model.Customer, error) {
	log.Logger.Infoln("Repository function invoked -> GetAllCostumers")
	// creating a new costumer model
	entries := []model.Customer{}
	// goroutine for getting all costumers in DB
	dbChanResponse := make(chan *gorm.DB)
	go func() {
		dbChanResponse <- DBContext.Find(&entries)
	}()
	// checking if there was an error getting all costumers in DB
	if dbResult := <-dbChanResponse; dbResult.Error != nil {
		return nil, dbResult.Error
	}
	log.Logger.Infoln("Repository function finished -> GetAllCostumers")
	return entries, nil
}
// GetCostumerByEmail is a repository function to get a costumer by email in the DB
func GetCostumerByEmail(email string) (*model.Customer, error) {
	log.Logger.Infoln("Repository function invoked -> GetCostumerByEmail")
	// creating a new costumer model
	entry := &model.Customer{}
	// goroutine for getting costumer by email in DB
	dbChanResponse := make(chan *gorm.DB)
	go func() {
		dbChanResponse <- DBContext.Where("email = ?", email).First(&entry)
	}()
	// checking if there was an error getting costumer by email in DB
	if dbResult := <-dbChanResponse; dbResult.Error != nil {
		return nil, dbResult.Error
	}
	log.Logger.Infoln("Repository function finished -> GetCostumerByEmail")
	return entry, nil
}