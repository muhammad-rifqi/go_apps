package controllers

import "github.com/gofiber/fiber/v2"

func UserControllerShow(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"id":       "1",
		"username": "rifqi",
		"password": "12345",
	})
}
