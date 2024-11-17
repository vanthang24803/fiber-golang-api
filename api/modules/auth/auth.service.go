package auth

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/vanthang24803/fiber-api/internal/models"
)

func findOne(db *sql.DB, userID uuid.UUID) (*models.User, error) {
	query := "SELECT id, username, first_name, last_name FROM users WHERE id = $1"
	row := db.QueryRow(query, userID)

	var user models.User
	err := row.Scan(&user.Id, &user.Username, &user.FirstName, &user.LastName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("error fetching user: %v", err)
	}
	return &user, nil
}

func findAll(db *sql.DB) ([]models.User, error) {
	query := "SELECT id, user_name, first_name, last_name , created_at, updated_at FROM users"
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching users: %v", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning user: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading rows: %v", err)
	}

	return users, nil
}
