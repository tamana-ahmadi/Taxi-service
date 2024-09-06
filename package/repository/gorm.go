package repository

import (
	"Taxi_service/errs"
	"errors"

	"gorm.io/gorm"
)

func translateErrors(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errs.ErrRecordNotFound
	}
	return err
}
