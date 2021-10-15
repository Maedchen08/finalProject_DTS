package config

import (
	"github.com/subosito/gotenv"
	"log"
	"os"
)

var appConfig *Config

type Config struct {
	AppName        string
	AppPort        string
	LogLevel       string
	Environment    string
	JWTSecret      string
	RedisAddress   string
	DBUsername     string
	DBPassword     string
	DBHost         string
	DBPort         int
	DBName         string
	MinioEndpoint  string
	MinioAccessKey string
	MinioSecretKey string
	MinioRegion    string
	MinioBucket    string
	

}

func Init() *Config {
	defaultEnv := ".env"

	err := gotenv.Load(defaultEnv)
	if err != nil {
		log.Fatalf("Could not load environment")
	}

	log.SetOutput(os.Stdout)

	appConfig = &Config{
		AppName:        GetString("APP_NAME"),
		AppPort:        GetString("APP_PORT"),
		LogLevel:       GetString("LOG_LEVEL"),
		Environment:    GetString("ENVIRONMENT"),
		JWTSecret:      GetString("JWT_SECRET"),
		RedisAddress:   GetString("REDIS_ADDRESS"),
		DBUsername:     GetString("DB_USERNAME"),
		DBPassword:     GetString("DB_PASSWORD"),
		DBHost:         GetString("DB_HOST"),
		DBPort:         GetInt("DB_PORT"),
		DBName:         GetString("DB_NAME"),
		MinioEndpoint:  GetString("MINIO_ENDPOINT"),
		MinioAccessKey: GetString("MINIO_ACCESS_KEY"),
		MinioSecretKey: GetString("MINIO_SECRET_KEY"),
		MinioRegion:    GetString("MINIO_REGION"),
		MinioBucket:    GetString("MINIO_BUCKET"),
	}


	return appConfig

}
