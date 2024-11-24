package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db_username := os.Getenv("DB_USERNAME")
	db_pass := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true",
		db_username, db_pass, db_host, db_port, db_name)

	cmd := exec.Command("goose", "-dir=internal/database/migrations", "mysql", connStr, "up")
	cmd.Dir = "../"

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmdErr := cmd.Run()
	if cmdErr != nil {
		log.Fatalf("Error running goose command: %v", cmdErr)
	}

	fmt.Println("Migration completed successfully ✅")
}
