package models

import "time"

type Order struct {
	ID         int       `gorm:"primary_key"`
	UserID     int       `gorm:"references users(id)" json:"user_id"`
	RouteID    int       `gorm:"references routes(id)" json:"route_id"`
	TaxicompID int       `gorm:"references taxicompanies(id)" json:"taxicomp_id"`
	IsDone     bool      `gorm:"default false" json:"is_done"`
	IsDeleted  bool      `gorm:"default false" json:"is_deleted"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Order) TableName() string {
	return "orders"
}
