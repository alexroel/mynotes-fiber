package routers

import (
	"mynotes/handlers"
	"mynotes/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Registrar el middleware CurrentUser
    app.Use(middlewares.CurrentUser)

	// Configuración de archivos estáticos
	app.Static("/", "./public") 

	// Rutas de la aplicación
	app.Get("/", handlers.Index)
	app.Get("/about", handlers.About)

	// Rutas de notas
	notes := app.Group("/notes", middlewares.AuthRequired)
	notes.Get("/", handlers.Notes)         // Listar notas
	notes.Post("/", handlers.CreateNote)   // Crear nota
	notes.Get("/edit/:id", handlers.EditNote)
	notes.Post("/edit", handlers.EditNote) // Editar nota
	notes.Get("/:id", handlers.DeleteNote) // Eliminar nota

	// Rutas de usuarios
	app.Get("/auth/register", handlers.UserRegister)
	app.Post("/auth/register", handlers.UserRegister)
	app.Get("/auth/login", handlers.UserLogin)
	app.Post("/auth/login", handlers.UserLogin)
	app.Get("/auth/logout", handlers.UserLogout)
}