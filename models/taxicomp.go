package models

import "time"

type TaxiComp struct {
	ID        int       `gorm:"primary key" json:"-"`
	CompTitle string    `gorm:"not null" json:"company_title"`
	IsDeleted bool      `gorm:"default false" json:"-"`
	DriverID  int       `json:"driver_id"`
	Driver    User      `gorm:"foreignKey:DriverID;references users(id)" json:"-"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`
}

func (TaxiComp) TableName() string {
	return "taxicompanies"
}

type GetAllTaxicomp struct {
	CompTitle string `gorm:"not null" json:"company_title"`
	IsDeleted bool   `gorm:"default false" json:"is_deleted"`
	DriverID  int    `json:"-"`
	Driver    User   `gorm:"foreignKey:DriverID;references users(id)" json:"driver_id"`
}
