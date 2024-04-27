package common

import "github.com/gofiber/fiber/v2"

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type APIResponseMsg struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func Response(c *fiber.Ctx, statuCode int, msg string, data interface{}) error {
	resp := APIResponse{
		Success: true,
		Message: msg,
		Data:    data,
	}
	return c.Status(statuCode).JSON(&resp)
}

func ResponseMsg(c *fiber.Ctx, statuCode int, msg string) error {
	resp := APIResponseMsg{
		Success: true,
		Message: msg,
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

func ResponseErrorMsg(c *fiber.Ctx, statuCode int, msg string) error {
	resp := APIResponseMsg{
		Success: false,
		Message: msg,
	}
	return c.Status(statuCode).JSON(&resp)
}
