package routes

import (
	"github.com/anucha-tk/go_jwt/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Hello)
	app.Post("/register", controllers.Register)
}
