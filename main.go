package main

import (
	"log"
	"os"

	"github.com/anucha-tk/go_jwt/database"
	"github.com/anucha-tk/go_jwt/routes"
	"github.com/anucha-tk/go_jwt/utils"
	"github.com/gofiber/fiber/v3"
)

func main() {
	utils.LoadENV()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT not found")
	}
	database.Connect()

	app := fiber.New()
	routes.Setup(app)
	err := app.Listen(":" + port)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
