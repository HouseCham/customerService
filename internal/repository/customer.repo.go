package repository

import (
	"github.com/HouseCham/customerService/internal/log"
	"github.com/HouseCham/customerService/internal/model"
	"gorm.io/gorm"
)

// InsertNewCustomerDB is a repository function to insert a new Customer in the DB
func InsertNewCustomerDB(entry *model.Customer) (uint, error) {
	log.Logger.Infoln("Repository function invoked -> InsertNewCustomerDB")
	// goroutine for inserting new Customer in DB
	dbChanResponse := make(chan *gorm.DB)
	go func() {
		dbChanResponse <- DBContext.Create(&entry)
	}()
	// checking if there was an error inserting new Customer in DB
	if dbResult := <-dbChanResponse; dbResult.Error != nil {
		return 0, dbResult.Error
	}
	log.Logger.Infoln("Repository function finished -> InsertNewCustomerDB")
	// returning the ID of the new Customer
	return entry.ID, nil
}
// UpdateCustomerDB is a repository function to update a Customer in the DB
func UpdateCustomerDB(entry *model.Customer) error {
	log.Logger.Infoln("Repository function invoked -> UpdateCustomerDB")
	// goroutine for updating Customer in DB
	dbChanResponse := make(chan *gorm.DB)
	go func() {
		dbChanResponse <- DBContext.Save(&entry)
	}()
	// checking if there was an error updating Customer in DB
	if dbResult := <-dbChanResponse; dbResult.Error != nil {
		return dbResult.Error
	}
	log.Logger.Infoln("Repository function finished -> UpdateCustomerDB")
	return nil
}
// DeleteCustomerDB is a repository function to delete a Customer in the DB
func DeleteCustomerDB(entry *model.Customer) error {
	log.Logger.Infoln("Repository function invoked -> DeleteCustomerDB")
	// goroutine for deleting Customer in DB
	dbChanResponse := make(chan *gorm.DB)
	go func() {
		dbChanResponse <- DBContext.Delete(&entry)
	}()
	// checking if there was an error deleting Customer in DB
	if dbResult := <-dbChanResponse; dbResult.Error != nil {
		return dbResult.Error
	}
	log.Logger.Infoln("Repository function finished -> DeleteCustomerDB")
	return nil
}
// GetCustomerByID is a repository function to get a Customer by ID in the DB
func GetCustomerByID(id uint) (*model.Customer, error) {
	log.Logger.Infoln("Repository function invoked -> GetCustomerByID")
	// creating a new Customer model
	entry := &model.Customer{}
	// goroutine for getting Customer by ID in DB
	dbChanResponse := make(chan *gorm.DB)
	go func() {
		dbChanResponse <- DBContext.First(&entry, id)
	}()
	// checking if there was an error getting Customer by ID in DB
	if dbResult := <-dbChanResponse; dbResult.Error != nil {
		return nil, dbResult.Error
	}
	log.Logger.Infoln("Repository function finished -> GetCustomerByID")
	return entry, nil
}
// GetAllCustomers is a repository function to get all Customers in the DB
func GetAllCustomers() ([]model.Customer, error) {
	log.Logger.Infoln("Repository function invoked -> GetAllCustomers")
	// creating a new Customer model
	entries := []model.Customer{}
	// goroutine for getting all Customers in DB
	dbChanResponse := make(chan *gorm.DB)
	go func() {
		dbChanResponse <- DBContext.Find(&entries)
	}()
	// checking if there was an error getting all Customers in DB
	if dbResult := <-dbChanResponse; dbResult.Error != nil {
		return nil, dbResult.Error
	}
	log.Logger.Infoln("Repository function finished -> GetAllCustomers")
	return entries, nil
}
// GetCustomerByEmail is a repository function to get a Customer by email in the DB
func GetCustomerByEmail(email string) (*model.Customer, error) {
	log.Logger.Infoln("Repository function invoked -> GetCustomerByEmail")
	// creating a new Customer model
	entry := &model.Customer{}
	// goroutine for getting Customer by email in DB
	dbChanResponse := make(chan *gorm.DB)
	go func() {
		dbChanResponse <- DBContext.Where("email = ?", email).First(&entry)
	}()
	// checking if there was an error getting Customer by email in DB
	if dbResult := <-dbChanResponse; dbResult.Error != nil {
		return nil, dbResult.Error
	}
	log.Logger.Infoln("Repository function finished -> GetCustomerByEmail")
	return entry, nil
}