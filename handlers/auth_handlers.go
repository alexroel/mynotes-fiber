package handlers

import (
	"mynotes/database"
	"mynotes/middlewares"
	"mynotes/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func UserRegister(c *fiber.Ctx) error {
	if c.Method() == "POST" {
		// Obtener los datos del formulario
		name := c.FormValue("name")
		email := c.FormValue("email")
		password := c.FormValue("password")
		

		// Validar los datos (esto es solo un ejemplo, deberías hacer una validación más robusta)
		if name == "" || password == "" || email == "" {
			return c.Render("auth/register", fiber.Map{
				"Title":   "Registro de Usuario",
				"Error":   "Todos los campos son obligatorios",
				"Success": false,
			})
		}

		// Verificar si el usuario ya con el correo electrónico existe
		var existingUser models.User
		if err := database.DB.Where("email = ?", email).First(&existingUser).Error; err == nil {
			return c.Render("auth/register", fiber.Map{
				"Title":   "Registro de Usuario",
				"Error":   "El correo electrónico ya está registrado",
				"Success": false,
			})
		}

		// Hash de la contraseña (esto es solo un ejemplo, deberías usar una función de hash segura)
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return c.Render("auth/register", fiber.Map{
				"Title":   "Registro de Usuario",
				"Error":   "Error al procesar la contraseña",
				"Success": false,
			})
		}
		// Aquí podrías agregar la lógica para guardar el usuario en la base de datos
		user := models.User{
			Name:     name,
			Email:    email,
			Password: string(hashedPassword),
		}

		// Guardar el usuario en la base de datos
		if err := database.DB.Create(&user).Error; err != nil {
			return c.Render("auth/register", fiber.Map{
				"Title":   "Registro de Usuario",
				"Error":   "Error al guardar el usuario",
				"Success": false,
			})
		}

		return c.Redirect("/auth/login", 302)
	}
	// Renderizar la vista de registro de usuario
	return c.Render("auth/register", fiber.Map{
		"Title": "Registro de Usuario",
	})
}

func UserLogin(c *fiber.Ctx) error {
    if c.Method() == "POST" {
        email := c.FormValue("email")
        password := c.FormValue("password")

        if email == "" || password == "" {
            return c.Render("auth/login", fiber.Map{
                "Title":   "Inicio de Sesión",
                "Error":   "Todos los campos son obligatorios",
                "Success": false,
				"currentUser": c.Locals("currentUser"),
            })
        }

        var user models.User
        if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
            return c.Render("auth/login", fiber.Map{
                "Title":   "Inicio de Sesión",
                "Error":   "Credenciales incorrectas",
                "Success": false,
            })
        }

        if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
            return c.Render("auth/login", fiber.Map{
                "Title":   "Inicio de Sesión",
                "Error":   "Credenciales incorrectas",
                "Success": false,
            })
        }

        // Crear sesión
        sess, err := middlewares.Store.Get(c)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).SendString("Error al crear la sesión")
        }

        // Guardar el ID del usuario en la sesión
        sess.Set("userID", user.ID)
        if err := sess.Save(); err != nil {
            return c.Status(fiber.StatusInternalServerError).SendString("Error al guardar la sesión")
        }

        return c.Redirect("/notes", 302)
    }
    return c.Render("auth/login", fiber.Map{
        "Title": "Inicio de Sesión",
    })
}

// UserLogout maneja el cierre de sesión del usuario
func UserLogout(c *fiber.Ctx) error {
    sess, err := middlewares.Store.Get(c)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).SendString("Error al obtener la sesión")
    }

    // Destruir la sesión
    if err := sess.Destroy(); err != nil {
        return c.Status(fiber.StatusInternalServerError).SendString("Error al destruir la sesión")
    }

    return c.Redirect("/", 302)
}