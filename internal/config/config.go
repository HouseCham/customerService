package config

import (
	"fmt"

	"github.com/HouseCham/customerService/internal/model"
	"github.com/spf13/viper"
)

var ConfigFile model.Config

// getConfigFile reads the config.json file and returns a Config struct
func GetConfigFile() error {
	v := viper.New()
	v.SetConfigFile("./../config.json")

	// Read the config file
	err := v.ReadInConfig()
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	// Unmarshal the config file into a struct
	var configFile model.Config
	err = v.Unmarshal(&configFile)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	ConfigFile = configFile
	return nil
}
