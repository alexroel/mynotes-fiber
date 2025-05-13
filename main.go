package main

import (
	"log"
	"mynotes/database"
	"mynotes/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Conectar a la base de datos
	database.ConnectDB()

	// Inicializar el motor de plantillas
	engine := html.New("./views", ".html")

	// Inicializar Fiber
	app := fiber.New(fiber.Config{
		Views: engine,
		ViewsLayout: "layouts/main",
	})

	// Configurar rutas
	routers.SetupRoutes(app)

	// Inicializar servidor
	log.Fatal(app.Listen(":3000"))

}