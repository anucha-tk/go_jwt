package controllers

import (
	"github.com/anucha-tk/go_jwt/common"
	"github.com/gofiber/fiber/v2"
)

func HealthChecker(c *fiber.Ctx) error {
	return common.ResponseMsg(c, fiber.StatusOK, "Golang JWT Authenticate App")
}
