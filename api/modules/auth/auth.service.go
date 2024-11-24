package auth

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/vanthang24803/fiber-api/api/modules/auth/common"
	"golang.org/x/crypto/bcrypt"
)

type NormalResponse struct {
	Code    int    `json:"httpCode"`
	Message string `json:"message"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Payload struct {
	Email    string `json:"email"`
	FullName string `json:"fullName"`
}

func register(db *sql.DB, req *common.RegisterRequest) *NormalResponse {

	var existingEmail string
	err := db.QueryRow("SELECT email FROM users WHERE email = ?", req.Email).Scan(&existingEmail)

	if err != sql.ErrNoRows {
		if err != nil {
			log.Printf("Error checking email: %v", err)
			return &NormalResponse{
				Code:    400,
				Message: fmt.Sprintf("Database error: %v", err),
			}
		}
		return &NormalResponse{
			Code:    400,
			Message: "Email already in use",
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return &NormalResponse{
			Code:    400,
			Message: "Error hashing password",
		}
	}

	userId := uuid.New().String()

	_, err = db.Exec("INSERT INTO users (id, email, password, user_name, first_name, last_name) VALUES (?, ?, ?, ?, ?, ?)",
		userId, req.Email, string(hashedPassword), req.Username, req.FirstName, req.LastName)

	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return &NormalResponse{
			Code:    400,
			Message: err.Error(),
		}
	}

	var roleId string
	err = db.QueryRow("SELECT id FROM roles WHERE name = ?", "user").Scan(&roleId)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Role 'user' not found")
			return &NormalResponse{
				Code:    400,
				Message: "Role 'user' not found",
			}
		}
		log.Printf("Error fetching role: %v", err)
		return &NormalResponse{
			Code:    400,
			Message: fmt.Sprintf("Error fetching role: %v", err),
		}
	}

	_, err = db.Exec("INSERT INTO user_roles (user_id, role_id) VALUES (?, ?)", userId, roleId)
	if err != nil {
		log.Printf("Error assigning role to user: %v", err)
		return &NormalResponse{
			Code:    400,
			Message: fmt.Sprintf("Error assigning role: %v", err),
		}
	}

	return &NormalResponse{
		Code:    201,
		Message: "User registered successfully!",
	}
}

func login(db *sql.DB, req *common.LoginRequest) interface{} {
	var storedPassword string
	var userId string
	var email string
	var firstName string
	var lastName string
	err := db.QueryRow("SELECT id, password,email, first_name, last_name  FROM users WHERE email = ?", req.Email).Scan(&userId, &storedPassword, &email, &firstName, &lastName)

	if err != nil {
		log.Printf("Error retrieving user: %v", err)
		return &NormalResponse{
			Code:    400,
			Message: "Invalid email or password",
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(req.Password))
	if err != nil {
		return &NormalResponse{
			Code:    400,
			Message: "Invalid credentials",
		}
	}

	accessToken, refreshToken, err := generateTokens(userId, &Payload{
		Email:    email,
		FullName: firstName + " " + lastName,
	})
	if err != nil {
		log.Printf("Error generating tokens: %v", err)
		return &NormalResponse{
			Code:    500,
			Message: "Internal server error",
		}
	}

	return &TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func generateTokens(userId string, payload *Payload) (string, string, error) {
	accessTokenClaims := jwt.MapClaims{
		"sub":     userId,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"payload": &payload,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", "", fmt.Errorf("could not generate access token: %v", err)
	}

	refreshTokenClaims := jwt.MapClaims{
		"sub": userId,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("JWT_REFRESH")))
	if err != nil {
		return "", "", fmt.Errorf("could not generate refresh token: %v", err)
	}

	return accessTokenString, refreshTokenString, nil
}
