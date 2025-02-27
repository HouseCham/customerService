package model

// Config struct is used to read the config.json file
type Config struct {
	DB  Database    `mapstructure:"database"`
	App Application `mapstructure:"app"`
}

// Database struct is used to read the database configuration
type Database struct {
	Host     string `mapstructure:"host"`
	Port     uint16 `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"dbname"`
}

// Application struct is used to read the application configuration
type Application struct {
	Host string `mapstructure:"host"`
	Port uint16 `mapstructure:"port"`
}
