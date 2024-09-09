package db

import (
	"Taxi_service/models"
	"errors"
)

func Migrate() error {
	err := conn.AutoMigrate(models.Route{}, models.TaxiComp{}, models.User{})
	if err != nil {
		return errors.New("Failed to begin transaction: " + err.Error())
	}

	return nil
}
