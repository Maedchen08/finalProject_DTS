package config

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/subosito/gotenv"
)

var appConfig *Config

type Config struct {
	AppName     string
	AppPort     string
	LogLevel    string
	Environment string
	DBUsername  string
	DBPassword  string
	DBHost      string
	DBPort      int
	DBName      string
}

func Init() *Config {
	defaultEnv := ".env"

	err := gotenv.Load(defaultEnv)
	if err != nil {
		log.Warning("Could not load environment")
	}

	log.SetOutput(os.Stdout)

	appConfig = &Config{
		AppName:     GetString("APP_NAME"),
		AppPort:     GetString("APP_PORT"),
		LogLevel:    GetString("LOG_LEVEL"),
		Environment: GetString("ENVIRONMENT"),
		DBUsername:  GetString("DB_USERNAME"),
		DBPassword:  GetString("DB_PASSWORD"),
		DBHost:      GetString("DB_HOST"),
		DBPort:      GetInt("DB_PORT", 3306),
		DBName:      GetString("DB_NAME"),
	}

	return appConfig

}
