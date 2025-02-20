package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func FileList(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
	})
}

func FileCreate(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
	})
}

func FileGet(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
	})
}

func FileDelete(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
	})
}
