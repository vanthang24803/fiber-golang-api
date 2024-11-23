package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vanthang24803/fiber-api/internal/config"
)

func ConnectionDB() *sql.DB {
	config.LoadEnvFile()

	db_username := os.Getenv("DB_USERNAME")
	db_pass := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_username, db_pass, db_host, db_port, db_name)

	db, err := sql.Open("mysql", connStr)
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
