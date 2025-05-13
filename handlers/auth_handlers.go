package handlers

import "github.com/gofiber/fiber/v2"

func UserRegister(c *fiber.Ctx) error {
	// Renderizar la vista de registro de usuario
	return c.Render("auth/register", fiber.Map{
		"Title": "Registro de Usuario",
	})
}