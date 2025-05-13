package database

import (
	"fmt"
	"log"
	"mynotes/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Configuración de conexión a PostgreSQL
	dsn := "host=localhost user=alexroel password=123456 dbname=mynotes_db port=5432 sslmode=disable TimeZone=UTC"

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