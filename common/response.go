package common

import "github.com/gofiber/fiber/v2"

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(c *fiber.Ctx, statuCode int, msg string, data interface{}) error {
	resp := APIResponse{
		Success: true,
		Message: msg,
		Data:    data,
	}
	return c.Status(statuCode).JSON(&resp)
}

func ResponseError(c *fiber.Ctx, statuCode int, msg string, data interface{}) error {
	resp := APIResponse{
		Success: false,
		Message: msg,
		Data:    data,
	}
	return c.Status(statuCode).JSON(&resp)
}
