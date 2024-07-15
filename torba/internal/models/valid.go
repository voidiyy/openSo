package model

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func (p *Postgres) Validate(user *User) (*User, error) {

	validN, err := p.ValidateUserName(user.Username)
	if err != nil {
		log.Println("User name validation error:", err)
	}
	validE, err := p.ValidEmail(user.Email)
	if err != nil {
		log.Println("Email validation error:", err)
	}
	validP, err := p.HashPassword(user.PasswordHash)
	if err != nil {
		log.Println("Password hashing error:", err)
	}

	return &User{
		Username:     validN,
		Email:        validE,
		PasswordHash: validP,
	}, nil
}

func (p *Postgres) ValidateFloat64(str string) (float64, error) {
	float, err := strconv.ParseFloat(str, 64)
	if err != nil {
		// Check if the error is due to invalid syntax
		var numError *strconv.NumError
		if errors.As(err, &numError) {
			fmt.Printf("Cannot convert %s to float64: invalid syntax\n", str)
		}
		return float, err
	}
	return float, nil
}

func (p *Postgres) ValidateUserName(username string) (string, error) {
	username = strings.TrimSpace(username)
	if username == "" {
		return "", fmt.Errorf("username cannot be empty")
	}
	if len(username) < 3 || len(username) > 50 {
		return "", fmt.Errorf("username must be between 3 and 50 characters")
	}
	return username, nil
}

func (p *Postgres) ValidEmail(email string) (string, error) {
	var rxEmail = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	email = strings.TrimSpace(email)
	if email == "" {
		return "", fmt.Errorf("email cannot be empty")
	}
	if !rxEmail.MatchString(email) {
		return "", fmt.Errorf("invalid email format")
	}
	return email, nil
}

func (p *Postgres) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %v", err)
	}
	return string(hashedPassword), nil
}
