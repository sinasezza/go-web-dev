package models

import (
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash"`
}

type UserService struct {
	DB *sql.DB
}

func (us *UserService) Create(email, password string) (*User, error) {
	email = strings.ToLower(email)
	password = strings.TrimSpace(password)
	if len(email) == 0 || len(password) == 0 {
		return nil, fmt.Errorf("length of email and password must be greater than zero")
	}

	// Check if user already exists
	_, err := us.GetByEmail(email)
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("user already exists")
	}

	// At this point, we know the user doesn't exist, so we can create it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := User{Email: email, PasswordHash: string(hashedPassword)}
	query := "INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id"
	err = us.DB.QueryRow(query, user.Email, user.PasswordHash).Scan(&user.ID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (us *UserService) GetByEmail(email string) (*User, error) {
	email = strings.ToLower(email)
	query := "SELECT id, email, password_hash FROM users WHERE email=$1"
	row := us.DB.QueryRow(query, email)

	var user User // Change: Initialize user as a value, not a pointer
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash)

	switch {
	case err == sql.ErrNoRows:
		return nil, sql.ErrNoRows
	case err != nil:
		return nil, err
	default:
		return &user, nil // Return the address of user
	}
}
