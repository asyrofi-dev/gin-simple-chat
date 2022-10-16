package config

import (
	"fmt"
	"gin-simple-chat/entity"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/joho/godotenv/autoload"
)

var DB *gorm.DB
var err error

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"), os.Getenv("DB_SSL_MODE"), os.Getenv("DB_TIMEZONE"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("Database Connection Error :", err)
		return
	}

	log.Println("Database Connection Succeed")
}

func DropAllTables() {
	err = DB.Debug().Migrator().DropTable(&entity.User{})

	if err != nil {
		log.Println("Drop Table Error :", err)
		return
	}

	log.Println("Drop Table Succeed")
}

func MigrateDB() {
	err = DB.Debug().Migrator().AutoMigrate(&entity.User{})

	if err != nil {
		log.Println("Database Migration Error :", err)
		return
	}

	log.Println("Database Migration Succeed")
}
