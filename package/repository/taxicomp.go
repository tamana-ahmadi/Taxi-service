package repository

import (
	"Taxi_service/db"
	"Taxi_service/logger"
	"Taxi_service/models"
)

func InsertTaxicomps(txcm models.TaxiComp) error {

	err := db.GetconnectDB().Create(&txcm).Error
	if err != nil {
		logger.Error.Printf("[repository.InsertTaxicomps]error in added taxi company %s\n", err.Error())
	}
	return nil
}
func EditTaxicomps(comptitle string, id int) error {
	err := db.GetconnectDB().Save(&models.TaxiComp{ID: id, CompTitle: comptitle}).Error
	if err != nil {
		logger.Error.Printf("[repository.edittaxicomps]error in update taxi company %s\n", err.Error())
	}
	return nil
}

func SoftDeleteTaxicomps(isdeleted bool, id int) error {
	err := db.GetconnectDB().Model(&models.TaxiComp{}).Where("id=?", id).Update("is_deleted", isdeleted).Error
	if err != nil {
		logger.Error.Printf("[repository.softdeletetaxicomps]error in deleted taxi company %s\n", err.Error())
	}
	return nil
}

func GetAllTaxicomps(isdeletedt, isblocked, isdeletedu bool) (txcm []models.TaxiComp, err error) {
	err = db.GetconnectDB().Preload("User").Joins("Join users ON users.id=taxicompanies.driver_id").Where("taxicompanies.is_deleted=?", isdeletedt).Where("users.is_blocked=? AND users.is_deleted=?", isblocked, isdeletedu).Find(&txcm).Error
	if err != nil {
		logger.Error.Printf("[repository.getalltaxicomps]error in getting all taxi companies %s\n", err.Error())
		return txcm, err
	}
	return txcm, nil
}
func GetAllTaxicompsByID(isdeleted, isblocked, isdeletedu bool, id int) (txcm []models.TaxiComp, err error) {
	err = db.GetconnectDB().Preload("User").Joins("Join users ON users.id=taxicompanies.driver_id").Where("taxicompanies.is_deleted=?", isdeleted).Where("taxicompanies.id=?", id).Where("users.is_blocked=? AND users.is_deleted=?", isblocked, isdeletedu).Find(&txcm).Error
	if err != nil {
		logger.Error.Printf("[repository.getalltaxicompsbyid]error in getting all taxi company by id %s\n", err.Error())
		return txcm, err
	}
	return txcm, nil
}
