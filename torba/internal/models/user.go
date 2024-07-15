package model

import (
	"context"
	"log"
	"time"
)

type UserStorage interface {
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, id uint32, user *User) (*User, error)
	DeleteUser(ctx context.Context, id uint32) error
	GetUserByID(ctx context.Context, id uint32) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
}

type User struct {
	ID              uint32    `json:"id"`
	Username        string    `json:"username"`
	Email           string    `json:"email"`
	PasswordHash    string    `json:"password_hash"`
	DonationSum     float64   `json:"donation_sum"`
	ProfileImageURL string    `json:"profile_image_url"`
	CreatedAt       time.Time `json:"created_at"`
	LastLogin       time.Time `json:"last_login"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (p *Postgres) UserExistsByID(ctx context.Context, id uint32) bool {

	query := "SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)"

	var exists bool
	err := p.DB.QueryRow(ctx, query, id).Scan(&exists)
	if err != nil {
		log.Printf("Error checking user existence: %v", err)
		return false
	}

	return exists
}

func (p *Postgres) CreateUser(ctx context.Context, user *User) error {

	return nil
}

func (p *Postgres) UpdateUser(ctx context.Context, id uint32, user *User) (*User, error) {

	return nil, nil
}

func (p *Postgres) DeleteUser(ctx context.Context, id uint32) error {

	return nil
}

func (p *Postgres) GetUserByID(ctx context.Context, id uint32) (*User, error) {

	return nil, nil
}

func (p *Postgres) GetUserByUsername(ctx context.Context, name string) (*User, error) {
	return nil, nil
}
