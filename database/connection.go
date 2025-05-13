package database

import (
	"fmt"
	"log"
	"mynotes/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	// Configuración de conexión a PostgreSQL
	// dsn := "host=localhost user=alexroel password=123456 dbname=mynotes_db port=5432 sslmode=disable TimeZone=UTC"
	dsn := os.Getenv("DB_DSN")

	var err error
	
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}

	// Migraciones automáticas
	err = DB.AutoMigrate(&models.User{}, &models.Note{})
	if err != nil {
		log.Fatalf("Error al realizar las migraciones: %v", err)
	}

	fmt.Println("Conexión a la base de datos y migraciones realizadas con éxito")
}