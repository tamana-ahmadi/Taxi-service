package models

import "time"

type Route struct {
	ID         int       `gorm:"primary key"`
	From       string    `gorm:"not null" json:"from"`
	Into       string    `gorm:"not null" json:"into"`
	Distance   int       `gorm:"not null" json:"distance"`
	Price      int       `gorm:"default false" json:"price"`
	ClientID   int       `gorm:"references users(id)" json:"client_id"`
	DriverID   int       `gorm:"references users(id)" json:"driver_id"`
	IsResponse bool      `gorm:"default false" json:"-"`
	IsDeleted  bool      `gorm:"default false" json:"-"`
	CreatedAt  time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"-" gorm:"autoUpdateTime"`
}

func (Route) TableName() string {
	return "routes"
}
