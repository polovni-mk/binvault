package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func BucketList(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
	})
}

func BucketCreate(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
	})
}

func BucketGet(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
	})
}

func BucketDelete(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
	})
}
