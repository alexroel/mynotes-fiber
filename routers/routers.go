package routers

import (
	"mynotes/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Rutas de la aplicación
	app.Get("/", handlers.Index)
	app.Get("/about", handlers.About)

	// Rutas de notas
	app.Get("/notes", handlers.Notes)
	app.Post("/notes", handlers.CreateNote)

	// Rutas de usuarios
	app.Get("/auth/register", handlers.UserRegister)
}