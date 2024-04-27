package controllers

import (
	"strings"
	"time"

	"github.com/anucha-tk/go_jwt/common"
	"github.com/anucha-tk/go_jwt/initializers"
	"github.com/anucha-tk/go_jwt/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	payload := new(models.SignupInput)

	err := c.BodyParser(&payload)
	if err != nil {
		return common.ResponseError(c, fiber.StatusBadRequest, "Error Body Parser", err)
	}
	errors := models.ValidateStruct(payload)
	if errors != nil {
		return common.ResponseError(c, fiber.StatusUnprocessableEntity, "Error invalid body", errors)
	}
	if payload.Password != payload.PasswordConfirm {
		return common.ResponseErrorMsg(c, fiber.StatusUnprocessableEntity, "Error password and passowrdConfirm not match")
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return common.ResponseErrorMsg(c, fiber.StatusInternalServerError, "Error generate hash password")
	}
	var existUser models.User

	nameExist := initializers.DB.Where("name = ?", payload.Name).First(&existUser)
	if nameExist.RowsAffected > 0 {
		return common.ResponseMsg(c, fiber.StatusConflict, "Error: Name already exist")
	}
	emailExist := initializers.DB.Where("email = ?", payload.Email).First(&existUser)
	if emailExist.RowsAffected > 0 {
		return common.ResponseMsg(c, fiber.StatusConflict, "Error: Email already exist")
	}
	user := models.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: string(hashPassword),
		Age:      payload.Age,
	}
	result := initializers.DB.Create(&user)
	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return common.ResponseMsg(c, fiber.StatusConflict, "Error name or email Exist")
	} else if result.Error != nil {
		return common.Response(c, fiber.StatusInternalServerError, "Something went wrong", result.Error.Error())
	}
	return common.Response(c, 200, "Create User successful", user)
}

func Login(c *fiber.Ctx) error {
	payload := new(models.SinginInput)
	err := c.BodyParser(&payload)
	if err != nil {
		return common.ResponseError(c, fiber.StatusInternalServerError, "error", err.Error())
	}
	errors := models.ValidateStruct(payload)
	if errors != nil {
		return common.ResponseError(c, fiber.StatusUnprocessableEntity, "Error invalid body", errors)
	}
	var user models.User
	result := initializers.DB.Where("email = ?", payload.Email).First(&user)
	if result.Error != nil && result.RowsAffected == 0 {
		return common.ResponseErrorMsg(c, fiber.StatusNotFound, "User not found")
	} else if result.Error != nil {
		return common.Response(c, fiber.StatusInternalServerError, "Something went wrong", result.Error.Error())
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return common.ResponseErrorMsg(c, fiber.StatusBadRequest, "Invalid password")
	}
	config, _ := initializers.LoadConfig(".")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": user.ID,
			"exp": time.Now().Add(config.JwtExpiresIn).Unix(),
		})
	tokenString, err := token.SignedString([]byte(config.JwtSecret))
	if err != nil {
		return common.ResponseErrorMsg(c, fiber.StatusInternalServerError, "could't not login")
	}
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		MaxAge:   60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	return common.Response(c, fiber.StatusOK, "Login success", fiber.Map{"accessToken": tokenString})
}

func Logout(c *fiber.Ctx) error {
	expired := time.Now().Add(-time.Hour * 24)
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: expired,
	})
	return common.ResponseMsg(c, fiber.StatusOK, "Logout success")
}
