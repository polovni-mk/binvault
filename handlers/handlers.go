package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func BucketList(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
	})
}
