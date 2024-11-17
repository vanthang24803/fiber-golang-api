package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/vanthang24803/fiber-api/internal/config"
)

func ConnectionDB() *sql.DB {
	config.LoadEnvFile()

	db_username := os.Getenv("DB_USERNAME")
	db_pass := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%v port=%v sslmode=disable", db_username, db_name, db_pass, db_host, db_port)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database: ", err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err)
		return nil
	}

	log.Println("Database connection established successfully.")
	return db
}
