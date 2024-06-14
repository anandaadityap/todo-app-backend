package helper

import (
	"github.com/gofiber/fiber/v2"
)

type ResponseWithData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponseWithoutData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Response(c *fiber.Ctx, code int, message string, payload interface{}) error {
	var result any
	status := "success"

	if code > 400 {
		status = "failed"
	}

	if payload != nil {
		result = &ResponseWithData{
			Status:  status,
			Message: message,
			Data:    payload,
		}
	} else {
		result = &ResponseWithoutData{
			Status:  status,
			Message: message,
		}
	}

	return c.Status(code).JSON(result)
}
