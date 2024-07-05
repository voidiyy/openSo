package storage

import (
	"fmt"
	"iblan/cmd/structures"
)

type MemberStorage interface {
	CreateMember(nickname, password, email, category string) error
	UpdateMember(id int, nickname, password, email, category string) error
	DeleteMember(int) error
	GetMembers() ([]*structures.Member, error)
	GetMemberByID(int) (*structures.Member, error)
}

func (s *PostgresStore) CreateMember(nickname, password, email, category string) error {
	create := structures.Member{
		Nickname: nickname,
		Password: password,
		Email:    email,
		Category: category,
	}
	if err := s.db.Create(&create).Error; err != nil {
		return fmt.Errorf("failed to create a member: %v", err)
	}
	fmt.Printf("member %s created", nickname)
	return nil
}

func (s *PostgresStore) UpdateMember(id int, nickname, password, email, category string) error {
	member := structures.Member{}
	if err := s.db.Model(&member).Where("id = ?", id).Updates(map[string]interface{}{"nickname": nickname, "password": password, "email": email, "category": category}).Error; err != nil {
		return fmt.Errorf("failed to update a member: %v", err)
	}
	fmt.Printf("member %s updated", nickname)
	return nil
}

func (s *PostgresStore) DeleteMember(id int) error {
	if err := s.db.Delete(&structures.Member{}, id).Error; err != nil {
		return fmt.Errorf("error deleting member %v: %w", id, err)
	}
	return nil
}

func (s *PostgresStore) GetMemberByID(id int) (*structures.Member, error) {
	member := structures.Member{}
	query := s.db.Model(&structures.Member{}).Where("id = ?", id)

	if err := query.First(&member).Error; err != nil {
		return nil, fmt.Errorf("error getting a member %v: %w", id, err)
	}
	return &member, nil
}

func (s *PostgresStore) GetMembers() ([]*structures.Member, error) {
	var members []*structures.Member
	if err := s.db.Find(&members).Error; err != nil {
		return nil, fmt.Errorf("error getting members: %w", err)
	}
	return members, nil
}
