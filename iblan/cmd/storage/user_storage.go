package storage

import (
	"iblan/cmd/structures"
)

//qwerzxcfgh2

import (
	"fmt"
	_ "github.com/lib/pq"
)

type UserStorage interface {
	CreateUser(nickname, email, password string) error
	UpdateUser(id int, nickname, password, email string) (*structures.User, error)
	DeleteUser(int) error
	GetUserByID(id int) (*structures.User, error)
	GetUsers() ([]*structures.User, error)
}

func (s *PostgresStore) GetUsers() ([]*structures.User, error) {
	users := []*structures.User{}
	if err := s.db.First(&users); err != nil {
		return nil, fmt.Errorf("error selecting users: %v", err)
	}
	return users, nil
}

func (s *PostgresStore) CreateUser(nickname, email, password string) error {
	create := structures.User{
		Nickname: nickname,
		Email:    email,
		Password: password,
	}
	if err := s.db.Create(create).Error; err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}
	fmt.Printf("user %s created", nickname)
	return nil
}

func (s *PostgresStore) GetUserByID(id int) (*structures.User, error) {
	user := structures.User{}
	if err := s.db.First(&user, "id = ?", id); err != nil {
		return nil, fmt.Errorf("error getting user: %v", err)
	}
	return &user, nil
}

func (s *PostgresStore) UpdateUser(id int, nickname, password, email string) (*structures.User, error) {
	user := structures.User{}

	if err := s.db.Model(&user).Where("id = ?", id).Updates(map[string]interface{}{"nickname": nickname, "password": password, "email": email}).Error; err != nil {
		return nil, fmt.Errorf("error updating user: %v", err)
	}
	fmt.Printf("user %s updated", nickname)
	return &user, nil
}
func (s *PostgresStore) DeleteUser(id int) error {
	if err := s.db.Delete(&structures.User{}).Where("id = ?", id); err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}
	fmt.Printf("user %s deleted", id)
	return nil
}

//*********************************************************
