package utils

import "github.com/gofiber/fiber/v2"

func JsonError(c *fiber.Ctx, status int, msg string) {
	c.Status(status).JSON(fiber.Map{
		"message": msg,
	})
}
