package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	JWTSecret  string
	Addr       string
	Port       string
	AccessTTL  time.Duration
	RefreshTTL time.Duration
}

var ENV *Config

func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("No .env file found, reading environment variables")
	}

	accessTTL, err := time.ParseDuration(viper.GetString("ACCESS_TTL"))
	if err != nil {
		accessTTL = 24 * time.Hour
	}

	refreshTTL, err := time.ParseDuration(viper.GetString("REFRESH_TTL"))
	if err != nil {
		refreshTTL = 7 * 24 * time.Hour
	}

	ENV = &Config{
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetString("DB_PORT"),
		DBName:     viper.GetString("DB_NAME"),
		JWTSecret:  viper.GetString("JWT_SECRET"),
		Addr:       viper.GetString("ADDR"),
		Port:       viper.GetString("PORT"),
		AccessTTL:  accessTTL,
		RefreshTTL: refreshTTL,
	}

	return ENV
}
