package middlewares

import (
	"mynotes/database"
	"mynotes/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store = session.New()

func AuthRequired(c *fiber.Ctx) error {
	// Obtener la sesión
	sess, err := Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error al obtener la sesión")
	}

	// Verificar si el usuario está autenticado
	if sess.Get("userID") == nil {
		return c.Redirect("/auth/login", 302)
	}

	// Continuar con la solicitud
	return c.Next()
}

func CurrentUser(c *fiber.Ctx) error {
	// Obtener la sesión
	sess, err := Store.Get(c)
	if err != nil {
		return c.Next()
	}

	// Obtener el userID de la sesión
	userID := sess.Get("userID")
	if userID == nil {
		return c.Next()
	}

	// Buscar el usuario en la base de datos
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Next()
	}

	// Pasar el usuario al contexto
	c.Locals("currentUser", user)
	return c.Next()
}