package routes

import (
	"github.com/anucha-tk/go_jwt/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(router fiber.Router) {
	router.Get("/healthchecker", controllers.HealthChecker)
	router.Post("/register", controllers.Register)
	router.Post("/login", controllers.Login)
	router.Get("/logout", controllers.Logout)
	router.All("*", controllers.Notfound)
}
