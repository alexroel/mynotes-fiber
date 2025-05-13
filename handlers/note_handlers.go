package handlers

import (
	"mynotes/database"
	"mynotes/middlewares"
	"mynotes/models"

	"github.com/gofiber/fiber/v2"
)

func Notes(c *fiber.Ctx) error {
	// Obtener la sesión
	sess, err := middlewares.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error al obtener la sesión")
	}

	// Obtener el ID del usuario desde la sesión
	userID := sess.Get("userID")
	if userID == nil {
		return c.Redirect("/auth/login", 302)
	}

	// Renderizar las notas del usuario
	var notes []models.Note
	if err := database.DB.Where("user_id = ?", userID).Find(&notes).Error; err != nil {
		return c.Status(500).SendString("Error al obtener las notas")
	}

	return c.Render("notes", fiber.Map{
		"Title": "Mis Notas",
		"Notes": notes,
		"currentUser": c.Locals("currentUser"),
	})
}

// CreateNote crea una nueva nota
func CreateNote(c *fiber.Ctx) error {
	// Obtener la sesión
	sess, err := middlewares.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error al obtener la sesión")
	}

	// Obtener el ID del usuario desde la sesión
	userID := sess.Get("userID")
	if userID == nil {
		return c.Redirect("/auth/login", 302)
	}

	// Obtener el título y el contenido de la nota del formulario
	title := c.FormValue("title")
	content := c.FormValue("content")

	// Crear una nueva nota asociada al usuario
	note := models.Note{
		Title:   title,
		Content: content,
		UserID:  userID.(uint), // Asegúrate de que el campo UserID sea del tipo correcto
	}

	// Guardar la nota en la base de datos
	if err := database.DB.Create(&note).Error; err != nil {
		return c.Status(500).SendString("Error al crear la nota")
	}

	return c.Redirect("/notes")
}

// EditNote actualiza una nota existente
func EditNote(c *fiber.Ctx) error {
	if c.Method() == "GET" {
		// Obtener el ID de la nota desde los parámetros
		id := c.Params("id")
		// Buscar la nota en la base de datos
		var note models.Note
		if err := database.DB.First(&note, id).Error; err != nil {
			return c.Status(404).SendString("Nota no encontrada")
		}
		// Renderizar el formulario de edición
		return c.Render("notes", fiber.Map{
			"Title": "Editar Nota",
			"Note":  note,
			"currentUser": c.Locals("currentUser"),
		})
	}

	// Obtener el ID de la nota desde el formulario
	id := c.FormValue("id")
	title := c.FormValue("title")
	content := c.FormValue("content")

	// Buscar la nota en la base de datos
	var note models.Note
	if err := database.DB.First(&note, id).Error; err != nil {
		return c.Status(404).SendString("Nota no encontrada")
	}

	// Actualizar los campos de la nota
	note.Title = title
	note.Content = content

	// Guardar los cambios
	if err := database.DB.Save(&note).Error; err != nil {
		return c.Status(500).SendString("Error al actualizar la nota")
	}

	return c.Redirect("/notes")
}

// DeleteNote elimina una nota existente
func DeleteNote(c *fiber.Ctx) error {
	// Obtener el ID de la nota desde los parámetros
	id := c.Params("id")

	// Eliminar la nota de la base de datos
	if err := database.DB.Delete(&models.Note{}, id).Error; err != nil {
		return c.Status(500).SendString("Error al eliminar la nota")
	}

	return c.Redirect("/notes")
}