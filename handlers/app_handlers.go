package handlers

import "github.com/gofiber/fiber/v2"

func Index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":       "¡Bienvenido a Mynotes!",
		"Description": "Una aplicación para tomar notas de manera sencilla y rápida. Crea una cuenta y empieza a organizar tus ideas.",
	})
}

func About(c *fiber.Ctx) error {
	return c.Render("about", fiber.Map{
		"Title": "Acerca de MyNotes",
	})
}