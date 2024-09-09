package models

import "time"

type Ord struct {
	ID         int       `gorm:"primary_key"`
	ClientID   int       `gorm:"references users(id)" json:"client_id"`
	DriverID   int       `gorm:"references users(id)" json:"driver_id"`
	RouteID    int       `gorm:"references routes(id)" json:"route_id"`
	Distance   int       `gorm:"foreignKey:RouteID" json:"distance"`
	Price      int       `gorm:"foreignKey:RouteID" json:"price"`
	TaxicompID int       `gorm:"references taxicompanies(id)" json:"taxicomp_id"`
	Comptitle  string    `gorm:"foreignKey:TaxicompID" json:"comp_title"`
	IsDone     bool      `gorm:"default false" json:"is_done"`
	IsDeleted  bool      `gorm:"default false" json:"is_deleted"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Ord) TableName() string {
	return "orders"
}
