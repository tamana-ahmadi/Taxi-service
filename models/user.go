package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"primary_key"`
	FullName  string    `json:"full_name"`
	Username  string    `json:"username" gorm:"unique"`
	Password  string    `json:"password,omitempty" gorm:"not null"`
	Role      string    `json:"role"`
	Rating    int       `json:"rating"`
	IsBlocked bool      `json:"is_blocked" gorm:"default:false"`
	IsDeleted bool      `json:"is_deleted" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (User) TableName() string {
	return "users"
}

type SwagSignUp struct {
	FullName string `json:"full_name"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password" gorm:"not null"`
	Role     string `json:"role"`
}

type SwagSignIn struct {
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password" gorm:"not null"`
}

type Rating struct {
	Rating int `json:"rating"`
}
