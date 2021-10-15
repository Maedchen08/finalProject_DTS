package database

import (
	"AntarJemput-Be-C/config"
	"AntarJemput-Be-C/models"
	"fmt"
<<<<<<< HEAD

	// "fmt"
=======
>>>>>>> 4e04e7ccabce770f32a4c5c36ac73dee2ed0ca21
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

<<<<<<< HEAD

func InitDb(appConfig *config.Config) *gorm.DB {

	maxIdleConn := config.GetInt("DB_MAX_IDLE_CONNECTIONS")
	maxConn := config.GetInt("DB_MAX_CONNECTIONS")
	maxLifetimeConn := config.GetInt("DB_MAX_LIFETIME_CONNECTIONS")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
	appConfig.DBUsername,appConfig.DBPassword,appConfig.DBHost,appConfig.DBPort,appConfig.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
=======
func InitDb(appConfig *config.Config) *gorm.DB {
	maxIdleConn := config.GetInt("DB_MAX_IDLE_CONNECTIONS", 10)
	maxConn := config.GetInt("DB_MAX_CONNECTIONS", 100)
	maxLifetimeConn := config.GetInt("DB_MAX_LIFETIME_CONNECTIONS", 2)

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		appConfig.DBUsername,
		appConfig.DBPassword,
		appConfig.DBHost,
		appConfig.DBPort,
		appConfig.DBName,
	)
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
>>>>>>> 4e04e7ccabce770f32a4c5c36ac73dee2ed0ca21

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
<<<<<<< HEAD
	log.Println("database connect success")
	log.Println("dsn :",dsn)
=======
	log.Println("database connect success!!!")
>>>>>>> 4e04e7ccabce770f32a4c5c36ac73dee2ed0ca21
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
