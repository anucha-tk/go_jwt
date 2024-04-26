package database

import (
	"fmt"
	"log"
	"os"

	"github.com/anucha-tk/go_jwt/models"
	"github.com/anucha-tk/go_jwt/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	utils.LoadENV()
	dbHost := os.Getenv("DATABASE_HOST")
	dbUserName := os.Getenv("DATABASE_USERNAME")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		dbHost, dbUserName, dbPassword, dbName)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error can't connect database: %v", err)
	}
	DB = conn
	err = conn.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Error auto migrate fail: %v", err)
	}
	log.Print("Database connect successful")
}
