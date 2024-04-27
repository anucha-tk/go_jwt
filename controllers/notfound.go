package controllers

import (
	"fmt"

	"github.com/anucha-tk/go_jwt/common"
	"github.com/gofiber/fiber/v2"
)

func Notfound(c *fiber.Ctx) error {
	path := c.Path()
	return common.ResponseErrorMsg(c, fiber.StatusOK, fmt.Sprintf("Path %s not found", path))
}
