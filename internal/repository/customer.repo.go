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