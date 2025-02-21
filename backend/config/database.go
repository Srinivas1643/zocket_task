package config

import (
	"log"
	"os"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

var DB *sqlx.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DB_URL")
	database, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	DB = database
	log.Println("Database connected successfully")
}
