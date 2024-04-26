package controllers

import (
	"log"

	"github.com/anucha-tk/go_jwt/common"
	"github.com/anucha-tk/go_jwt/database"
	"github.com/anucha-tk/go_jwt/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("ready")
}

type User struct {
	Name     string `json:"name" validate:"required,min=3,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Passowrd string `json:"password" validate:"required"`
	Age      int    `json:"age"`
}

func Register(c *fiber.Ctx) error {
	body := new(User)
	err := c.BodyParser(body)
	if err != nil {
		return err
	}
	// Validate request body
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return common.ResponseError(c, 422, "error", err.Error())
	}
	password, err := bcrypt.GenerateFromPassword([]byte(body.Passowrd), 14)
	if err != nil {
		log.Fatalf("Error bcrypt password: %v", err)
	}
	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: string(password),
		Age:      body.Age,
	}
	database.DB.Create(&user)
	// TODO: if error eg. duplication return bad request 400
	return common.Response(c, 200, "Create User successful", user)
}
