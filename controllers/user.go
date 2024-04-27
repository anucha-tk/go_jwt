package controllers

import (
	"github.com/anucha-tk/go_jwt/common"
	"github.com/anucha-tk/go_jwt/models"
	"github.com/gofiber/fiber/v2"
)

func GetMe(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserResponse)
	return common.Response(c, fiber.StatusOK, "GetMe success", user)
}
