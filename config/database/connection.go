package database

import (
	"AntarJemput-Be-C/config"
	"AntarJemput-Be-C/models"
	"fmt"

	// "fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func InitDb() *gorm.DB {

	maxIdleConn := config.GetInt("DB_MAX_IDLE_CONNECTIONS")
	maxConn := config.GetInt("DB_MAX_CONNECTIONS")
	maxLifetimeConn := config.GetInt("DB_MAX_LIFETIME_CONNECTIONS")
	dbUsername := config.GetString("DB_USERNAME")
	dbPassword := config.GetString("DB_PASSWORD")
	dbHost := config.GetString("DB_HOST")
	dbPort := config.GetInt("DB_PORT")
	dbName := config.GetString("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
	dbUsername,dbPassword,dbHost,dbPort,dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetMaxOpenConns(maxConn)
	sqlDB.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

	InitCreateTable(db)
	log.Println("database connect success")
	log.Println("dsn :",dsn)
	return db

}

func InitCreateTable(db *gorm.DB) {
	db.AutoMigrate(
		&models.Province{},
		&models.Regency{},
		&models.District{},
		&models.Agents{},
		&models.Customer{},
		&models.Users{},
		&models.ServiceTransaction{},
		&models.Transactions{},
		&models.TypeServiceTransaction{},
		
	)

}
