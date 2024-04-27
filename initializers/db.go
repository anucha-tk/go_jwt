package initializers

import (
	"fmt"
	"log"

	"github.com/anucha-tk/go_jwt/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(config *Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		config.DatabaseHost, config.DatabaseUsername, config.DatabasePassword, config.DatabaseName)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error can't connect database: %v", err)
	}
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Error auto migrate fail: %v", err)
	}
	log.Println("🚀 Connected Successfully to the Database")
}
