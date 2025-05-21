package config

import (
	"aplikasi-daftar-belanja/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// dsn := "user=postgres password=postgres dbname=aplikasi_daftar_belanja port=5432"
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database %v", err)
	}

	DB = db

	//Auto migrate tabel item
	err = db.AutoMigrate(&models.Item{}, &models.User{})
	if err != nil {
		log.Fatalf("Failed to auto migrate : %v", err)
	}
	fmt.Println("Connected to database")

}