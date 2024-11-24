package models

import (
	"time"
)

type User struct {
	Id        string    `json:"id"`
	Username  string    `json:"user_name"`
	FirstName string    `json:"first_name"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
