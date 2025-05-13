package handlers

import (
	"mynotes/database"
	"mynotes/models"

	"github.com/gofiber/fiber/v2"
)

func Notes(c *fiber.Ctx) error {
	// Renderizar la vista de notas
	var notes []models.Note

	// Obtener las notas del modelo
	database.DB.Find(&notes)

	return c.Render("notes", fiber.Map{
		"Title": "Mis Notas",
		"Notes": notes,
	})

}

// CreateNote crea una nueva nota
func CreateNote(c *fiber.Ctx) error {
	// Obtener el título y el contenido de la nota del formulario
	title := c.FormValue("title")
	content := c.FormValue("content")

	// Crear una nueva nota
	note := models.Note{
		Title:   title,
		Content: content,
	}

	// Guardar la nota en la base de datos
	if err := database.DB.Create(&note).Error; err != nil {
		return c.Status(500).SendString("Error al crear la nota")
	}

	return c.Redirect("/notes")
}