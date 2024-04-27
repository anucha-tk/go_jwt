package main

import (
	"log"

	"github.com/anucha-tk/go_jwt/initializers"
	"github.com/anucha-tk/go_jwt/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	initializers.Connect(&config)
}

func main() {
	app := fiber.New()
	v1 := app.Group("/api/v1")
	app.Use(logger.New())
	routes.Setup(v1)
	err := app.Listen("localhost:" + "5000")
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
