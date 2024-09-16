package models

import "time"

type Route struct {
	ID         int       `gorm:"primary key" json:"-"`
	From       string    `gorm:"not null" json:"from"`
	Into       string    `gorm:"not null" json:"into"`
	Distance   int       `gorm:"not null" json:"distance"`
	Pricekm    int       `json:"-"`
	AllPrice   int       `json:"-"`
	ClientID   int       `gorm:"references users(id)" json:"-"`
	DriverID   int       `gorm:"references users(id)" json:"-"`
	IsResponse bool      `gorm:"default false" json:"-"`
	IsDeleted  bool      `gorm:"default false" json:"-"`
	CreatedAt  time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"-" gorm:"autoUpdateTime"`
}

func (Route) TableName() string {
	return "routes"
}

type GetRoutes struct {
	From       string    `gorm:"not null" json:"from"`
	Into       string    `gorm:"not null" json:"into"`
	Distance   int       `gorm:"not null" json:"distance"`
	Pricekm    int       `gorm:"default 0" json:"price_km"`
	AllPrice   int       `json:"all_price"`
	ClientID   int       `gorm:"references users(id)" json:"-"`
	DriverID   int       `gorm:"references users(id)" json:"-"`
	IsResponse bool      `gorm:"default false" json:"is_response"`
	IsDeleted  bool      `gorm:"default false" json:"-"`
	CreatedAt  time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"-" gorm:"autoUpdateTime"`
}

type Checkresponse struct {
	IsResponse bool `gorm:"default false" json:"is_response"`
}
