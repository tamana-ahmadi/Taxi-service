package models

import "time"

type Order struct {
	ID         int       `gorm:"primary_key"`
	ClientID   int       `gorm:"references users(id)" json:"client_id"`
	DriverID   int       `gorm:"references users(id)" json:"driver_id"`
	RouteID    int       `gorm:"references routes(id)" json:"route_id"`
	Distance   Route     `gorm:"foreignKey:RouteID" json:"distance"`
	Price      Route     `gorm:"foreignKey:RouteID" json:"price"`
	TaxicompID int       `gorm:"references taxicompanies(id)" json:"taxicomp_id"`
	Comptitle  TaxiComp  `gorm:"foreignKey:TaxicompID" json:"comp_title"`
	IsDone     bool      `gorm:"default false" json:"is_done"`
	IsDeleted  bool      `gorm:"default false" json:"is_deleted"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Order) TableName() string {
	return "orders"
}
