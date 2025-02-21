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

	dbURL := os.Getenv("postgresql://postgres:Vasu%401643%23@db.zyigqdeabwscdtebkjms.supabase.co:5432/postgres")
	database, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	DB = database
	log.Println("Database connected successfully")
}
