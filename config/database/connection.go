package database
import(
	"log"
	"time"
	"AntarJemput-Be-C/models"
	"AntarJemput-Be-C/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	maxIdleConn := config.GetInt("DB_MAX_IDLE_CONNECTIONS")
	maxConn := config.GetInt("DB_MAX_CONNECTIONS")
	maxLifetimeConn := config.GetInt("DB_MAX_LIFETIME_CONNECTIONS")

	dsn := config.GetString("DB_SERVER_URL")
	db,err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

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
	return db
	
}

func InitCreateTable(db *gorm.DB){
	
	db.AutoMigrate(&models.Regency{})
	db.AutoMigrate(&models.Province{})
	db.AutoMigrate(&models.District{})
	// db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.TypeTransaction{})
	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.TypeTransaction{})


	
	



}

