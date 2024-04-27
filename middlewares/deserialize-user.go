package middlewares

import (
	"fmt"
	"strings"

	"github.com/anucha-tk/go_jwt/common"
	"github.com/anucha-tk/go_jwt/initializers"
	"github.com/anucha-tk/go_jwt/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func DeserializeUser(c *fiber.Ctx) error {
	var tokenString string
	authorization := c.Get("Authorization")
	if strings.HasPrefix(authorization, "Bearer") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("token") != "" {
		tokenString = c.Cookies("token")
	}
	if tokenString == "" {
		return common.ResponseErrorMsg(c, fiber.StatusUnauthorized, "Error not Authorization")
	}
	config, _ := initializers.LoadConfig(".")
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("token error")
		}
		return []byte(config.JwtSecret), nil
	})
	if err != nil {
		return common.ResponseErrorMsg(c, fiber.StatusUnauthorized, fmt.Sprintf("Invalid token %v", err))
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return common.ResponseErrorMsg(c, fiber.StatusUnauthorized, "Error invalid token claims")
	}
	var user models.User
	result := initializers.DB.Where("id = ?", claims["iss"]).First(&user)
	if result.Error != nil && result.RowsAffected == 0 {
		return common.ResponseErrorMsg(c, fiber.StatusNotFound, "User not found")
	} else if result.Error != nil {
		return common.Response(c, fiber.StatusInternalServerError, "Something went wrong", result.Error.Error())
	}
	c.Locals("user", models.FilterUserRecord(&user))

	return c.Next()
}
