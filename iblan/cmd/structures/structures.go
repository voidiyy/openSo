package structures

import (
	"gorm.io/gorm"
	"math/rand"
)

type User struct {
	gorm.Model
	Nickname string `json:"nickname" gorm:"unique" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email" gorm:"unique"`
}

func NewUser(nickname, password, email string) *User {
	return &User{
		Nickname: nickname,
		Email:    email,
		Password: password,
	}
}

//********************************************************************

type Member struct {
	gorm.Model
	Nickname string `json:"nickname" gorm:"unique" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email" gorm:"unique" gorm:"not null"`
	Category string `json:"category"`
	Loyalty  int    `json:"loyalty"`
	Verified bool   `json:"Verified"`
}

func NewMember(nickname, password, email, category string) *Member {
	return &Member{
		Nickname: nickname,
		Password: password,
		Email:    email,
		Category: category,
		Loyalty:  rand.Intn(10),
		Verified: false,
	}
}

//**********************************************************************

type Article struct {
	gorm.Model
	Title    string `json:"title" gorm:"not null"`
	Category string `json:"category" gorm:"not null"`
	Body     string `json:"body" gorm:"not null"`
	Payments string `json:"payments" gorm:"not null"`
	Link     string `json:"link" gorm:"not null"`
}

func NewArticle(title, category, body, payments, link string) *Article {
	return &Article{
		Title:    title,
		Category: category,
		Body:     body,
		Payments: payments,
		Link:     link,
	}
}
