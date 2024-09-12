package models

import "time"

type TaxiComp struct {
	ID        int       `gorm:"primary key"`
	CompTitle string    `gorm:"not null" json:"company_title"`
	IsDeleted bool      `gorm:"default false" json:"-"`
	UserID    int       `gorm:"references users(id)" json:"-"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`
}

func (TaxiComp) TableName() string {
	return "taxicompanies"
}
