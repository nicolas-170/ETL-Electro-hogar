package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	App      App
	Database Database
}

type App struct {	
	Timeout *time.Duration
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func LoadConfig() *Config {
	godotenv.Load(os.Getenv("ENV_FILE"))
	return &Config{
		App: App{
			Timeout: GetDurationMilliEnv("TIMEOUT", 20000),
		},
		Database: Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
	}
}

func GetDurationMilliEnv(envName string, defaultValue int) *time.Duration {
	if val, err := strconv.Atoi(os.Getenv(envName)); err == nil && val > 0 {
		defaultValue = val
	}
	dur := time.Duration(defaultValue) * time.Millisecond
	return &dur
}
