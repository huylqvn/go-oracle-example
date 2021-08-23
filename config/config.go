package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Port   string
	DBPort string
	DBHost string
	DBUser string
	DBPass string
	DBName string
}

var config Config

func NewFromFile() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file", err)
		return nil, err
	}
	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("unable to decode into struct", err)
	}
	return &config, nil
}

func New() (*Config, error) {
	godotenv.Load()
	config.Port = os.Getenv("PORT")
	config.DBHost = os.Getenv("DB_HOST")
	config.DBPort = os.Getenv("DB_PORT")
	config.DBUser = os.Getenv("DB_USER")
	config.DBPass = os.Getenv("DB_PASS")
	config.DBName = os.Getenv("DB_NAME")

	return &config, nil
}
