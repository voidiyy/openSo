package postgres

import (
	"context"
	"errors"
	"strconv"
)

func (p *DB) UserExists(ctx context.Context, username string) error {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE username=$1)"

	err := p.pdb.QueryRow(ctx, query, username).Scan(&exists)
	if err != nil {
		return errors.New("cant check if user exists")
	}

	if exists {
		return errors.New("username already exists")
	}
	return nil
}

func (p *DB) EmailExistsU(ctx context.Context, email string) error {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)"

	err := p.pdb.QueryRow(ctx, query, email).Scan(&exists)
	if err != nil {
		return errors.New("cant check if user exists")
	}
	if exists {
		return errors.New("email already exists")
	}
	return nil
}

type CreateUserParams struct {
	Username     string
	Email        string
	PasswordHash string
}

func (p *DB) CreateUser(ctx context.Context, arg *CreateUserParams) error {
	query := "INSERT INTO users (username, email, password_hash, created_at) VALUES ($1, $2, $3, now())"

	_, err := p.pdb.Exec(ctx, query, arg.Username, arg.Email, arg.PasswordHash)
	if err != nil {
		p.errl.Printf("Error inserting user: %v", err)
		return err
	}

	p.infol.Printf("User successfully created :%v ", arg.Username)
	return nil
}

type UpdateUserParams struct {
	UserID       int64
	Username     string
	Email        string
	PasswordHash string
}

func (p *DB) UpdateUser(ctx context.Context, arg *UpdateUserParams) error {
	query := "UPDATE users SET username = $2, email = $3, password_hash = $4, updated_at = now() WHERE user_id = $1"

	_, err := p.pdb.Exec(ctx, query, arg.UserID, arg.Username, arg.Email, arg.PasswordHash)
	if err != nil {
		p.errl.Printf("Error updating user: %v", err)
		return err
	}

	p.infol.Printf("User successfully updated :%v ", arg.Username)
	return nil
}

func (p *DB) DeleteUser(ctx context.Context, userID int64) error {
	query := "DELETE FROM users WHERE user_id = $1"

	_, err := p.pdb.Exec(ctx, query, userID)
	if err != nil {
		p.errl.Printf("Error deleting user: %v", err)
		return err
	}

	p.infol.Printf("User successfully deleted :%v", userID)
	return nil
}

func (p *DB) GetUserByID(ctx context.Context, userID int64) (*User, error) {
	query := `SELECT * FROM users WHERE user_id = $1 LIMIT 1`

	user := User{}

	err := p.pdb.QueryRow(ctx, query, userID).Scan(&user.UserID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.LastLogin, &user.UpdatedAt)
	if err != nil {
		p.errl.Printf("Error getting user by ID: %v", err)
		return &User{}, errors.New("user does not exist id: " + strconv.FormatInt(userID, 10))
	}

	p.infol.Printf("fetched user by ID :%v", user.Username)
	return &user, nil
}

func (p *DB) GetUserByName(ctx context.Context, username string) (*User, error) {
	query := "SELECT * FROM users WHERE username = $1 LIMIT 1"

	var user User

	err := p.pdb.QueryRow(ctx, query, username).Scan(&user.UserID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.LastLogin, &user.UpdatedAt)
	if err != nil {
		p.errl.Printf("Error getting user by name: %v", err)
		return &User{}, err
	}

	p.infol.Printf("fetched user by name:%v", user.Username)
	return &user, nil
}

func (p *DB) ListUserID(ctx context.Context) ([]User, error) {
	query := "SELECT * FROM users ORDER BY user_id"

	var users []User

	rows, err := p.pdb.Query(ctx, query)
	if err != nil {
		p.errl.Printf("Error listing users by ID: %v", err)
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var u User
		if err := rows.Scan(
			&u.UserID,
			&u.Username,
			&u.Email,
			&u.PasswordHash,
			&u.CreatedAt,
			&u.LastLogin,
			&u.UpdatedAt); err != nil {
			p.errl.Printf("Error scanning users by ID in loop: %v", err)
			return users, err
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		p.errl.Printf("Error scanning users by ID: %v", err)
		return users, err
	}

	p.infol.Printf("listing users by ID: %v", users)
	return users, nil
}

func (p *DB) ListUserName(ctx context.Context) ([]User, error) {
	query := "SELECT * FROM users ORDER BY username"

	var users []User

	rows, err := p.pdb.Query(ctx, query)
	if err != nil {
		p.errl.Printf("Error listing users: %v", err)
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var u User
		if err := rows.Scan(
			&u.UserID,
			&u.Username,
			&u.Email,
			&u.PasswordHash,
			&u.CreatedAt,
			&u.LastLogin,
			&u.UpdatedAt); err != nil {
			p.errl.Printf("Error scanning users by name in loop: %v", err)
			return users, err
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		p.errl.Printf("Error scanning users by name: %v", err)
		return users, err
	}

	p.infol.Printf("listing users by name: %v ", users)
	return users, nil
}
