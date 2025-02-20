package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func FileList(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
	})
}
