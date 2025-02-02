package config

import (
	"os"
	"strconv"

	"github.com/HouseCham/customerService/internal/model"
)

var ConfigFile model.Config

// getConfigFile reads the config.json file and returns a Config struct
func GetConfigFile() error {
	// Read the environment variables
	appPort, err := strconv.ParseUint(os.Getenv("APP_PORT"), 10, 16)
	if err != nil {
		return err 
	}
	dbPort, err := strconv.ParseUint(os.Getenv("DB_PORT"), 10, 16)
	if err != nil {
		return err
	}

	// Unmarshal the config file into a struct
	var configFile model.Config = model.Config{
		App: model.Application{
			Host: os.Getenv("APP_HOST"),
			Port: uint16(appPort),
		},
		DB: model.Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     uint16(dbPort),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DbName:   os.Getenv("DB_NAME"),
		},
	}

	ConfigFile = configFile
	return nil
}
