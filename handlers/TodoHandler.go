package handlers

import "github.com/gofiber/fiber/v2"

func CreateTodo(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success"})
}
