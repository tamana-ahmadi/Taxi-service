package models

import "time"

type Route struct {
	ID         int       `gorm:"primary key"`
	From       string    `gorm:"not null" json:"from"`
	Into       string    `gorm:"not null" json:"into"`
	Distance   int       `gorm:"not null" json:"distance"`
	Price      int       `gorm:"default false" json:"price"`
	UserID     int       `gorm:"references users(id)" json:"user_id"`
	IsResponse bool      `gorm:"default false" json:"is_response"`
	IsDeleted  bool      `gorm:"default false" json:"is_deleted"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Route) TableName() string {
	return "routes"
}
